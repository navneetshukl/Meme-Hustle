package main

import (
	"log"
	db "memes-hustle/internals/adapters/persistence"

	"github.com/joho/godotenv"
)

func main() {
	err:=godotenv.Load()
	if err!=nil{
		log.Println("Error in loading the .env file")
		return
	}
	_,err=db.ConnectToDB()
	if err!=nil{
		log.Println("Error in connecting to DB ",err)
		return
	}
}