package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	goqu "gopkg.in/doug-martin/goqu.v5"
	_ "gopkg.in/doug-martin/goqu.v5/adapters/postgres"
)

func main() {

	host := "192.168.0.16"
	port := 5432
	user := "postgres-dev"
	password := "s3cr3tp4ssw0rd"
	dbname := "dev"
	sslmode := "disable"

	t := "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s"

	connectionString := fmt.Sprintf(t, host, port, user, password, dbname, sslmode)

	pgDb, err := sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Println("Error in postgres connection: ", err)
	}

	err = pgDb.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected!")

	db := goqu.New("postgres", pgDb)

	sql, _, _ := db.From("test").Where(
		goqu.I("col").Eq(10),
		goqu.L(`"json"::TEXT = "other_json"::TEXT`),
	).ToSql()
	fmt.Println(sql)

}
