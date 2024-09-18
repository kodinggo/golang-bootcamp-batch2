package main

import (
	"fmt"
	"log"
)

// Custom error khusus untuk error 404
type ErrorNotFound interface {
	Error() string
}

type errorNotFound struct {
	Msg string
}

func (e errorNotFound) Error() string {
	return e.Msg
}

func NewErrorNotFound(msg string) ErrorNotFound {
	return &errorNotFound{Msg: msg}
}

func main() {
	isValid, err := cekUmur(15)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("error: ", err.Error())
	fmt.Println(isValid)

	switch err.(type) {
	case ErrorNotFound:
		// tentukan HTTP status code
	}
}

func cekUmur(umur int64) (bool, error) {
	if umur < 17 {
		return false, NewErrorNotFound("dibawah umur") // error
	}
	return true, nil // tidak error
}
