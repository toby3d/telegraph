package telegraph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	t.Run("invalid", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			_, err := CreateAccount(Account{})
			assert.Error(t, err)
		})
		t.Run("without shortname", func(t *testing.T) {
			_, err := CreateAccount(Account{
				ShortName:  "",
				AuthorName: "Anonymous",
			})
			assert.Error(t, err)
		})
	})
	t.Run("valid", func(t *testing.T) {
		account, err := CreateAccount(Account{
			ShortName:  "Sandbox",
			AuthorName: "Anonymous",
		})
		assert.NoError(t, err)
		assert.NotNil(t, account)
	})
}
