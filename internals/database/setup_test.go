package database

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	if err := godotenv.Load("./../../.env"); err != nil {
		log.Fatalf("cannot load the env files: %v", err)
	}
	os.Exit(m.Run())
}
