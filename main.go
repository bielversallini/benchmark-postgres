package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "hippo"
	password = "update-your-postgres-pass-here"
	dbname   = "hippo"
)

// const DB_IN_MEMORY bool = false
const TOTAL_CLUSTERS int = 1 // Number of SNO clusters to simulate.
const PRINT_RESULTS bool = true
const SINGLE_TABLE bool = true // Store relationships in single table or separate table.
const UPDATE_TOTAL int = 1000  // Number of records to update.
const DELETE_TOTAL int = 1000  // Number of records to delete.

var lastUID string

func main() {
	fmt.Printf("Loading %d clusters from template data.\n\n", TOTAL_CLUSTERS)

	// Open the PostgreSQL database.
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", // sslmode=disable",
		host, port, user, password, dbname)
	database, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = database.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

	// Initialize the database tables.
	var nodeStmt, edgeStmt *sql.Stmt
	if SINGLE_TABLE {
		database.Exec("DROP TABLE resources")
		database.Exec("CREATE TABLE IF NOT EXISTS resources (uid TEXT PRIMARY KEY, data TEXT, relatedto TEXT)")
		nodeStmt, _ = database.Prepare("INSERT INTO resources (uid, data, relatedto) VALUES ($1, $2, $3)")
	} else {
		database.Exec("CREATE TABLE IF NOT EXISTS resources (uid TEXT PRIMARY KEY, data TEXT)")
		database.Exec("CREATE TABLE IF NOT EXISTS relationships (sourceId TEXT, destId TEXT)")
		nodeStmt, _ = database.Prepare("INSERT INTO resources (uid, data) VALUES ($1, $2)")
		edgeStmt, _ = database.Prepare("INSERT INTO relationships (sourceId, destId) VALUES ($1, $2)")
	}

	// Load data from file and unmarshall JSON only once.
	addNodes, addEdges := readTemplate()

	// Start counting here to exclude time it takes to read file and unmarshall json
	start := time.Now()

	// LESSON: When using BEGIN and COMMIT TRANSACTION saving to a file is comparable to in memory.
	for i := 0; i < TOTAL_CLUSTERS; i++ {
		// database.Exec("BEGIN TRANSACTION")
		insert(addNodes, nodeStmt, fmt.Sprintf("cluster-%d", i))
		if !SINGLE_TABLE {
			insertEdges(addEdges, edgeStmt, fmt.Sprintf("cluster-%d", i))
		}
		// database.Exec("COMMIT TRANSACTION")
	}

	fmt.Println("\nInsert took", time.Since(start))

	// Benchmark queries
	fmt.Println("\nBENCHMARK QUERIES")

	fmt.Println("\nDESCRIPTION: Find a record using the UID")
	benchmarkQuery(database, fmt.Sprintf("SELECT uid, data FROM resources WHERE uid='%s'", lastUID), true)

	fmt.Println("\nDESCRIPTION: Count all resources")
	benchmarkQuery(database, "SELECT count(*) from resources", true)

	if !SINGLE_TABLE {
		fmt.Println("\nDESCRIPTION: Count all relationships")
		benchmarkQuery(database, "SELECT count(*) FROM relationships", true)
	}

	// fmt.Println("\nDESCRIPTION: Find records with a status name containing `Run`")
	// benchmarkQuery(database, "SELECT uid, data, relatedTo from resources where json_extract(data, \"$.status\") LIKE 'run%' LIMIT 10", PRINT_RESULTS)

	// fmt.Println("\nDESCRIPTION: Find all the values for the field 'namespace'")
	// benchmarkQuery(database, "SELECT DISTINCT json_extract(resources.data, '$.namespace') from resources", PRINT_RESULTS)

	// // LESSON: Adding ORDER BY increases execution time by 2x.
	// fmt.Println("\nDESCRIPTION: Find all the values for the field 'namespace' and sort in ascending order")
	// benchmarkQuery(database, "SELECT DISTINCT json_extract(resources.data, '$.namespace') as namespace from resources ORDER BY namespace ASC", PRINT_RESULTS)

	// fmt.Println("\nDESCRIPTION: Find count of all values for the field 'kind'")
	// benchmarkQuery(database, "SELECT json_extract(resources.data, '$.kind') as kind , count(json_extract(resources.data, '$.kind')) as count FROM resources GROUP BY kind ORDER BY count DESC", PRINT_RESULTS)

	// fmt.Println("\nDESCRIPTION: Find count of all values for the field 'kind' using subquery")
	// benchmarkQuery(database, "SELECT kind, count(*) as count FROM (SELECT json_extract(resources.data, '$.kind') as kind FROM resources) GROUP BY kind ORDER BY count DESC", PRINT_RESULTS)

	// fmt.Println("\nDESCRIPTION: Update a single record.")
	// benchmarkQuery(database, fmt.Sprintf("UPDATE resources SET data = json_set(data, '$.kind', 'value was updated') WHERE uid='%s'", lastUID), true)
	// // Print record to verify it was modified.
	// // benchmarkQuery(database, fmt.Sprintf("SELECT uid, data FROM resources WHERE uid='%s'", lastUID), true)

	fmt.Printf("DESCRIPTION: Update %d records in the database.\n", UPDATE_TOTAL)
	benchmarkUpdate(database, UPDATE_TOTAL)

	fmt.Println("\nDESCRIPTION: Delete a single record.")
	benchmarkQuery(database, fmt.Sprintf("DELETE FROM resources WHERE id='%s'", lastUID), true)
	// Print record to verify it was deleted.
	// benchmarkQuery(database, fmt.Sprintf("SELECT uid, data FROM resources WHERE uid='%s'", lastUID), true)

	fmt.Printf("DESCRIPTION: Delete %d records from the database.\n", DELETE_TOTAL)
	benchmarkDelete(database, DELETE_TOTAL)

	fmt.Println("\nWon't exit so I can get memory usage from OS.")
	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}

