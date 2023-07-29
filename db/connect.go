package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func Connect() *sql.DB {
	mysqlConnString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	sqlDb, err := sql.Open("mysql", mysqlConnString)

	if err != nil {
		panic(err)
	}

	if _, err := sqlDb.Exec("CREATE TABLE IF NOT EXISTS department(id VARCHAR(225) PRIMARY KEY, department_name VARCHAR(225))"); err != nil {
		panic(err)
	}

	if _, err := sqlDb.Exec("CREATE TABLE IF NOT EXISTS course(id VARCHAR(225) PRIMARY KEY, course_name VARCHAR(225),credit_total INT, semester INT, major_name VARCHAR(225), expectation VARCHAR(225), department_id VARCHAR(225), FOREIGN KEY (department_id) REFERENCES department(id) ON DELETE CASCADE ON UPDATE CASCADE)"); err != nil {
		panic(err)
	}

	if _, err := sqlDb.Exec("CREATE TABLE IF NOT EXISTS major(id VARCHAR(225) PRIMARY KEY, major_name VARCHAR(225), department_id VARCHAR(225), FOREIGN KEY (department_id) REFERENCES department(id) ON DELETE CASCADE ON UPDATE CASCADE)"); err != nil {
		panic(err)
	}
	return sqlDb
}
