package telegraph_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"source.toby3d.me/toby3d/telegraph"
)

func TestGetPage(t *testing.T) {
	t.Run("invalid", func(t *testing.T) {
		_, err := telegraph.GetPage("wtf", true)
		assert.Error(t, err)
	})
	t.Run("valid", func(t *testing.T) {
		page, err := telegraph.GetPage("Sample-Page-12-15", true)
		assert.NoError(t, err)
		assert.NotNil(t, page)
	})
}