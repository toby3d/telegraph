package telegraph

import (
	"testing"
)

var (
	demoAccount Account
	demoPage    Page
	demoContent = `<p>Hello, World!</p>`
)

func TestCreateAccount(t *testing.T) {
	acc, err := CreateAccount("Sandbox", "Anonymous", "")
	if err != nil {
		t.Error(err.Error())
	}

	demoAccount = *acc
	t.Logf("New account created!\n%#v", *acc)
}

func TestCreatePage(t *testing.T) {
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

func TestEditAccountInfo(t *testing.T) {
	update := &Account{
		ShortName:  "Sandbox",
		AuthorName: "Anonymous",
	}

	info, err := demoAccount.EditAccountInfo(update)
	if err != nil {
		t.Error(err.Error())
	}

	t.Logf("Account updated!\n%#v", info)
}

func TestEditPage(t *testing.T) {
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

func TestGetAccountInfo(t *testing.T) {
	account, err := demoAccount.GetAccountInfo([]string{"short_name", "page_count"})
	if err != nil {
		t.Error(err.Error())
	}

	t.Logf("Account info:\nShort Name: %s\nPage Count: %d", account.ShortName, account.PageCount)
}

func TestGetPageList(t *testing.T) {
	pages, err := demoAccount.GetPageList(0, 3)
	if err != nil {
		t.Error(err.Error())
	}

	t.Logf("Total %d pages\n%#v", pages.TotalCount, pages.Pages)
}

func TestGetPage(t *testing.T) {
	page, err := GetPage(demoPage.Path, true)
	if err != nil {
		t.Error(err.Error())
	}

	t.Logf("%#v", page)
}

func TestGetViews(t *testing.T) {
	views, err := GetViews("Sample-Page-12-15", 2016, 12, 0, -1)
	if err != nil {
		t.Error(err.Error())
	}

	t.Logf("This page have %d views", views.Views)
}

func TestRevokeAccessToken(t *testing.T) {
	t.Logf("Old Access Token: %s", demoAccount.AccessToken)

	token, err := demoAccount.RevokeAccessToken()
	if err != nil {
		t.Error(token)
	}

	t.Logf("New Access Token: %s", token.AccessToken)
}
