package telegraph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePage(t *testing.T) {
	content, err := ContentFormat(`<p>Hello, world!</p>`)
	assert.NoError(t, err)

	t.Run("invalid", func(t *testing.T) {
		var a Account
		_, err := a.CreatePage(Page{
			Title:      "Sample Page",
			AuthorName: "Anonymous",
			Content:    content,
		}, true)
		assert.Error(t, err)
	})
	t.Run("valid", func(t *testing.T) {
		a := Account{
			AccessToken: "b968da509bb76866c35425099bc0989a5ec3b32997d55286c657e6994bbb",
			ShortName:   "Sandbox",
			AuthorName:  "Anonymous",
		}

		page, err := a.CreatePage(Page{
			Title:      "Sample Page",
			AuthorName: "Anonymous",
			Content:    content,
		}, true)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		assert.NotEmpty(t, page.URL)
	})
}
