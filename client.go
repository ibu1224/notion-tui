package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func getPages() []interface{} {
	var jsonStr = []byte(`{"space_id":""}`)
	res, err := http.Post("http://127.0.0.1:5000/v1/api/pages", "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Fatal(err)
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
}

func getPageContent(pageID string) []interface{} {
	var jsonStr = []byte(`{"page_id":`)
	jsonStr = append(jsonStr, `"`...)
	jsonStr = append(jsonStr, pageID...)
	jsonStr = append(jsonStr, `"`...)
	jsonStr = append(jsonStr, "}"...)

	res, err := http.Post("http://127.0.0.1:5000/v1/api/page/block", "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Fatal(err)
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
}

func createPageContent(pageID string, title string) {
	// var jsonStr = []byte(`{"page_id":"e4376ca0-5e61-401b-8fd2-878e2e6e68e4","title":"hoge"}`)
	titleKey := []byte(`"title":`)
	var jsonStr = []byte(`{"page_id":`)
	jsonStr = append(jsonStr, `"`...)
	jsonStr = append(jsonStr, pageID...)
	jsonStr = append(jsonStr, `"`...)
	jsonStr = append(jsonStr, `,`...)
	jsonStr = append(jsonStr, titleKey...)
	jsonStr = append(jsonStr, `"`...)
	jsonStr = append(jsonStr, title...)
	jsonStr = append(jsonStr, `"`...)
	jsonStr = append(jsonStr, "}"...)

	res, err := http.Post("http://127.0.0.1:5000/v1/api/page/block/new", "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Fatal(err)
	}
	// body
	defer res.Body.Close()
	// body, error := ioutil.ReadAll(res.Body)
	// if error != nil {
	// 	log.Fatal(error)
	// }
}
