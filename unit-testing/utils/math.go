package utils

import "errors"

type Math interface {
	Tambah(a, b int64) (int64, error)
	Kurang(a, b int64) int64
}

type math struct{}

func NewMath() Math {
	return &math{}
}

func (m *math) Tambah(a, b int64) (int64, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("number can't be negative")
	}
	return a + b, nil
}

func (m *math) Kurang(a, b int64) int64 {
	return a - b
}
