package main

import (
	"net/http"
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
	e.Logger.Fatal(e.Start(":9090"))
}
