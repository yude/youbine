package database

import (
	"fmt"
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
		var exists bool
		_, err := db.QueryOne(pg.Scan(&exists), `
        SELECT EXISTS (
            SELECT FROM information_schema.tables 
            WHERE table_schema = 'public' 
            AND table_name = ?
        )`, "messages")

		if err != nil {
			fmt.Println(err)
		}

		if exists {
			break
		}

		err = db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp: false,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
