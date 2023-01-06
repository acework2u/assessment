package main

import (
	"fmt"

	"github.com/acework2u/assessment/expense"
	"github.com/labstack/echo/v4"
)

var pl = fmt.Println

func main() {
	pl("Please use server.go for main file")
	//DB
	// pl("start at port:", os.Getenv("PORT"))
	e := echo.New()

	//Mid

	//Router
	getExps := expense.GetExpense
	saveExps := expense.SaveExpense
	e.GET("/expenses", getExps)
	e.GET("/expenses/:id", getExps)
	e.POST("/expenses", saveExps)
	e.PUT("/expenses/:id", getExps)
	e.DELETE("/expenses", getExps)

	e.Logger.Fatal(e.Start(":2565"))
}
