package expense

import (
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

	return c.JSON(http.StatusOK, expsRaw)

}
