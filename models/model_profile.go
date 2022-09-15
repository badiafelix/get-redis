package models

import (
	"api-get-redis/config"
	"fmt"

	"gorm.io/gorm"
)

type Payload struct {
	Usr_name    string `json:"usr_name"`
	Usr_city    string `json:"usr_city"`
	Usr_country string `json:"usr_country"`
}

type Output struct {
	Status  string
	Message string
	Data    Payload
}

func GetProfile() (*Payload, bool, *gorm.DB) {
	var Sel Payload
	var isValid bool
	config.ConnectDB()
	err := config.Db.Raw("SELECT * FROM users LIMIT 1").Scan(&Sel)
	// err := config.Db.Raw("SELECT usr_name,usr_city,usr_country FROM users WHERE usr_name = 'Firman Manurung'").Scan(&Sel)
	// Sel.Name = "Firman Manurung"
	// Sel.City = "Jakarta"
	// Sel.Country = "Indonesia"

	fmt.Println(Sel)
	if err.Error != nil {
		isValid = false
		fmt.Println("errornya", err)
	} else {
		isValid = true
		fmt.Println("asdsdas", &err)
	}
	return &Sel, isValid, err
}
