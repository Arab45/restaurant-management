package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	// Load .env early so package-level vars can read environment values.
	godotenv.Load()
}

var SECRET_KEY = func() string {
	key := os.Getenv("SECRET_KEY")
	if key == "" {
		log.Fatal("SECRET_KEY not set in environment")
	}
	return key
}()

// DOCS_ENABLED controls whether the static docs are served by the server.
// Set to "true" or "1" to enable. Defaults to false.
var DOCS_ENABLED = func() bool {
	v := os.Getenv("DOCS_ENABLED")
	return v == "true" || v == "1"
}()

// DOCS_PATH is the filesystem path to the generated docs folder.
// Defaults to `docs` in the project root.
var DOCS_PATH = func() string {
	p := os.Getenv("DOCS_PATH")
	if p == "" {
		return "docs"
	}
	return p
}()
