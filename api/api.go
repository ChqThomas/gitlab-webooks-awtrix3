package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Client struct {
    baseUrl string
	username string
	password string
}

func NewClient(baseUrl string, username string, password string) *Client {
    return &Client{
        baseUrl: baseUrl,
		username: username,
		password: password,
    }
}

func (c *Client) PostRequest(path string, data interface{}) {
	body, _ := json.Marshal(data)
	bodyString := bytes.NewBuffer(body)
	fmt.Println(bodyString)
	req, err := http.NewRequest("POST", c.baseUrl + path, bodyString)

	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(c.username, c.password)

	client := &http.Client{}
	response, err := client.Do(req)
	
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	defer response.Body.Close()

	fmt.Println("Status:", response.Status)
	fmt.Println("Response:", response.Body)
}