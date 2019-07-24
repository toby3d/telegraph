package telegraph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPage(t *testing.T) {
	t.Run("invalid", func(t *testing.T) {
		_, err := GetPage("wtf", true)
		assert.Error(t, err)
	})
	t.Run("valid", func(t *testing.T) {
		page, err := GetPage("Sample-Page-12-15", true)
		assert.NoError(t, err)
		assert.NotNil(t, page)
	})
}
