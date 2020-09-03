package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Quote struct {
	Message string `json:"quote"`
}

func main() {
	res, err := getQuote()
	if err != nil {
		panic(err)
	}
	fmt.Print(res.Message)
}

func getQuote() (*Quote, error) {
	client := new(http.Client)
	req, err := http.NewRequest("GET", "https://api.kanye.rest", nil)
	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var quote Quote
	err = json.Unmarshal(body, &quote)
	if err != nil {
		return nil, err
	}
	return &quote, nil
}
