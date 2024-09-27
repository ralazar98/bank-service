package migration

import (
	"database/sql"
	"flag"
	"github.com/pressly/goose"
	"log"
)

func main() {
	command := flag.String("c", "status", "command")
	dir := flag.String("dir", "./migration", "mgt dir")
	flag.Parse()

	dsn := "postgres://postgres:postgres@localhost:5438"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal(err)
	}

	if err := goose.Run(*command, db, *dir); err != nil {
		log.Fatal(err)
	}
	goose.Up(dsn)
}
