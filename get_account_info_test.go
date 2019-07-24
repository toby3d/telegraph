package telegraph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAccountInfo(t *testing.T) {
	t.Run("invalid", func(t *testing.T) {
		var a Account
		_, err := a.GetAccountInfo()
		assert.Error(t, err)
	})
	t.Run("valid", func(t *testing.T) {
		a := Account{
			AccessToken: "b968da509bb76866c35425099bc0989a5ec3b32997d55286c657e6994bbb",
			ShortName:   "Sandbox",
			AuthorName:  "Anonymous",
		}

		info, err := a.GetAccountInfo(FieldShortName, FieldPageCount)
		assert.NoError(t, err)
		assert.Equal(t, a.ShortName, info.ShortName)
		assert.NotZero(t, info.PageCount)
	})
}
