package main

import (
	"abc/db"
	"context"
	_ "github.com/joho/godotenv/autoload"
)

func init() {
	db.InitDB()
}
func main() {
	ctx := context.Background()
	client := db.GetDB()
	db.AutoMigrate(client)
	client.Test.Create().SetNote("abc").Save(ctx)
}
