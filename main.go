package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jlpadilla/benchmark-postgres/pkg/dbclient"
)

func main() {

	database := dbclient.GetConnection()

	fmt.Println("\nStarting setup...")
	process(database, "./data/setup.sql", false)
	fmt.Println("Setup done successfully.")

	fmt.Println("\nSTARTING BENCHMARKS")

	fmt.Printf("\nDESCRIPTION: Insert records in the database.\n")
	process(database, "./data/insert/1.sql", true)
	process(database, "./data/insert/2.sql", true)

	fmt.Println("\nDONE.")
}

/*****************************
Helper functions
*****************************/

func process(database *pgxpool.Pool, filename string, print bool) {
	externalSql := readSQLFile(filename)
	requests := strings.Split(string(externalSql), ";")

	startRoutine := time.Now()

	for _, request := range requests {
		_, err := database.Exec(context.Background(), request)
		if err != nil {
			fmt.Println("ERROR    : ", err)
		}
	}

	if print {
		fmt.Printf("TOTAL TIME : %vms \n", time.Since(startRoutine).Milliseconds())
	}
}

func readSQLFile(filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}
