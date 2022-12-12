package database

import (
	"time"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type Message struct {
	Id       int64
	Content  string
	IpAddr   string
	DateTime time.Time
}

func create_schema(db *pg.DB) error {
	models := []interface{}{
		(*Message)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp: false,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
