package database

import (
	"fmt"
	"log"
	"time"
)

func Init() {
	db := GetCredentials()
	defer db.Close()

	err := create_schema(db)
	if err != nil {
		fmt.Println(err)
	}
}

func AddMessage(content string, ip_addr string) {
	db := GetCredentials()
	defer db.Close()

	message := &Message{
		Content:  content,
		IpAddr:   ip_addr,
		DateTime: time.Now(),
	}
	_, err := db.Model(message).Insert()
	if err != nil {
		log.Fatal(err)
	}
}

func ReturnMessage() []Message {
	db := GetCredentials()
	defer db.Close()

	var messages []Message
	err := db.Model(&messages).Select()
	if err != nil {
		log.Fatal(err)
	}

	return messages
}
