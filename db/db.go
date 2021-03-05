package db

import (
	"abc/ent"
	"abc/ent/migrate"
	"context"
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
	"os"
)

var client *ent.Client

func InitDB() {
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", dbHost, username, dbName, password, dbPort) //Build connection string
	fmt.Println("db.go_25", dbUri)
	db, err := sql.Open("pgx", dbUri)
	if err != nil {
		log.Fatal(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	conn := ent.NewClient(ent.Driver(drv))
	client = conn
}

func GetDB() *ent.Client {
	return client.Debug()
}

func AutoMigrate(client *ent.Client) error {
	return client.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true))
}
