package telegraph

import (
	"testing"
)

var (
	demoAccount Account
	demoPage    Page
	demoContent = `<p>Hello, World!</p>`
)

func TestContentFormatByString(t *testing.T) {
	_, err := ContentFormat(demoContent)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestContentFormatByBytes(t *testing.T) {
	_, err := ContentFormat([]byte(demoContent))
	if err != nil {
		t.Error(err.Error())
	}
}

func TestContentFormatByWTF(t *testing.T) {
	_, err := ContentFormat(42)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestCreateInvalidAccount(t *testing.T) {
	_, err := CreateAccount("", "", "")
	if err != nil {
		t.Error(err.Error())
	}
}

func TestCreateValidAccount(t *testing.T) {
	acc, err := CreateAccount("Sandbox", "Anonymous", "")
	if err != nil {
		t.Error(err.Error())
	}

	demoAccount = *acc
	t.Logf("New account created!\n%#v", *acc)
}

func TestCreateInvalidPage(t *testing.T) {
	newPage := &Page{
		AuthorURL: "lolwat",
	}
	_, err := demoAccount.CreatePage(newPage, false)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestCreateValidPage(t *testing.T) {
	content, err := ContentFormat(demoContent)
	if err != nil {
		t.Error(err.Error())
	}

	newPage := &Page{
		Title:      "Sample Page",
		AuthorName: "Anonymous",
		Content:    content,
	}

	page, err := demoAccount.CreatePage(newPage, true)
	if err != nil {
		t.Error(err.Error())
	}

	demoPage = *page
	t.Logf("%#v", *page)
}

func TestEditInvalidAccountInfo(t *testing.T) {
	var update Account

	_, err := demoAccount.EditAccountInfo(&update)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestEditValidAccountInfo(t *testing.T) {
	update := &Account{
		ShortName:  "Sandbox",
		AuthorName: "Anonymous",
		AuthorURL:  "https://telegram.me/telegraph",
	}

	info, err := demoAccount.EditAccountInfo(update)
	if err != nil {
		t.Error(err.Error())
	}

	t.Logf("Account updated!\n%#v", info)
}

func TestEditInvalidPage(t *testing.T) {
	update := &Page{
		AuthorURL: "lolwat",
	}

	_, err := demoAccount.EditPage(update, false)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestEditValidPage(t *testing.T) {
	content, err := ContentFormat(demoContent)
	if err != nil {
		t.Error(err)
	}

	update := &Page{
		Path:       demoPage.Path,
		Title:      "Sample Page",
		AuthorName: "Anonymous",
		Content:    content,
	}

	page, err := demoAccount.EditPage(update, true)
	if err != nil {
		t.Error(err.Error())
	}

	t.Logf("%#v", *page)
}

func TestGetInvalidAccountInfo(t *testing.T) {
	var account Account
	_, err := account.GetAccountInfo([]string{"short_name", "page_count"})
	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetValidAccountInfo(t *testing.T) {
	account, err := demoAccount.GetAccountInfo([]string{"short_name", "page_count"})
	if err != nil {
		t.Error(err.Error())
	}

	t.Logf("Account info:\nShort Name: %s\nPage Count: %d", account.ShortName, account.PageCount)
}

func TestGetInvalidPageList(t *testing.T) {
	var account Account
	_, err := account.GetPageList(0, 3)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetValidPageList(t *testing.T) {
	pages, err := demoAccount.GetPageList(0, 3)
	if err != nil {
		t.Error(err.Error())
	}

	t.Logf("Total %d pages\n%#v", pages.TotalCount, pages.Pages)
}

func TestGetInvalidPage(t *testing.T) {
	_, err := GetPage("lolwat", true)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetValidPage(t *testing.T) {
	page, err := GetPage(demoPage.Path, true)
	if err != nil {
		t.Error(err.Error())
	}

	t.Logf("%#v", page)
}

func TestGetInvalidViewsByPage(t *testing.T) {
	_, err := GetViews("lolwat", 2016, 12, 0, -1)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetInvalidViewsByHour(t *testing.T) {
	_, err := GetViews("Sample-Page-12-15", 0, 0, 0, 42)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetInvalidViewsByDay(t *testing.T) {
	_, err := GetViews("Sample-Page-12-15", 0, 0, 42, -1)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetInvalidViewsByMonth(t *testing.T) {
	_, err := GetViews("Sample-Page-12-15", 0, 22, 24, -1)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetInvalidViewsByYear(t *testing.T) {
	_, err := GetViews("Sample-Page-12-15", 1984, 12, 24, 23)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetValidViews(t *testing.T) {
	views, err := GetViews("Sample-Page-12-15", 2016, 12, 0, -1)
	if err != nil {
		t.Error(err.Error())
	}

	t.Logf("This page have %d views", views.Views)
}

func TestRevokeInvalidAccessToken(t *testing.T) {
	var account Account
	_, err := account.RevokeAccessToken()
	if err != nil {
		t.Error(err.Error())
	}
}

func TestRevokeAccessToken(t *testing.T) {
	t.Logf("Old Access Token: %s", demoAccount.AccessToken)

	token, err := demoAccount.RevokeAccessToken()
	if err != nil {
		t.Error(err.Error())
	}

	t.Logf("New Access Token: %s", token.AccessToken)
}
