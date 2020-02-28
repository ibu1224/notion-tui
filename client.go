package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Pages struct {
	id   string
	page Page
}

type Page struct {
	id    string `json:"id"`
	title string `json:id`
}

func main() {

	var jsonStr = []byte(`{"space_id":"6be959a0-721f-479e-aa12-77f609a30eab"}`)
	res, err := http.Post("http://127.0.0.1:5000/v1/api/space", "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Fatal(err)
	}

	// header
	fmt.Printf("[status] %d\n", res.StatusCode)
	for k, v := range res.Header {
		fmt.Print("[header] " + k)
		fmt.Println(": " + strings.Join(v, ","))
	}

	data := new(Pages)
	// body
	defer res.Body.Close()
	body, error := ioutil.ReadAll(res.Body)
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(string(body))

	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return
	}

	fmt.Println(data)
}

// resp, err := http.Post("http:127.0.0.1:5000", "application/json", "{"space_id": "6be959a0-721f-479e-aa12-77f609a30eab"}")
// print(resp)
