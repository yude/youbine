package database

import (
	"os"

	"github.com/go-pg/pg"
)

func GetCredentials() *pg.DB {
	db := pg.Connect(&pg.Options{
		User:     GetEnv("POSTGRES_USER", "app"),
		Password: GetEnv("POSTGRES_PASSWORD", "app"),
		Database: GetEnv("POSTGRES_DB", "app"),
	})
	return db
}

func GetEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}
