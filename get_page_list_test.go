package telegraph_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"source.toby3d.me/toby3d/telegraph"
)

func TestGetPageList(t *testing.T) {
	t.Run("invalid", func(t *testing.T) {
		var a telegraph.Account
		_, err := a.GetPageList(0, 0)
		assert.Error(t, err)
	})
	t.Run("valid", func(t *testing.T) {
		a := telegraph.Account{
			AccessToken: "b968da509bb76866c35425099bc0989a5ec3b32997d55286c657e6994bbb",
			ShortName:   "Sandbox",
			AuthorName:  "Anonymous",
		}

		list, err := a.GetPageList(1, 1)
		assert.NoError(t, err)
		assert.NotNil(t, list)
	})
}