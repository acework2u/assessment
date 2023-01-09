package expense

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

func PutExpense(c echo.Context) error {

	// DB Connect
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

	expsInfo := new(Expense)
	if err2 := c.Bind(expsInfo); err != nil {
		return err2
	}

	sqlStmt := "UPDATE expenses SET title=$1,amount=$2,note=$3,tags=$4 WHERE id=$5"

	res, err := db.Query(sqlStmt, expsInfo.Title, expsInfo.Amount, expsInfo.Note, pq.Array(expsInfo.Tags), expsInfo.Id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
		return c.JSON(http.StatusCreated, expsInfo)
	}
	return c.String(http.StatusOK, expsInfo.Id)

}
