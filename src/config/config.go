package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load(".env.local")
	if err != nil {
		err := godotenv.Load() 
		if err != nil {
			fmt.Println(err)
			panic("Error loading .env file")
		}
	}
	
}

func GetPort() string {
	port, err := os.LookupEnv("GO_PORT")
	if !err {
		fmt.Println("Error get value from .env file")
	}
	return ":" + port
}