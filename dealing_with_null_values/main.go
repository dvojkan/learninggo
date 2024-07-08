package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "test"
)

func main() {
	//PostgreSQL connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, age, score, active, created_at FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		//this construction is used to define several variables in one block. It is the same as defining 5 variables one by one.
		var (
			name sql.NullString
			age  sql.NullInt64
			//age       sql.NullFloat64 - error
			score     sql.NullFloat64
			active    sql.NullBool
			createdAt sql.NullTime
		)
		// sql.Null types must correspond appropriate database types:
		// SQL Type: VARCHAR, TEXT, etc. -> Go Type: sql.NullString
		// SQL Type: INT, BIGINT, etc. -> Go Type: sql.NullInt64
		// SQL Type: FLOAT, DOUBLE, REAL, etc. -> Go Type: sql.NullFloat64
		// SQL Type: BOOLEAN -> Go Type: sql.NullBool
		// SQL Type: TIMESTAMP, DATETIME, DATE -> Go Type: sql.NullTime

		if err := rows.Scan(&name, &age, &score, &active, &createdAt); err != nil {
			log.Fatal(err)
		}

		if name.Valid {
			fmt.Println("Name:", name.String)
		} else {
			fmt.Println("Name is NULL")
		}

		if age.Valid {
			fmt.Println("Age:", age.Int64)
		} else {
			fmt.Println("Age is NULL")
		}

		if score.Valid {
			fmt.Println("Score:", score.Float64)
		} else {
			fmt.Println("Score is NULL")
		}

		if active.Valid {
			fmt.Println("Active:", active.Bool)
		} else {
			fmt.Println("Active is NULL")
		}

		if createdAt.Valid {
			fmt.Println("Created At:", createdAt.Time)
		} else {
			fmt.Println("Created At is NULL")
		}
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
