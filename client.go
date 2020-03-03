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

func getPages() []interface{} {

	var jsonStr = []byte(`{"space_id":"6be959a0-721f-479e-aa12-77f609a30eab"}`)
	res, err := http.Post("http://127.0.0.1:5000/v1/api/pages", "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Fatal(err)
	}

	// header
	fmt.Printf("[status] %d\n", res.StatusCode)
	for k, v := range res.Header {
		fmt.Print("[header] " + k)
		fmt.Println(": " + strings.Join(v, ","))
	}

	// body
	defer res.Body.Close()
	body, error := ioutil.ReadAll(res.Body)
	if error != nil {
		log.Fatal(error)
	}

	var decoded []interface{}
	if err := json.Unmarshal(body, &decoded); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
	}

	return decoded
	// s := decoded[1].(map[string]interface{})["title"].(string)
	// for _, s := range decoded {
	// fmt.Println(s.(map[string]interface{})["title"].(string))
	// }

}

// resp, err := http.Post("http:127.0.0.1:5000", "application/json", "{"space_id": "6be959a0-721f-479e-aa12-77f609a30eab"}")
// print(resp)
