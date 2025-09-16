package main

import (
	"fmt"
	. "home/caddi/websocket/conn"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/markbates/goth"
)

var hub = NewHub()

func complexHandler(w http.ResponseWriter, r *http.Request) {
	ServeWs(hub, w, r, &goth.User{})

}

func client1ConnectToWs(wait1 chan bool) {
	s := httptest.NewServer(http.HandlerFunc(complexHandler))
	defer s.Close()
	u := "ws" + strings.TrimPrefix(s.URL, "http")
	conn, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()
	if err := conn.WriteMessage(websocket.TextMessage, []byte("hello")); err != nil {
		fmt.Println(err.Error())
		return
	}
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf("connect1 receive: %v\n", string(p))
		if err := conn.WriteMessage(websocket.TextMessage, []byte("exit")); err != nil {
			fmt.Println(err.Error())
			return
		}
		_, p, err = conn.ReadMessage()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf("connect1 receive: %v\n", string(p))
		if strings.Contains(string(p), "exit") {
			break
		}
	}
	wait1 <- true
}

func client2ConnectToWs(wait2 chan bool) {
	s := httptest.NewServer(http.HandlerFunc(complexHandler))
	defer s.Close()
	u := "ws" + strings.TrimPrefix(s.URL, "http")
	conn, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf("connect2 receive: %v\n", string(p))
		if strings.Contains(string(p), "exit") {
			break
		}
	}
	wait2 <- true
}
func TestComplexWebsocket(t *testing.T) {
	go hub.Run()
	wait1 := make(chan bool)
	wait2 := make(chan bool)
	go client1ConnectToWs(wait1)
	go client2ConnectToWs(wait2)
	<-wait1
	<-wait2
}
