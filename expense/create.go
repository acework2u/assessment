package expense

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// type ExpenseInfo struct {
// 	TITLE  string   `json:"title" form:"title"`
// 	AMOUNT float32  `json:"amount" form:"amount`
// 	NOTE   string   `json:"note" form:"note"`
// 	TAGS   []string `json:tags form:"tags"`
// }

type expenseDTO struct {
	Title  string
	Amount float32
	Note   string
	Tags   []string
}

func SaveExpense(c echo.Context) (err error) {
	u := new(ExpenseInfo)
	if err = c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	// load in to struct
	expsRaw := expenseDTO{
		Title:  u.TITLE,
		Amount: u.AMOUNT,
		Note:   u.NOTE,
		Tags:   u.TAGS,
	}

	url := "postgres://yyrrnwpr:Avk9HhlWGrrmGcvGAhtC0nYrAj0JMQA6@john.db.elephantsql.com/yyrrnwpr"
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal("Connect to database err", err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	} else {
		log.Println("Connect database...")
	}

	sqlStmt := "INSERT INTO expense (title,amount,note,tags) VALUES ($1,$2,$3,ARRAY$4)"

	res, err := db.Query(sqlStmt, expsRaw.Title, expsRaw.Amount, expsRaw.Note, expsRaw.Tags)
	// res, err := db.Exec("INSERT INTO expense (title,amount,note,tags) VALUES ('strawberry smoothie',67,'night market promotion discount 10 bath',ARRAY['food', 'beverage'])")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
		return c.JSON(http.StatusCreated, expsRaw)
	}

	return c.JSON(http.StatusOK, "OK")

}
