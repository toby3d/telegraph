package telegraph

import "testing"

var (
	demoAccount Account
	demoPage    Page
	demoContent = `<p>Hello, World!</p>`
)

func TestContentFormatByString(t *testing.T) {
	_, err := ContentFormat(demoContent)
	if err != nil {
		t.Error(err)
	}
}

func TestContentFormatByBytes(t *testing.T) {
	_, err := ContentFormat([]byte(demoContent))
	if err != nil {
		t.Error(err)
	}
}

func TestCreateValidAccount(t *testing.T) {
	account, err := CreateAccount("Sandbox", "Anonymous", "")
	if err != nil {
		t.Error(err)
	}

	demoAccount = *account
	t.Logf("New account created!")
}

func TestCreateValidPage(t *testing.T) {
	content, err := ContentFormat(demoContent)
	if err != nil {
		t.Error(err)
	}

	newPage := &Page{
		Title:      "Sample Page",
		AuthorName: "Anonymous",
		AuthorURL:  "https://telegram.me/telegraph",
		Content:    content,
	}

	page, err := demoAccount.CreatePage(newPage, true)
	if err != nil {
		t.Error(err)
	}

	demoPage = *page
	t.Logf("%#v", *page)
}

func TestEditValidAccountInfo(t *testing.T) {
	update := &Account{
		ShortName:  "Sandbox",
		AuthorName: "Anonymous",
		AuthorURL:  "https://telegram.me/telegraph",
	}

	_, err := demoAccount.EditAccountInfo(update)
	if err != nil {
		t.Error(err)
	}

	t.Logf("Account updated!")
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
		t.Error(err)
	}

	t.Logf("%#v", *page)
}

func TestGetValidAccountInfo(t *testing.T) {
	account, err := demoAccount.GetAccountInfo("short_name", "page_count")
	if err != nil {
		t.Error(err)
	}

	t.Logf("%s create %d pages", account.ShortName, account.PageCount)
}

func TestGetValidPageList(t *testing.T) {
	pages, err := demoAccount.GetPageList(0, 3)
	if err != nil {
		t.Error(err)
	}

	t.Logf("Total %d pages\n%#v", pages.TotalCount, pages.Pages)
}

func TestGetValidPage(t *testing.T) {
	page, err := GetPage(demoPage.Path, true)
	if err != nil {
		t.Error(err)
	}

	t.Logf("%#v", *page)
}

func TestGetValidViews(t *testing.T) {
	views, err := GetViews("Sample-Page-12-15", -1, 0, 12, 2016)
	if err != nil {
		t.Error(err)
	}

	t.Logf("This page have %d views", views.Views)
}

func TestRevokeAccessToken(t *testing.T) {
	_, err := demoAccount.RevokeAccessToken()
	if err != nil {
		t.Error(err)
	}

	t.Logf("New Access Token set!")
}
