package dbclient

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	pgxpool "github.com/jackc/pgx/v4/pgxpool"
)

const maxConnections = 8

var pool *pgxpool.Pool
var InsertChan chan *Record

type Record struct {
	UID        string
	Cluster    string
	Name       string
	Properties map[string]interface{}
}

func init() {
	InsertChan = make(chan *Record, 100)
	createPool()

	// Start go routines to process insert.
	go batchInsert("A")
	go batchInsert("B")
}

// Initializes the connection pool.
func createPool() {
	DB_USER := getEnvOrUseDefault("POSTGRES_USER", "")
	DB_PASSWORD := url.QueryEscape(getEnvOrUseDefault("POSTGRES_PASSWORD", ""))
	DB_HOST := getEnvOrUseDefault("POSTGRES_HOST", "")

	DB_PORT, err := strconv.Atoi(getEnvOrUseDefault("POSTGRES_PORT", ""))
	if err != nil {
		DB_PORT = 5432
		fmt.Println("Error parsing db port, using default 5432")
	}

	DB_NAME := getEnvOrUseDefault("POSTGRES_DB", "")

	database_url := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)

	time.Sleep(8 * time.Second)

	fmt.Println("Connecting to PostgreSQL at: ", strings.ReplaceAll(database_url, DB_PASSWORD, "*****"))
	config, connerr := pgxpool.ParseConfig(database_url)

	if connerr != nil {
		fmt.Println("Error connecting to DB:", connerr)
	}

	config.MaxConns = maxConnections
	conn, err := pgxpool.ConnectConfig(context.Background(), config)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	pool = conn
}

func GetConnection() *pgxpool.Pool {
	err := pool.Ping(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	return pool
}

func getEnvOrUseDefault(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
