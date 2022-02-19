package config

import "github.com/joho/godotenv"

func InitConfig() {
	err := godotenv.Load(".env")
	if err != nil {

	}
}
