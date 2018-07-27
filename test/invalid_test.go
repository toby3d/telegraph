package telegraph_test

import (
	"testing"
	"time"

	"gitlab.com/toby3d/telegraph"
)

const (
	invalidAuthorURL = "lolwat"
	invalidPageURL   = "sukablyat'"
	invalidContent   = 42
)

var invalidAccount = new(telegraph.Account)

func TestInvalidContentFormat(t *testing.T) {
	if _, err := telegraph.ContentFormat(invalidContent); err != telegraph.ErrInvalidDataType {
		t.Error()
	}
}

func TestInvalidCreateAccount(t *testing.T) {
	if _, err := telegraph.CreateAccount(invalidAccount); err == nil {
		t.Error()
	}

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
	if _, err := invalidAccount.CreatePage(&telegraph.Page{
		AuthorURL: invalidAuthorURL,
	}, false); err == nil {
		t.Error()
	}
}

func testInvalidEditAccountInfo(t *testing.T) {
	if _, err := invalidAccount.EditAccountInfo(&telegraph.Account{
		AuthorURL: invalidAuthorURL,
	}); err == nil {
		t.Error()
	}
}

func testInvalidEditPage(t *testing.T) {
	if _, err := invalidAccount.EditPage(&telegraph.Page{
		AuthorURL: invalidAuthorURL,
	}, false); err == nil {
		t.Error()
	}
}

func testInvalidGetAccountInfo(t *testing.T) {
	if _, err := invalidAccount.GetAccountInfo(
		telegraph.FieldShortName,
		telegraph.FieldPageCount,
	); err == nil {
		t.Error()
	}
}

func testInvalidGetPageList(t *testing.T) {
	if _, err := invalidAccount.GetPageList(0, 3); err == nil {
		t.Error()
	}
}

func testInvalidGetPageListByOffset(t *testing.T) {
	if _, err := invalidAccount.GetPageList(-42, 3); err == nil {
		t.Error()
	}
}

func testInvalidGetPageListByLimit(t *testing.T) {
	if _, err := invalidAccount.GetPageList(0, 9000); err == nil {
		t.Error()
	}
}

func TestInvalidGetPage(t *testing.T) {
	if _, err := telegraph.GetPage(invalidPageURL, true); err == nil {
		t.Error()
	}
}

func TestInvalidGetViewsByPage(t *testing.T) {
	if _, err := telegraph.GetViews(
		invalidPageURL,
		time.Date(2016, time.December, 0, 0, 0, 0, 0, time.UTC),
	); err == nil {
		t.Error()
	}
}

func TestInvalidGetViewsByHour(t *testing.T) {
	if _, err := telegraph.GetViews(
		validPageURL,
		time.Date(0, 0, 0, 42, 0, 0, 0, time.UTC),
	); err == nil {
		t.Error()
	}
}

func TestInvalidGetViewsByDay(t *testing.T) {
	if _, err := telegraph.GetViews(
		validPageURL,
		time.Date(0, 0, 42, 23, 0, 0, 0, time.UTC),
	); err == nil {
		t.Error()
	}
}

func TestInvalidGetViewsByMonth(t *testing.T) {
	if _, err := telegraph.GetViews(
		validPageURL,
		time.Date(0, 22, 24, 23, 0, 0, 0, time.UTC),
	); err == nil {
		t.Error()
	}
}

func TestInvalidGetViewsByYear(t *testing.T) {
	if _, err := telegraph.GetViews(
		validPageURL,
		time.Date(1980, time.December, 24, 23, 0, 0, 0, time.UTC),
	); err == nil {
		t.Error()
	}
}

func testInvalidRevokeAccessToken(t *testing.T) {
	if _, err := invalidAccount.RevokeAccessToken(); err == nil {
		t.Error()
	}
}
