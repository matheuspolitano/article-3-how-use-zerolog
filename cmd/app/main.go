package main

import (
	"article-3-how-use-zerolog/config"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env.dev", ".env")
	cfg, err := config.LoadDefaultConfig()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(cfg)

}
