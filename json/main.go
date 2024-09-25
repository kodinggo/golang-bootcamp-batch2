package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"nama_lengkap,omitempty"`
	Address string `json:"-"`
}

func main() {
	// Encode
	jsonStr := `{"foo":"bar"}`

	var m map[string]string
	json.Unmarshal([]byte(jsonStr), &m)
	fmt.Println(m["foo"])

	// Decode
	pesonMap := map[string]string{
		"nama": "bambang",
	}
	byteJSON, _ := json.Marshal(pesonMap)
	fmt.Println(string(byteJSON))

	// Encode struct
	p := Person{
		Name:    "John",
		Address: "Jakarta",
	}
	byteJSON, _ = json.Marshal(p)
	fmt.Println(string(byteJSON))
}
