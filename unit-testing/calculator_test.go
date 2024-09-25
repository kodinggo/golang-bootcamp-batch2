package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Tambah(t *testing.T) {
	result := Tambah(1, 1)
	assert.Equal(t, 3, result)
}
