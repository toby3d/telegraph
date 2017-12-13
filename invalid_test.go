package telegraph

import "testing"

const (
	invalidAuthorURL = "lolwat"
	invalidPageURL   = "sukablyat'"
	invalidContent   = 42
)

var invalidAccount = &Account{}

func TestInvalidContentFormat(t *testing.T) {
	if _, err := ContentFormat(invalidContent); err != ErrInvalidDataType {
		t.Error()
	}
}

func TestInvalidCreateAccount(t *testing.T) {
	if _, err := CreateAccount(invalidAccount); err == nil {
		t.Error()
	}

	t.Run("invalidCreatePage", testInvalidCreatePage)
	t.Run("invalidEditAccountInfo", testInvalidEditAccountInfo)
	t.Run("invalidEditPage", testInvalidEditPage)
	t.Run("invalidGetAccountInfo", testInvalidGetAccountInfo)
	t.Run("invalidGetPageList", testInvalidGetPageList)
	t.Run("invalidRevokeAccessToken", testInvalidRevokeAccessToken)
}

func testInvalidCreatePage(t *testing.T) {
	if _, err := invalidAccount.CreatePage(&Page{
		AuthorURL: invalidAuthorURL,
	}, false); err == nil {
		t.Error()
	}
}

func testInvalidEditAccountInfo(t *testing.T) {
	if _, err := invalidAccount.EditAccountInfo(&Account{
		AuthorURL: invalidAuthorURL,
	}); err == nil {
		t.Error()
	}
}

func testInvalidEditPage(t *testing.T) {
	if _, err := invalidAccount.EditPage(&Page{
		AuthorURL: invalidAuthorURL,
	}, false); err == nil {
		t.Error()
	}
}

func testInvalidGetAccountInfo(t *testing.T) {
	if _, err := invalidAccount.GetAccountInfo("short_name", "page_count"); err == nil {
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
	if _, err := GetPage(invalidPageURL, true); err == nil {
		t.Error()
	}
}

func TestInvalidGetViewsByPage(t *testing.T) {
	if _, err := GetViews(validPageURL, 2016, 12, 0, -1); err == nil {
		t.Error()
	}
}

func TestInvalidGetViewsByHour(t *testing.T) {
	if _, err := GetViews(validPageURL, 42, 0, 0, 0); err == nil {
		t.Error()
	}
}

func TestInvalidGetViewsByDay(t *testing.T) {
	if _, err := GetViews(validPageURL, 23, 42, 0, 0); err == nil {
		t.Error()
	}
}

func TestInvalidGetViewsByMonth(t *testing.T) {
	if _, err := GetViews(validPageURL, 23, 24, 22, 0); err == nil {
		t.Error()
	}
}

func TestInvalidGetViewsByYear(t *testing.T) {
	if _, err := GetViews(validPageURL, 23, 24, 12, 1980); err == nil {
		t.Error()
	}
}

func testInvalidRevokeAccessToken(t *testing.T) {
	if _, err := invalidAccount.RevokeAccessToken(); err == nil {
		t.Error()
	}
}
