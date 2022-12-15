package database

import (
	"github.com/go-pg/pg"
	"github.com/yude/youbine/utils"
)

func GetCredentials() *pg.DB {
	db := pg.Connect(&pg.Options{
		User:     utils.GetEnv("POSTGRES_USER", "app"),
		Password: utils.GetEnv("POSTGRES_PASSWORD", "app"),
		Database: utils.GetEnv("POSTGRES_DB", "app"),
	})
	return db
}
