package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	DBhost     = "localhost"
	DBPort     = 5432
	DBuser     = "postgres"
	DBpassword = "postgres"
	DBname     = "libDB"
	DBsslmode  = "disable"
)

var DB *sql.DB
var err error
var ConnStr string

func CheckErr(e error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Initialization() {
	ConnStr = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname =%s sslmode=%s", DBhost, DBPort, DBuser, DBpassword, DBname, DBsslmode)
	DB, err = sql.Open("postgres", ConnStr)

	CheckErr(err)

}

func CloseDB() {
	DB.Close()
}
