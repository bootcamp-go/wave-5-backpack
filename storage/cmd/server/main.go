package main

import (
	"github.com/joho/godotenv"
	cnx "wave-5-backpack/storage/db"
)

func main() {
	loadEnv()

	cnx.MySQLConnection()

}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}