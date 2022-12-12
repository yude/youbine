package database

import (
	"fmt"
	"log"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type Message struct {
	Id      int64
	Content string
	IpAddr  string
}

func Init() {
	fmt.Println("aaa")
	db := pg.Connect(&pg.Options{
		User:     "app",
		Password: "app",
		Database: "app",
	})
	defer db.Close()

	err := create_schema(db)
	if err != nil {
		fmt.Println(err)
	}
}

func AddMessage(content string, ip_addr string) {
	db := pg.Connect(&pg.Options{
		User:     "app",
		Password: "app",
		Database: "app",
	})
	defer db.Close()

	message := &Message{
		Content: content,
		IpAddr:  ip_addr,
	}
	_, err := db.Model(message).Insert()
	if err != nil {
		log.Fatal(err)
	}
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

func get_credentials() *pg.DB {
	db := pg.Connect(&pg.Options{
		User:     "app",
		Password: "app",
		Database: "app",
	})
	return db
}
