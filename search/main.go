package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
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
	defer pgDb.Close()

	type Article struct {
		Title       string    `db:"title"`
		PublishDate time.Time `db:"publish_date"`
		Summary     string    `db:"summary"`
		TopImage    string    `db:"top_image"`
	}

	//db := goqu.New("postgres", pgDb)

	// 	SELECT *
	// FROM products
	// WHERE to_tsvector('english', product_name) @@ to_tsquery('english', 'hello');
	// product_name := "product_name"
	// sql, _, _ := db.From("products").Where(
	// 	goqu.L("to_tsvector(?, ?) @@ to_query(?, ?)", "\"english\"", product_name, "\"english\"", "hello"),
	// ).ToSql()
	// fmt.Println(sql)
	// productName := "product_name"
	// sql, _, _ := db.From("products").Prepared(true).Where(
	// 	goqu.L("to_tsvector(?, ?) @@ to_query(?, ?)", "\"english\"", productName, "\"english\"", "hello"),
	// ).ScanStructs(&items)

}
