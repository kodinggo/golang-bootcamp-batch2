package unittesting

import (
	"testing"
	"unit-testing/utils/mocks"

	"github.com/stretchr/testify/assert"
)

func Test_Calculator(t *testing.T) {
	t.Run("success using add operator", func(t *testing.T) {
		mathMock := new(mocks.Math)
		calc := NewCalculator(mathMock)

		mathMock.On("Tambah", int64(1), int64(1)).Return(int64(2), nil).Once()

		result, err := calc.Calculate(1, 1, "+")
		assert.NoError(t, err)
		assert.Equal(t, int64(2), result)
	})

	t.Run("success using subtract operator", func(t *testing.T) {
		mathMock := new(mocks.Math)
		calc := NewCalculator(mathMock)

		mathMock.On("Kurang", int64(4), int64(1)).Return(int64(3)).Once()

		result, err := calc.Calculate(4, 1, "-")
		assert.NoError(t, err)
		assert.Equal(t, int64(3), result)
	})
}
