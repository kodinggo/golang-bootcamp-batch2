package unittesting

import (
	"testing"
	"unit-testing/utils"

	"github.com/stretchr/testify/assert"
)

func Test_Calculator(t *testing.T) {
	t.Run("success using add operator", func(t *testing.T) {
		math := utils.NewMath()
		calc := NewCalculator(math)

		result, err := calc.Calculate(1, 1, "+")
		assert.NoError(t, err)
		assert.Equal(t, int64(2), result)
	})
}
