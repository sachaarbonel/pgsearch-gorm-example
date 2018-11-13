package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	_ "gopkg.in/doug-martin/goqu.v5/adapters/postgres"
)

func main() {
	// pgDb, err := sql.Open("postgres", "user=postgres dbname=goqupostgres sslmode=disable ")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// db := goqu.New("postgres", pgDb)
	// //interpolated sql
	// sql, _ := db.From("user").Where(goqu.Ex{
	// 	"id": 10,
	// }).ToSql()
	// fmt.Println(sql)

	host := "192.168.0.16"
	port := 5432
	user := "postgres-dev"
	password := "s3cr3tp4ssw0rd"
	dbname := "dev"
	sslmode := "disable"

	t := "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s"

	connectionString := fmt.Sprintf(t, host, port, user, password, dbname, sslmode)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Println("Error in postgres connection: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected!")

}
