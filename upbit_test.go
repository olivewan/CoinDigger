package upbit

import (
	"log"
	"os"
	"testing"
	"github.com/joho/godotenv"
)

// fix me
var (
	accessKey string
	secretKey string

	marketID = "KRW-SOL"
	currency = "KRW"
)

func TestMain(m *testing.M) {
	err := godotenv.Load()
	
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	os.Exit(m.Run())
}
