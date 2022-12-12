package database

import (
	"fmt"
	"log"
	"time"
)

func Init() {
	fmt.Println("aaa")
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
