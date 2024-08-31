package dao

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewDB() (*sqlx.DB, error) {

	driverName := "postgres"
	dataSourceName := "user=myuser password=mypassword dbname=mydatabase sslmode=disable host=localhost port=5432"

	db, err := sqlx.Connect(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	return db, nil
}
