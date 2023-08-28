package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type EmailProfile struct {
	Id              int    `json:"id"`
	ProfileName     string `json:"profile_name"`
	FromName        string `json:"FromName"`
	FromEmail       string `json:"from_email"`
	UpdateDate      string `json:"update_date"`
	RelyToEmail     string `json:"reply_to_email"`
	SubjectPostfix  string `json:"subject_postfix"`
	IsUseDomainKeys bool   `json:"is_use_domain_keys"`
	IsDefault       bool   `json:"is_default"`
}

type SMSProfile struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	FromNumber string `json:"fromNumber"`
	UpdatedOn  string `json:"updatedOn"`
	IsDefault  bool   `json:"isDefault"`
}

type AccountBalance struct {
	email Balance
	sms   Balance
}

type Balance struct {
	credits    int
	percent    float32
	alert_type string
}

type ATClient struct {
	APIKey  string
	BaseURL string
}

func NewClient(apiKey string) *ATClient {
	return &ATClient{
		APIKey:  apiKey,
		BaseURL: "https://webapi.mymarketing.co.il/api",
	}
}

func (c *ATClient) GetBalance() (AccountBalance, error) {
	resp, err := c.GET("account/balance")
	if err != nil {
		return AccountBalance{}, err
	}
	fmt.Println(resp)
	balance := AccountBalance{}
	err = json.Unmarshal([]byte(resp), &balance)
	return balance, err
}

func (c *ATClient) GetEmailProfiles() ([]EmailProfile, error) {
	resp, err := c.GET("account/sendingprofiles")
	fmt.Println(resp)
	if err != nil {
		return nil, err
	}
	profiles := []EmailProfile{}
	err = json.Unmarshal([]byte(resp), &profiles)
	return profiles, err
}
func (c *ATClient) GetSMSProfiles() ([]SMSProfile, error) {
	resp, err := c.GET("account/sms-sendingprofiles")
	if err != nil {
		return nil, err
	}
	profiles := []SMSProfile{}
	err = json.Unmarshal([]byte(resp), &profiles)
	return profiles, err
}

func (c *ATClient) GET(url string) (string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", c.BaseURL, url), nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", c.APIKey)
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	respBody, err := io.ReadAll(resp.Body)
	return string(respBody), err
}
func (c *ATClient) POST(url string, data interface{}) (string, error) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", c.BaseURL, url), bytes.NewReader(dataBytes))
	if err != nil {
		return "", err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", c.APIKey)

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	respBody, err := io.ReadAll(resp.Body)
	return string(respBody), err
}
