package handler

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv()  {
	err:=godotenv.Load()
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal("Error get .env")
	}
}