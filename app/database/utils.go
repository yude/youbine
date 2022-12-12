package database

import "github.com/go-pg/pg"

func GetCredentials() *pg.DB {
	db := pg.Connect(&pg.Options{
		User:     "app",
		Password: "app",
		Database: "app",
	})
	return db
}
