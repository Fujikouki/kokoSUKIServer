package config

import (
	"github.com/jmoiron/sqlx"
)

func PostgresHost() {
	connStr := "user=myuser password=mypassword dbname=mydatabase sslmode=disable host=localhost port=5432"

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		panic(err)
	}

	defer db.Close()

}
