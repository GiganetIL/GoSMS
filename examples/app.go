package main

import (
	"fmt"
	"os"

	"github.com/giganetil/gosms"
	"github.com/joho/godotenv"
)

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		fmt.Println("Error loading .env file")
		fmt.Println(envErr)
		os.Exit(1)
	}
	apiKey, _ := os.LookupEnv("AT_APIKEY")
	if apiKey == "" {
		fmt.Println("AT_APIKEY environment variable not set")
		os.Exit(1)
	}
	atc := gosms.NewClient(apiKey)

	balance, err := atc.GetBalance()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if balance.Sms.Credits < 10 {
		fmt.Println("Low balance")
		os.Exit(1)
	}
	atc.SendSMS("Giganet", "0502332060", "Testing Testing 123")
}
