package expense

type ExpenseInfo struct {
	ID     string   `json:"id"`
	TITLE  string   `json:"title"`
	AMOUNT float32  `json:"amount"`
	NOTE   string   `json:"note"`
	TAGS   []string `json:tags`
}

type Expenses struct {
	Expenses []ExpenseInfo `json:"Expenses"`
}
