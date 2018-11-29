package telegraph_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gitlab.com/toby3d/telegraph"
)

const (
	validTitle         = "Testing"
	validShortName     = "Sandbox"
	validAuthorName    = "Anonymous"
	validNewAuthorName = "Gopher"
	validAuthorURL     = "https://t.me/telegraph"
	validPageURL       = "Sample-Page-12-15"
)

var (
	validContentDOM []telegraph.Node
	validAccount    *telegraph.Account
	validPage       *telegraph.Page

	validContent = `<p>Hello, World!</p>`
)

func testValidContentFormatByString(t *testing.T) {
	var err error
	validContentDOM, err = telegraph.ContentFormat(validContent)
	assert.NoError(t, err)
	assert.NotEmpty(t, validContentDOM)
}

func testValidContentFormatByBytes(t *testing.T) {
	var err error
	validContentDOM, err = telegraph.ContentFormat([]byte(validContent))
	assert.NoError(t, err)
	assert.NotEmpty(t, validContentDOM)
}

func TestValidCreateAccount(t *testing.T) {
	var err error
	validAccount, err = telegraph.CreateAccount(&telegraph.Account{
		ShortName: validShortName,
		// AuthorName: validAuthorName,
	})
	assert.NoError(t, err)

	t.Run("validCreatePage", testValidCreatePage)
	t.Run("validEditAccountInfo", testValidEditAccountInfo)
	t.Run("validGetAccountInfo", testValidGetAccountInfo)
	t.Run("validGetPageList", testValidGetPageList)
	t.Run("validRevokeAccessToken", testValidRevokeAccessToken)
}

func testValidCreatePage(t *testing.T) {
	t.Run("validContentFormatByString", testValidContentFormatByString)
	t.Run("validContentFormatByBytes", testValidContentFormatByBytes)

	var err error
	validPage, err = validAccount.CreatePage(&telegraph.Page{
		Title:      validTitle,
		AuthorName: validAuthorName,
		AuthorURL:  validAuthorURL,
		Content:    validContentDOM,
	}, true)
	assert.NoError(t, err)
	assert.NotEmpty(t, validPage.URL)
	t.Run("validEditPage", testValidEditPage)
}

func testValidEditAccountInfo(t *testing.T) {
	update, err := validAccount.EditAccountInfo(&telegraph.Account{
		ShortName:  validShortName,
		AuthorName: validNewAuthorName,
		AuthorURL:  validAuthorURL,
	})
	assert.NoError(t, err)
	assert.NotEqual(t, validAccount.AuthorName, update.AuthorName)
}

func testValidEditPage(t *testing.T) {
	var err error
	validPage, err = validAccount.EditPage(&telegraph.Page{
		Path:       validPage.Path,
		Title:      validTitle,
		AuthorName: validAuthorName,
		Content:    validContentDOM,
	}, true)
	assert.NoError(t, err)
}

func testValidGetAccountInfo(t *testing.T) {
	info, err := validAccount.GetAccountInfo(telegraph.FieldShortName, telegraph.FieldPageCount)
	assert.NoError(t, err)
	assert.Equal(t, validAccount.ShortName, info.ShortName)
}

func TestValidGetPage(t *testing.T) {
	page, err := telegraph.GetPage(validPage.Path, true)
	assert.NoError(t, err)
	assert.Equal(t, validPage.Title, page.Title)
}

func testValidGetPageList(t *testing.T) {
	pages, err := validAccount.GetPageList(0, 3)
	assert.NoError(t, err)
	assert.NotZero(t, pages.TotalCount)
}

func TestValidGetViews(t *testing.T) {
	stats, err := telegraph.GetViews(validPageURL, time.Date(2016, time.December, 0, 0, 0, 0, 0, time.UTC))
	assert.NoError(t, err)
	t.Log("get", stats.Views, "views")
}

func testValidRevokeAccessToken(t *testing.T) {
	oldToken := validAccount.AccessToken
	var err error
	validAccount, err = validAccount.RevokeAccessToken()
	assert.NoError(t, err)
	assert.NotEmpty(t, validAccount.AccessToken)
	assert.NotEqual(t, validAccount.AccessToken, oldToken)
}
