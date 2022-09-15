package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectDB() {
	dsn := "host=192.168.18.251 user=postgres password=postgres dbname=docker_db port=2022 sslmode=disable"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}
	Db = conn
}
