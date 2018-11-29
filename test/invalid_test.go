package telegraph_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gitlab.com/toby3d/telegraph"
)

const (
	invalidAuthorURL = "lolwat"
	invalidPageURL   = "sukablyat'"
	invalidContent   = 42
)

var invalidAccount = new(telegraph.Account)

func TestInvalidContentFormat(t *testing.T) {
	_, err := telegraph.ContentFormat(invalidContent)
	assert.EqualError(t, telegraph.ErrInvalidDataType, err.Error())
}

func TestInvalidCreateAccount(t *testing.T) {
	_, err := telegraph.CreateAccount(invalidAccount)
	assert.Error(t, err)

	t.Run("invalidCreatePage", testInvalidCreatePage)
	t.Run("invalidEditAccountInfo", testInvalidEditAccountInfo)
	t.Run("invalidEditPage", testInvalidEditPage)
	t.Run("invalidGetAccountInfo", testInvalidGetAccountInfo)
	t.Run("invalidGetPageList", testInvalidGetPageList)
	t.Run("invalidGetPageListByLimit", testInvalidGetPageListByLimit)
	t.Run("invalidGetPageListByOffset", testInvalidGetPageListByOffset)
	t.Run("invalidRevokeAccessToken", testInvalidRevokeAccessToken)
}

func testInvalidCreatePage(t *testing.T) {
	_, err := invalidAccount.CreatePage(&telegraph.Page{AuthorURL: invalidAuthorURL}, false)
	assert.Error(t, err)
}

func testInvalidEditAccountInfo(t *testing.T) {
	_, err := invalidAccount.EditAccountInfo(&telegraph.Account{AuthorURL: invalidAuthorURL})
	assert.Error(t, err)
}

func testInvalidEditPage(t *testing.T) {
	_, err := invalidAccount.EditPage(&telegraph.Page{AuthorURL: invalidAuthorURL}, false)
	assert.Error(t, err)
}

func testInvalidGetAccountInfo(t *testing.T) {
	_, err := invalidAccount.GetAccountInfo(telegraph.FieldShortName, telegraph.FieldPageCount)
	assert.Error(t, err)
}

func testInvalidGetPageList(t *testing.T) {
	_, err := invalidAccount.GetPageList(0, 3)
	assert.Error(t, err)
}

func testInvalidGetPageListByOffset(t *testing.T) {
	_, err := invalidAccount.GetPageList(-42, 3)
	assert.Error(t, err)
}

func testInvalidGetPageListByLimit(t *testing.T) {
	_, err := invalidAccount.GetPageList(0, 9000)
	assert.Error(t, err)
}

func TestInvalidGetPage(t *testing.T) {
	_, err := telegraph.GetPage(invalidPageURL, true)
	assert.Error(t, err)
}

func TestInvalidGetViewsByPage(t *testing.T) {
	_, err := telegraph.GetViews(invalidPageURL, time.Date(2016, time.December, 0, 0, 0, 0, 0, time.UTC))
	assert.Error(t, err)
}

func TestInvalidGetViewsByHour(t *testing.T) {
	_, err := telegraph.GetViews(validPageURL, time.Date(0, 0, 0, 42, 0, 0, 0, time.UTC))
	assert.Error(t, err)
}

func TestInvalidGetViewsByDay(t *testing.T) {
	_, err := telegraph.GetViews(validPageURL, time.Date(0, 0, 42, 23, 0, 0, 0, time.UTC))
	assert.Error(t, err)
}

func TestInvalidGetViewsByMonth(t *testing.T) {
	_, err := telegraph.GetViews(validPageURL, time.Date(0, 22, 24, 23, 0, 0, 0, time.UTC))
	assert.Error(t, err)
}

func TestInvalidGetViewsByYear(t *testing.T) {
	_, err := telegraph.GetViews(validPageURL, time.Date(1980, time.December, 24, 23, 0, 0, 0, time.UTC))
	assert.Error(t, err)
}

func testInvalidRevokeAccessToken(t *testing.T) {
	_, err := invalidAccount.RevokeAccessToken()
	assert.Error(t, err)
}
