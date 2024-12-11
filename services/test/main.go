package main

import (
	"flag"

	"pms.pkg/datastore/sqlite"
)

func main() {
	dsn := flag.String("dsn", "./test.db", "")
	flag.Parse()

	_, err := sqlite.Open(*dsn)
	if err != nil {
		print(err.Error())
	}
}
