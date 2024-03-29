package database

import (
	"fmt"
	"log"
	"os"

	"example.com/main/src/models"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

type DbConnectStruct struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
	sslmode  string
}

func Init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	log.Default().Output(1, "DB Connection started...")
	db_struct := DbConnectStruct{
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		user:     os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASSWORD"),
		dbname:   os.Getenv("DB_NAME"),
		sslmode:  os.Getenv("DB_SSLMODE"),
	}
	str := fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v?sslmode=disable",
		db_struct.user,
		db_struct.password,
		db_struct.host,
		db_struct.port,
		db_struct.dbname,
	)
	db, err := gorm.Open(
		"postgres",
		str,
	)
	if err != nil {
		log.Panicln(err, str)
		panic(err)
	}
	db.AutoMigrate(&models.User{})
	DB = db
	log.Default().Printf("DB Connection successful!")
}
