package database

import (
	"fmt"
	"go_news/model/entity"
	"log"
)

func RunMigration() {
	err := DB.AutoMigrate(&entity.User{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Migration Success")
}
