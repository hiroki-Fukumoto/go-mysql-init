package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

var DB *gorm.DB

type DBConfig struct {
	Host       string
	Port       int
	User       string
	DBName     string
	Password   string
	Parameters string
}

func InitDb() {
	dbConfig := DBConfig{
		Host:       os.Getenv("DB_HOST"),
		Port:       os.Getenv("DB_PORT"),
		User:       os.Getenv("DB_USER"),
		Password:   os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		Parameters: "charset=utf8mb4&parseTime=True&loc=Local&allowNativePasswords=true",
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DBName, dbConfig.Parameters)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = db
}
