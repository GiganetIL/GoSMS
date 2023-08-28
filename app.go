package main

import (
	"fmt"
	"os"

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
	atc := NewClient(apiKey)

	balance, err := atc.GetBalance()

	fmt.Printf("Email: %d - %f - %s\n", balance.email.credits, balance.email.percent, balance.email.alert_type)
	fmt.Printf("SMS: %d - %f - %s\n", balance.sms.credits, balance.sms.percent, balance.sms.alert_type)

	os.Exit(0)
	sms_profiles, err := atc.GetSMSProfiles()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	email_profiles, err := atc.GetEmailProfiles()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("SMS Profiles:")
	for _, profile := range sms_profiles {
		fmt.Printf("%s - %d - %v\n", profile.Name, profile.Id, profile.IsDefault)
	}
	fmt.Println("Email Profiles:")
	for _, profile := range email_profiles {
		fmt.Printf("%s - %d - %v\n", profile.ProfileName, profile.Id, profile.IsDefault)
	}
}
