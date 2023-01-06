package expense

type Expense struct {
	id     int      `json:"id"`
	title  string   `json:"title"`
	amount float32  `json:"amount"`
	note   string   `json:"note"`
	tags   []string `json:tags`
}
