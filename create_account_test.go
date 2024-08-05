package telegraph_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"source.toby3d.me/toby3d/telegraph"
)

func TestCreateAccount(t *testing.T) {
	t.Run("invalid", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			_, err := telegraph.CreateAccount(telegraph.Account{})
			assert.Error(t, err)
		})
		t.Run("without shortname", func(t *testing.T) {
			_, err := telegraph.CreateAccount(telegraph.Account{
				ShortName:  "",
				AuthorName: "Anonymous",
			})
			assert.Error(t, err)
		})
	})
	t.Run("valid", func(t *testing.T) {
		account, err := telegraph.CreateAccount(telegraph.Account{
			ShortName:  "Sandbox",
			AuthorName: "Anonymous",
		})
		assert.NoError(t, err)
		assert.NotNil(t, account)
	})
}