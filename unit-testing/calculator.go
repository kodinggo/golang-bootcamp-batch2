package unittesting

import "errors"

func Tambah(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("number can't be negative")
	}
	return a + b, nil
}

func Kurang(a, b int) int {
	return a - b
}
