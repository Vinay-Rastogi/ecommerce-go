package config

import (
    "database/sql"
    "log"

    _ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {

   connStr := "user=koushik dbname=ecommerce host=localhost port=5432 sslmode=disable"

    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    if err := db.Ping(); err != nil {
        log.Fatal("Failed to ping database:", err)
    }

    log.Println("Database connected successfully")

    return db
}
