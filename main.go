package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jlpadilla/benchmark-postgres/pkg/dbclient"
	"github.com/jlpadilla/benchmark-postgres/pkg/fileutil"
)

var database *pgxpool.Pool

func main() {

	database = dbclient.GetConnection()

	fmt.Println("\nStarting setup...")

	processFile("./data/setup.sql", false)

	fmt.Println("Setup done successfully.")

	color.Blue("\nSTARTING BENCHMARKS")

	color.Blue("\nDESCRIPTION: Insert records in the database.\n")
	processDir("./data/insert", false)

	color.Blue("\nDESCRIPTION: Update records in the database.\n")
	processDir("./data/update", true)

	color.Blue("\nDESCRIPTION: Query records in the database.\n")
	processDir("./data/query", true)

	color.Blue("\nDESCRIPTION: Delete records in the database.\n")
	processDir("./data/delete", true)

	color.Blue("\nDONE.")

}

/*****************************
Helper functions
*****************************/

func processDir(directoryName string, print bool) {
	insertFiles := fileutil.GetFilesOnDir(directoryName)
	for _, filename := range insertFiles {
		fmt.Println("\nSCRIPT:", filename)
		processFile(filename, print)
	}
}

func processFile(filename string, print bool) {
	externalSql := fileutil.ReadFile(filename)
	requests := strings.Split(string(externalSql), ";")
	startRoutine := time.Now()

	for _, request := range requests {
		if len(strings.TrimSpace(request)) > 0 {
			res, err := database.Exec(context.Background(), request)
			if err != nil {
				fmt.Println("ERROR:", err)
			}
			if print {
				fmt.Println("RESULT:", res.RowsAffected())
			}
		}
	}

	totalTime := time.Since(startRoutine).Milliseconds()
	fmt.Println("TIME: ", totalTime, "ms")
}
