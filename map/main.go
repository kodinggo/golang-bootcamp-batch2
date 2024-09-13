package main

import (
	"encoding/json"
	"fmt"
)

var isValidType = map[string]bool{
	"JPEG": true,
	"PNG":  true,
}

func main() {
	// var scores map[string]int
	// fmt.Println(scores)

	scores := map[string]int{
		"John":  10,
		"David": 40,
	}
	fmt.Println(scores)
	fmt.Println(scores["John"])

	scores["Bambang"] = 50
	fmt.Println(scores)

	// Pengecak dengan map
	imageType := "GIF"
	if isValidType[imageType] {
		fmt.Println("type is valid")
	}

	// Convert ke JSON
	data := map[string]any{
		"id":      "1",
		"name":    "Eric",
		"hobbies": []string{"swimming", "running"},
		"detail": map[string]string{
			"abc": "efg",
		},
	}
	byteData, _ := json.Marshal(data)
	fmt.Println(string(byteData))
}
