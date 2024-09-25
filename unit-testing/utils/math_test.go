package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var m = NewMath()

func Test_Tambah(t *testing.T) {
	t.Run("berhasil", func(t *testing.T) {
		result, err := m.Tambah(1, 1)
		assert.NoError(t, err)
		assert.Equal(t, int64(2), result)
	})

	t.Run("gagal", func(t *testing.T) {
		_, err := m.Tambah(-1, 1)
		assert.Error(t, err)
	})
}

func Test_Kurang(t *testing.T) {
	result := m.Kurang(2, 1)
	assert.Equal(t, int64(1), result)
}
