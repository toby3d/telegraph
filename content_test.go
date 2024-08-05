package telegraph_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"source.toby3d.me/toby3d/telegraph"
)

func TestContentFormat(t *testing.T) {
	t.Run("invalid", func(t *testing.T) {
		_, err := telegraph.ContentFormat(42)
		assert.EqualError(t, telegraph.ErrInvalidDataType, err.Error())
	})

	t.Run("valid", func(t *testing.T) {
		t.Run("string", func(t *testing.T) {
			validContentDOM, err := telegraph.ContentFormat(`<p>Hello, World!</p>`)
			assert.NoError(t, err)
			assert.NotEmpty(t, validContentDOM)
		})
		t.Run("bytes", func(t *testing.T) {
			validContentDOM, err := telegraph.ContentFormat([]byte(`<p>Hello, World!</p>`))
			assert.NoError(t, err)
			assert.NotEmpty(t, validContentDOM)
		})
	})
}