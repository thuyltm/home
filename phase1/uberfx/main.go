package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

type Route interface {
	http.Handler
	Pattern() string
}

func (*EchoHandler) Pattern() string {
	return "/echo"
}

// EchoHandler is an http.Handler that copies its request body back to the response.
type EchoHandler struct {
	log *zap.Logger
}

// NewEchoHandler builds a new EchoHandler
func NewEchoHandler(log *zap.Logger) *EchoHandler {
	return &EchoHandler{log: log}
}

// HelloHandler is an HTTP handler that prints a greeting to the user
type HelloHandler struct {
	log *zap.Logger
}

func NewHelloHandler(log *zap.Logger) *HelloHandler {
	return &HelloHandler{log: log}
}

func (*HelloHandler) Pattern() string {
	return "/hello"
}

type GreetHandler struct {
	log *zap.Logger
}

func NewGreetHandler(log *zap.Logger) *GreetHandler {
	return &GreetHandler{log: log}
}

func (*GreetHandler) Pattern() string {
	return "/greet"
}

func (*GreetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := io.Copy(w, r.Body); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to handle request:", err)
	}
}

type WelcomeHandler struct {
	log *zap.Logger
}

func newWelcomeHandler(log *zap.Logger) *WelcomeHandler {
	return &WelcomeHandler{log: log}
}

func (*WelcomeHandler) Pattern() string {
	return "/welcome"
}

func (*WelcomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := io.Copy(w, r.Body); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to handle request:", err)
	}
}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.Error("Failed to read request", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	if _, err := fmt.Fprintf(w, "Hello, %s\n", body); err != nil {
		h.log.Error("Failed to write response", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

// ServerHTTP handles an HTTP request to the /echo endpoint
func (*EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := io.Copy(w, r.Body); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to handle request:", err)
	}
}

// NewServeMux builds a ServeMux that will route requests to the given Route
func NewServeMux(routes []Route) *http.ServeMux {
	mux := http.NewServeMux()
	for _, route := range routes {
		mux.Handle(route.Pattern(), route)
	}
	return mux
}

// AsRoute annotates the given constructor to state that
// it provides a route to the "routes" group
func AsRoute(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Route)),
		fx.ResultTags(`group:"routes"`),
	)
}

// NewHTTPServer builds an HTTP server that will begin serving requests
// when the Fx application starts
func NewHTTPServer(lc fx.Lifecycle, mux *http.ServeMux, log *zap.Logger) *http.Server {
	srv := &http.Server{Addr: ":8080", Handler: mux}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			log.Info("Starting HTTP server at", zap.String("addr", srv.Addr))
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return srv
}
func NewLogger() *zap.Logger {
	rawJSON := []byte(`{
	  "level": "debug",
	  "encoding": "json",
	  "outputPaths": ["stdout", "/tmp/logs"],
	  "errorOutputPaths": ["stderr"],
	  "initialFields": {"foo": "bar"},
	  "encoderConfig": {
	    "messageKey": "message",
	    "levelKey": "level",
	    "levelEncoder": "lowercase"
	  }
	}`)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	logger := zap.Must(cfg.Build())
	defer logger.Sync()

	logger.Info("logger construction succeeded")
	return logger
}

func main() {

	fx.New(
		fx.Provide(
			NewHTTPServer,
			fx.Annotate(
				NewServeMux,
				fx.ParamTags(`group:"routes"`),
			),
			NewLogger,
			//this fails because FX does not allow two instance of the same type
			//to be present in the container without annotating them
			fx.Annotate(
				NewEchoHandler,
				fx.As(new(Route)),
				fx.ResultTags(`group:"routes"`),
			),
			fx.Annotate(
				NewHelloHandler,
				fx.As(new(Route)),
				fx.ResultTags(`group:"routes"`),
			),
			AsRoute(NewGreetHandler),
			AsRoute(newWelcomeHandler),
		),
		fx.Invoke(func(*http.Server) {}),
		//you can use the same Zap logger for Fx's own logs as well
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
	).Run()
}
