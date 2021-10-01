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

	fmt.Println("QUERY      : ", q)
	fmt.Println("ERROR      : ", err)

	if printResult {
		fmt.Println("RESULTS    :")
	}

	for rows.Next() {
		rowValues, _ := rows.Values()
		fmt.Println(rowValues)
	}

	fmt.Printf("TIME       : %v \n\n", time.Since(startQuery).Seconds())
}
