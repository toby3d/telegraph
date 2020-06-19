package telegraph_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/toby3d/telegraph"
)

func TestEditPage(t *testing.T) {
	content, err := telegraph.ContentFormat("<p>Hello, world!</p>")
	assert.NoError(t, err)

	t.Run("invalid", func(t *testing.T) {
		var a telegraph.Account
		_, err := a.EditPage(telegraph.Page{
			Title:      "Sample Page",
			AuthorName: "Anonymous",
			Content:    content,
		}, true)
		assert.Error(t, err)
	})
	t.Run("valid", func(t *testing.T) {
		a := telegraph.Account{
			AccessToken: "b968da509bb76866c35425099bc0989a5ec3b32997d55286c657e6994bbb",
			ShortName:   "Sandbox",
			AuthorName:  "Anonymous",
		}

		page, err := a.EditPage(telegraph.Page{
			Path:       "Sample-Page-12-15",
			Title:      "Sample Page",
			AuthorName: "Anonymous",
			Content:    content,
		}, true)
		assert.NoError(t, err)
		assert.NotEmpty(t, page.Content)
	})
}
