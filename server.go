package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

var pl = fmt.Println

func main() {
	pl("Please use server.go for main file")
	// pl("start at port:", os.Getenv("PORT"))
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":2526"))
}
