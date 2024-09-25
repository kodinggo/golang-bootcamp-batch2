package unittesting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Tambah(t *testing.T) {
	t.Run("berhasil", func(t *testing.T) {
		result, err := Tambah(1, 1)
		assert.NoError(t, err)
		assert.Equal(t, 2, result)
	})

	t.Run("gagal", func(t *testing.T) {
		_, err := Tambah(-1, 1)
		assert.Error(t, err)
	})
}

func Test_Kurang(t *testing.T) {
	result := Kurang(2, 1)
	assert.Equal(t, 1, result)
}
