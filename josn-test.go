package main

import "encoding/json"

type Message struct {
	Name string
	Body string
	Time int64
}

func main() {
	b := []byte(`{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`)
	var m Message
	json.Unmarshal(b, &m)
	println(m.Name) // Alice
	println(m.Body) // Hello
	println(m.Time) // 1294706395881547000
}
