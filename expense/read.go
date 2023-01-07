package expense

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

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

type Expenses struct {
	Expenses []Expense `json:"expenses"`
}

func GetExpense(c echo.Context) error {

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

	//sqlStmt := "SELECT id,title,amount,note,tags FROM expense order by id"
	sqlStmt := "SELECT tags FROM expense order by id"

	rows, err := db.Query(sqlStmt)
	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	//result := Expenses{}

	var tagsList TageList

	for rows.Next() {
		//var expense Expense
		var tags string

		err2 := rows.Scan(&tags)
		//err2 := rows.Scan(&expense.Id, &expense.Title, &expense.Amount, &expense.Note, &expense.Tags)

		if err2 != nil {
			log.Println(err2)
			return err2
		}
		//sTag := strings.Split(tags, ",")
		tagsList = []string{tags}
		//result.Expenses = append(result.Expenses, expense)

		log.Println(tagsList)
	}
	//log.Println(result)
	return c.JSON(http.StatusCreated, tagsList)
	//return c.String(http.StatusOK, "OK")
}
