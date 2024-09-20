package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	var t time.Time

	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal(err)
	}

	t = time.Now().In(loc)

	fmt.Println(t)

	strTime := "20/09/2024"
	layout := "02/01/2006" // dd/mm/yyyy
	parsedTime, _ := time.Parse(layout, strTime)

	fmt.Println(parsedTime.Year())
	fmt.Println(parsedTime.Month())
	fmt.Println(parsedTime.Day())

	fmt.Println(parsedTime.Format(time.RFC3339))
}
