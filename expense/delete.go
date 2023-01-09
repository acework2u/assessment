package expense

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func DelExpense(c echo.Context) error {
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

	id := c.Param("id")
	sqlStmt := "DELETE FROM expenses WHERE id=$1"
	res, err := db.Query(sqlStmt, id)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
		return c.JSON(http.StatusOK, "Delelted ")
	}

	return c.String(http.StatusOK, id+"Deleted")

}
