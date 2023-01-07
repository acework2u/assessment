package expense

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetExpense(c echo.Context) error {

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

	sqlStmt := "SELECT id,title,amount,note,tags FROM expense order by id"

	rows, err := db.Query(sqlStmt)
	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	result := Expenses{}
	// for rows.Next() {
	// 	expense := Expenses{}
	// 	err2 := rows.Scan(&expense.id, &expense.title, &expense.amount, &expense.note, &expense.tags)

	// 	if err2 != nil {
	// 		return err2
	// 	}

	// 	result.Expenses = append(result.Expenses, expense)
	// }
	log.Println(result)
	// return c.JSON(http.StatusCreated, result)
	return c.String(http.StatusOK, "Read")
}
