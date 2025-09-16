package main

import (
	"database/sql"
	"fmt"
	. "home/bamboo/server/handler"
	. "home/bamboo/server/service"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func servePingRequest(c echo.Context) error {
	return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
}

func initStore() (*sql.DB, error) {
	pgConnString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGDATABASE"),
		os.Getenv("PGUSER"),
	)
	openDB := func() (*sql.DB, error) {
		db, err := sql.Open("postgres", pgConnString)
		return db, err
	}
	db, err := openDB()
	if err != nil {
		return nil, err
	}
	if _, err := db.Exec(
		"create table if not exists message (value text primary key)"); err != nil {
		return nil, err
	}
	return db, nil
}

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	db, err := initStore()
	if err != nil {
		log.Fatal("Failed to initialize store:", err)
	}
	messageService := NewCockroachDBMessageService(db)
	defer messageService.Close()
	messageHandler := NewMessageHandler(messageService)
	e.GET("/", messageHandler.CountMessages)
	e.POST("/send", messageHandler.CreateMessage)
	e.GET("/ping", servePingRequest)

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "9090"
	}
	e.Logger.Fatal(e.Start(":" + httpPort))
}
