package telegraph_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/toby3d/telegraph"
)

func TestGetAccountInfo(t *testing.T) {
	t.Run("invalid", func(t *testing.T) {
		var a telegraph.Account
		_, err := a.GetAccountInfo()
		assert.Error(t, err)
	})
	t.Run("valid", func(t *testing.T) {
		a := telegraph.Account{
			AccessToken: "b968da509bb76866c35425099bc0989a5ec3b32997d55286c657e6994bbb",
			ShortName:   "Sandbox",
			AuthorName:  "Anonymous",
		}

		info, err := a.GetAccountInfo(telegraph.FieldShortName, telegraph.FieldPageCount)
		assert.NoError(t, err)
		assert.Equal(t, a.ShortName, info.ShortName)
		assert.NotZero(t, info.PageCount)
	})
}
