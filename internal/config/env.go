package config

import (
	"log"
	"os"
)

var SECRET_KEY = func() string {
	key := os.Getenv("SECRET_KEY")
	if key == "" {
		log.Fatal("SECRET_KEY not set in environment")
	}
	return key
}()
