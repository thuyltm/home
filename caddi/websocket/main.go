package main

import (
	"flag"
	"fmt"
	"home/caddi/websocket/conn"
	"log"
	"net/http"
	"os"
	"sort"
	"text/template"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
)

var addr = flag.String("addr", ":9090", "http service address")

type ProviderIndex struct {
	Providers    []string
	ProvidersMap map[string]string
}

func isAuthenticated(req *http.Request) bool {
	_, err := req.Cookie(USERID)
	return err == nil
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !isAuthenticated(r) {
			http.Redirect(w, r, "/auth/github", http.StatusFound)
		}
		next.ServeHTTP(w, r)
	})
}

func loadChatPage(res http.ResponseWriter, req *http.Request) {
	chatTemplate, _ := template.ParseFiles("caddi/websocket/templates/chat.html")
	chatTemplate.Execute(res, nil)
}

func main() {
	flag.Parse()
	hub := conn.NewHub()
	go hub.Run()
	os.Setenv("GITHUB_KEY", "Ov23lirxru6CJTb3ICG5")
	os.Setenv("GITHUB_SECRET", "af6421654c19c60d200fb610107ee9ea4c58d68e")
	os.Setenv("SESSION_SECRET", "something")
	goth.UseProviders(
		github.New(os.Getenv("GITHUB_KEY"), os.Getenv("GITHUB_SECRET"), "http://localhost:8080/auth/github/callback"),
	)
	key := []byte(os.Getenv("SESSION_SECRET"))
	redisStore := NewRedisStore(key)
	m := map[string]string{
		"github": "Github",
	}
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	router := mux.NewRouter()
	userTemplate, err := template.ParseFiles("caddi/websocket/templates/userprofile.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	router.HandleFunc("/auth/{provider}/callback", func(res http.ResponseWriter, req *http.Request) {
		user, err := gothic.CompleteUserAuth(res, req)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		redisStore.Save(req, res, user)
		userTemplate.Execute(res, user)
	})
	router.HandleFunc("/auth/{provider}", func(res http.ResponseWriter, req *http.Request) {
		//try to get the user without re-authenticating
		if gothUser, err := gothic.CompleteUserAuth(res, req); err == nil {
			redisStore.Save(req, res, gothUser)
			userTemplate.Execute(res, gothUser)
		} else {
			gothic.BeginAuthHandler(res, req)
		}
	})
	router.HandleFunc("/logout/{provider}", func(res http.ResponseWriter, req *http.Request) {
		gothic.Logout(res, req)
		res.Header().Set("Location", "/")
		res.WriteHeader(http.StatusTemporaryRedirect)
	})
	router.Handle("/chat", AuthMiddleware(http.HandlerFunc(loadChatPage)))
	router.Handle("/ws", AuthMiddleware(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		cookie, err := req.Cookie(USERID)
		switch err {
		case nil:
			log.Printf("Received cookie: %s = %s", cookie.Name, cookie.Value)
		case http.ErrNoCookie:
			log.Println("No 'github' received")
			http.Error(res, "No 'github' received", http.StatusInternalServerError)
		default:
			log.Printf("Error getting cookie: %v", err)
			http.Error(res, "Error getting cookie", http.StatusInternalServerError)
		}

		gothUser, err := redisStore.Get(req)
		if err == nil {
			conn.ServeWs(hub, res, req, gothUser)
		} else {
			conn.ServeWs(hub, res, req, &goth.User{})
		}
	})))

	CSRF := csrf.Protect([]byte("32-byte-long-auth-key"), csrf.Secure(false))
	err = http.ListenAndServe(*addr, CSRF(router))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
