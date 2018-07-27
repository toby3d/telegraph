package telegraph_test

import (
	"testing"
	"time"

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

	validAccount = new(telegraph.Account)
	validPage    = new(telegraph.Page)
	validContent = `<p>Hello, World!</p>`
)

func testValidContentFormatByString(t *testing.T) {
	var err error
	validContentDOM, err = telegraph.ContentFormat(validContent)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	if len(validContentDOM) <= 0 {
		t.Error("DOM content is nil")
	}
}

func testValidContentFormatByBytes(t *testing.T) {
	var err error
	validContentDOM, err = telegraph.ContentFormat([]byte(validContent))
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	if len(validContentDOM) <= 0 {
		t.Error("DOM content is nil")
	}
}

func TestValidCreateAccount(t *testing.T) {
	var err error
	validAccount, err = telegraph.CreateAccount(&telegraph.Account{
		ShortName: validShortName,
		// AuthorName: validAuthorName,
	})
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

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
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	if validPage.URL == "" {
		t.Error("new page not contain URL")
	}

	t.Run("validEditPage", testValidEditPage)
}

func testValidEditAccountInfo(t *testing.T) {
	update, err := validAccount.EditAccountInfo(&telegraph.Account{
		ShortName:  validShortName,
		AuthorName: validNewAuthorName,
		AuthorURL:  validAuthorURL,
	})
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	if update.AuthorName == validAccount.AuthorName {
		t.Error("account not updated")
	}
}

func testValidEditPage(t *testing.T) {
	var err error
	validPage, err = validAccount.EditPage(&telegraph.Page{
		Path:       validPage.Path,
		Title:      validTitle,
		AuthorName: validAuthorName,
		Content:    validContentDOM,
	}, true)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
}

func testValidGetAccountInfo(t *testing.T) {
	info, err := validAccount.GetAccountInfo(
		telegraph.FieldShortName,
		telegraph.FieldPageCount,
	)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	if info.ShortName != validAccount.ShortName {
		t.Error("get wrong account info")
	}
}

func TestValidGetPage(t *testing.T) {
	page, err := telegraph.GetPage(validPage.Path, true)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	if page.Title != validPage.Title {
		t.Error("get wrong page")
	}
}

func testValidGetPageList(t *testing.T) {
	pages, err := validAccount.GetPageList(0, 3)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	if pages.TotalCount <= 0 {
		t.Error("no one page in page list")
	}
}

func TestValidGetViews(t *testing.T) {
	stats, err := telegraph.GetViews(
		validPageURL,
		time.Date(2016, time.December, 0, 0, 0, 0, 0, time.UTC),
	)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	t.Log("get", stats.Views, "views")
}

func testValidRevokeAccessToken(t *testing.T) {
	oldToken := validAccount.AccessToken

	var err error
	validAccount, err = validAccount.RevokeAccessToken()
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	if validAccount.AccessToken == "" {
		t.Error("revokeAccessToken return nothing")
	}

	if oldToken == validAccount.AccessToken {
		t.Error("old and new tokens are equal")
	}
}