/*****************************
Helper functions
*****************************/

/*
 * Read cluster data from file.
 */
func readTemplate() ([]map[string]interface{}, []map[string]interface{}) {
	bytes, _ := ioutil.ReadFile("./data/sno-0.json")
	var data map[string]interface{}
	if err := json.Unmarshal(bytes, &data); err != nil {
		panic(err)
	}
	records := data["addResources"].([]interface{})
	edges := data["addEdges"].([]interface{})

	// Edges format is: { "edgeTypeA": ["destUID1", destUID2], "edgeTypeB": ["destUID3"]}
	findEdges := func(sourceUID string) string {
		result := make(map[string][]string)
		for _, edge := range edges {
			edgeMap := edge.(map[string]interface{})
			if edgeMap["SourceUID"] == sourceUID {
				edgeType := edgeMap["EdgeType"].(string)
				destUIDs, exist := result[edgeType]
				if exist {
					result[edgeType] = append(destUIDs, edgeMap["DestUID"].(string))
				} else {
					result[edgeType] = []string{edgeMap["DestUID"].(string)}
				}
			}
		}
		edgeJSON, _ := json.Marshal(result)
		return string(edgeJSON)
	}

	addResources := make([]map[string]interface{}, len(records))
	for i, record := range records {
		uid := record.(map[string]interface{})["uid"]
		properties := record.(map[string]interface{})["properties"]
		data, _ := json.Marshal(properties)

		e := findEdges(uid.(string))
		// LESSON - QUESTION: UIDs are long and use too much space. What is the risk of compressing?
		// uid = "local-cluster/" + strings.Split(uid.(string), "-")[5]
		addResources[i] = map[string]interface{}{"uid": uid, "data": string(data), "edges": e}
	}

	addEdges := make([]map[string]interface{}, len(edges))
	for i, edge := range edges {
		t := edge.(map[string]interface{})["EdgeType"]
		s := edge.(map[string]interface{})["SourceUID"]
		d := edge.(map[string]interface{})["DestUID"]
		addEdges[i] = map[string]interface{}{"sourceUID": s, "destUID": d, "type": t}
	}
	return addResources, addEdges
}

/*
 * Insert records
 */
