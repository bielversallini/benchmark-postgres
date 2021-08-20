package dbclient

import (
	"context"
	"fmt"
	"time"
)

func BenchmarkQuery(q string, printResult bool) {
	startQuery := time.Now()
	rows, err := pool.Query(context.Background(), q)
	defer rows.Close()
	if err != nil {
		fmt.Println("Error executing query: ", err)
	}

	fmt.Println("QUERY      :", q)
	if printResult {
		fmt.Println("RESULTS    :")
	} else {
		fmt.Println("RESULTS    : To print results set PRINT_RESULTS=true")
	}

	for rows.Next() {
		rowValues, _ := rows.Values()
		fmt.Println(rowValues)
	}

	fmt.Printf("TIME       : %v \n\n", time.Since(startQuery))
}
