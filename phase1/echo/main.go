package main

import (
	"fmt"
	"net/http"
	"os"
	"text/template"

	"github.com/labstack/echo/v4"
)

type PageData struct {
	Title   string
	Message string
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		textTemplate, err := template.ParseFiles("phase1/echo/templates/text_example.html")
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Error parsing template")
		}
		data := PageData{
			Title:   "Welcome to Go Templates!",
			Message: "This is a simple example of using Go's html/template package.",
		}
		err = textTemplate.Execute(c.Response(), data)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Error parsing template")
		}
		return c.String(http.StatusOK, "Hello from Echo!")
	})
	hostname, _ := os.Hostname()
	e.GET("/service/1", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("Hello from behind Envoy (service 1)!. Host name is %s", hostname))
	})
	e.GET("/service/2", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("Hello from behind Envoy (service 2)!. Host name is %s", hostname))
	})
	e.GET("/healthz", func(c echo.Context) error {
		// healthCheckHandler is a simple HTTP handler that returns a 200 OK status.
		return c.String(http.StatusOK, "Service is healthy")
	})
	e.Logger.Fatal(e.Start(":9090"))
}
