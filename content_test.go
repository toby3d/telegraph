package telegraph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContentFormat(t *testing.T) {
	t.Run("invalid", func(t *testing.T) {
		_, err := ContentFormat(42)
		assert.EqualError(t, ErrInvalidDataType, err.Error())
	})

	t.Run("valid", func(t *testing.T) {
		t.Run("string", func(t *testing.T) {
			validContentDOM, err := ContentFormat(`<p>Hello, World!</p>`)
			assert.NoError(t, err)
			assert.NotEmpty(t, validContentDOM)
		})
		t.Run("bytes", func(t *testing.T) {
			validContentDOM, err := ContentFormat([]byte(`<p>Hello, World!</p>`))
			assert.NoError(t, err)
			assert.NotEmpty(t, validContentDOM)
		})
	})
}
