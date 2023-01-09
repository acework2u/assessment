package expense

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/lib/pq"
	_ "github.com/lib/pq"

	"github.com/labstack/echo/v4"
)

type Expense struct {
	Id     string   `json:"id"`
	Title  string   `json:"title"`
	Amount float32  `json:"amount"`
	Note   string   `json:"note"`
	Tags   TageList `json:tags`
}

type TageList []string

type Expenses []Expense

func GetExpenses(c echo.Context) error {

	//DB Connect
	url := "postgres://yyrrnwpr:Avk9HhlWGrrmGcvGAhtC0nYrAj0JMQA6@john.db.elephantsql.com/yyrrnwpr"
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal("Connect to database err", err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	} else {
		log.Println("Connect database...3")
	}

	sqlStmt := "SELECT * FROM expenses"
	rows, err := db.Query(sqlStmt)
	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	result := Expenses{}
	for rows.Next() {
		var expense Expense
		var tags []string

		//err2 := rows.Scan(&tags)
		err2 := rows.Scan(&expense.Id, &expense.Title, &expense.Amount, &expense.Note, pq.Array(&tags))

		if err2 != nil {
			log.Println(err2)
			return err2
		}

		expense.Tags = tags
		result = append(result, expense)

		//log.Println(result)
	}
	//log.Println(result)
	return c.JSON(http.StatusOK, result)
}

func GetExpense(c echo.Context) error {

	id := c.Param("id")
	//DB Connect
	url := "postgres://yyrrnwpr:Avk9HhlWGrrmGcvGAhtC0nYrAj0JMQA6@john.db.elephantsql.com/yyrrnwpr"
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal("Connect to database err", err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	} else {
		log.Println("Connect database...3")
	}

	sqlStmt := "SELECT * FROM expenses WHERE id = $1"
	rows, err := db.Query(sqlStmt, id)
	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	result := Expenses{}
	for rows.Next() {
		var expense Expense
		var tags []string

		//err2 := rows.Scan(&tags)
		err2 := rows.Scan(&expense.Id, &expense.Title, &expense.Amount, &expense.Note, pq.Array(&tags))

		if err2 != nil {
			log.Println(err2)
			return err2
		}

		expense.Tags = tags
		result = append(result, expense)

		//log.Println(result)
	}
	//log.Println(result)
	return c.JSON(http.StatusOK, result)
}
