package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadDB() {
	dbString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?%v",
		ENV.DB_USER,
		ENV.DB_PASSWORD,
		ENV.DB_URL,
		ENV.DB_PORT,
		ENV.DB_NAME,
		"charset=utf8mb4&parseTime=True&loc=Asia%2fJakarta")
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Asia/Jakarta"
	db, err := gorm.Open(mysql.Open(dbString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db

	fmt.Println("Database running successfully")
}
