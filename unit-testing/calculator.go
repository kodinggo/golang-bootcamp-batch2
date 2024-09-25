package unittesting

import "unit-testing/utils"

type Calculator struct {
	math utils.Math
}

func NewCalculator(math utils.Math) *Calculator {
	return &Calculator{math: math}
}

func (c *Calculator) Calculate(n1, n2 int64, operator string) (result int64, err error) {
	switch operator {
	case "+":
		result, err = c.math.Tambah(n1, n2)
	case "-":
		result = c.math.Kurang(n1, n2)
	case "/":
	case "*":
	}

	return
}