func insert(records []map[string]interface{}, statement *sql.Stmt, clusterName string) {
	fmt.Print(".")
	for i, record := range records {
		lastUID = strings.Replace(record["uid"].(string), "local-cluster", clusterName, 1)
		var err error
		if SINGLE_TABLE {
			edges := record["edges"].(string)
			edges = strings.ReplaceAll(edges, "local-cluster", clusterName)
			_, err = statement.Exec(lastUID, record["data"], edges)
			fmt.Printf("Inserting is slow %d of %d\n", i, len(records))
		} else {
			_, err = statement.Exec(lastUID, record["data"])
		}
		if err != nil {
			fmt.Println("Error inserting record:", err, statement)
			panic(err)
		}
	}
}

/*
 * Insert edges in separate table.
 */
func insertEdges(edges []map[string]interface{}, statement *sql.Stmt, clusterName string) {
	fmt.Print(">")
	for _, edge := range edges {
		source := strings.Replace(edge["sourceUID"].(string), "local-cluster", clusterName, 1)
		dest := strings.Replace(edge["destUID"].(string), "local-cluster", clusterName, 1)
		_, err := statement.Exec(source, dest)

		if err != nil {
			fmt.Println("Error inserting edge:", err)
		}
	}
}

func benchmarkQuery(database *sql.DB, q string, printResult bool) {
	startQuery := time.Now()
	rows, queryError := database.Query(q)
	defer rows.Close()
	if queryError != nil {
		fmt.Println("Error executing query: ", queryError)
	}

	fmt.Println("QUERY      :", q)
	if printResult {
		fmt.Println("RESULTS    :")
	} else {
		fmt.Println("RESULTS    : To print results set PRINT_RESULTS=true")
	}

	for rows.Next() {
		columns, _ := rows.Columns()
		columnData := make([]string, 3)
		switch len(columns) {
		case 3:
			rows.Scan(&columnData[0], &columnData[1], &columnData[2])
		case 2:
			rows.Scan(&columnData[0], &columnData[1])
		default:
			rows.Scan(&columnData[0])
		}

		if printResult {
			fmt.Println("  *\t", columnData[0], columnData[1], columnData[2])
		}
	}
	// LESSON: We can stream results from rows, but using aggregation and sorting will delay results because we have to process al records first.
	fmt.Printf("TIME       : %v \n\n", time.Since(startQuery))
}

/*
 * Helper method to select records for Update and Delete.
 */
func selectRandomRecords(database *sql.DB, total int) []string {
	records, _ := database.Query("SELECT id FROM resources ORDER BY RANDOM() LIMIT ?", total)
	uids := make([]string, total)
	for i := 0; records.Next(); i++ {
		scanErr := records.Scan(&uids[i])
		if scanErr != nil {
			fmt.Println(scanErr)
		}
	}
	return uids
}

/*
 * Benchmark UPDATE
 */
func benchmarkUpdate(database *sql.DB, updateTotal int) {
	// First, let's find some random records to update.
	uids := selectRandomRecords(database, updateTotal)

	// Now that we have the UIDs we want to update, start benchmarking from here.
	start := time.Now()
	updateStmt, _ := database.Prepare("UPDATE resources SET data = json_set(data, '$.kind', 'Updated value') WHERE id = ?")
	defer updateStmt.Close()
	// Lesson: Using BEGIN/COMMIT TRANSACTION doesn't seem to affect performance.
	for _, uid := range uids {
		updateStmt.Exec(uid)
	}

	fmt.Printf("QUERY      : UPDATE resources SET data = json_set(data, '$.kind', 'Updated value') WHERE id = ? \n")
	fmt.Printf("TIME       : %v \n\n", time.Since(start))
}

/*
 * Benchmark DELETE
 */
func benchmarkDelete(database *sql.DB, deleteTotal int) {
	// First, let's find some random records to delete.
	uids := selectRandomRecords(database, deleteTotal)

	// Now that we have the UIDs we want to delete, start benchmarking from here.
	start := time.Now()
	database.Exec("DELETE from resources WHERE id IN (?)", strings.Join(uids, ", "))

	fmt.Printf("QUERY      : DELETE from resources WHERE id IN (?) \n") //, strings.Join(uids, ", "))
	fmt.Printf("TIME       : %v \n\n", time.Since(start))
}
