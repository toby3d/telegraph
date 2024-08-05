package telegraph_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"source.toby3d.me/toby3d/telegraph"
)

func TestRevokeAccessToken(t *testing.T) {
	t.Run("invalid", func(t *testing.T) {
		var account telegraph.Account
		_, err := account.RevokeAccessToken()
		assert.Error(t, err)
	})
	t.Run("valid", func(t *testing.T) {
		a, err := telegraph.CreateAccount(telegraph.Account{
			ShortName:  "Sandbox",
			AuthorName: "Anonymous",
		})
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		newAccount, err := a.RevokeAccessToken()
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		assert.NotEqual(t, a.AccessToken, newAccount.AccessToken)
		assert.NotEmpty(t, newAccount.AuthURL)
	})
}