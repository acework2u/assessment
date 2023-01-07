package expense

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func init() {
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

	// Check Table

	//creatTb := `CREATE TABLE IF NOT EXISTS expense (id SERIAL PRIMARY KEY, title varchar(255), amount INT, note varchar(255), tags varchar[] )`
	creatTb := `	CREATE TABLE IF NOT EXISTS expenses (
		id SERIAL PRIMARY KEY,
		title TEXT,
		amount FLOAT,
		note TEXT,
		tags TEXT[]
	);`

	_, err = db.Exec(creatTb)
	if err != nil {
		log.Fatal("Can't not create this table")
	} else {
		log.Println("Create table success")
	}

	defer db.Close()

}
