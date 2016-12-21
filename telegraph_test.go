package telegraph

import "testing"

var (
	demoAccount = &Account{
		AccessToken: "b968da509bb76866c35425099bc0989a5ec3b32997d55286c657e6994bbb",
	}
	demoPage = &Page{
		Path: "Sample-Page-12-15",
	}
	demoContent = `<p>Hello, world!<p>`
)

func TestCreateAccount(t *testing.T) {
	newAccount, err := CreateAccount("Sandbox", "Anonymous", "")
	if err != nil {
		t.Error(err)
	}

	t.Logf("New account created!\nAccess Token: %s\nAuth URL: %s\nShort Name: %s\nAuthor Name: %s\nPage Count: %d", newAccount.AccessToken, newAccount.AuthURL, newAccount.ShortName, newAccount.AuthorName, newAccount.PageCount)
}

/*
func TestCreatePage(t *testing.T) {
	newPage := &Page{
		Title:      "Sample Page",
		AuthorName: "Anonymous",
		Content:    demoContent,
	}

	demoPage, err = demoAccount.CreatePage(newPage, true)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v", demoPage)
}
*/

func TestEditAccountInfo(t *testing.T) {
	update := &Account{
		ShortName:  "Sandbox",
		AuthorName: "Anonymous",
	}

	info, err := demoAccount.EditAccountInfo(update)
	if err != nil {
		t.Error(err)
	}

	t.Logf("Account updated!\nNew Short Name: %s\nNew Author Name: %s", info.ShortName, info.AuthorName)
}

/*
func TestEditPage(t *testing.T) {
	update := &Page{
		Path:       demoPage.Path,
		Title:      "",
		AuthorName: "Anonymous",
		Content:    demoContent,
	}

	page, err := demoAccount.EditPage(update, true)
	if err != nil {
		t.Error(err)
	}

	t.Logf("%#v", page)
}
*/

func TestGetAccountInfo(t *testing.T) {
	account, err := demoAccount.GetAccountInfo([]string{"short_name", "page_count"})
	if err != nil {
		t.Error(err)
	}

	t.Logf("Account info:\nShort Name: %s\nPage Count: %d", account.ShortName, account.PageCount)
}

/*
func TestGetPage(t *testing.T) {
	page, err := GetPage(demoPage.Path, true)
	if err != nil {
		t.Error(err)
	}

	t.Logf("%#v", page)
}
*/

func TestGetPageList(t *testing.T) {
	list, err := demoAccount.GetPageList(0, 3)
	if err != nil {
		t.Error(err)
	}

	t.Logf("Total %d pages\nPages Raw: %#v", list.TotalCount, list.Pages)
}

func TestGetViews(t *testing.T) {
	views, err := GetViews(demoPage.Path, 2016, 12, 0, -1)
	if err != nil {
		t.Error(err)
	}

	t.Logf("This page have %d views", views.Views)
}

func TestRevokeAccessToken(t *testing.T) {
	account, err := CreateAccount("Sandbox", "Anonymous", "")
	if err != nil {
		t.Error(err)
	}

	t.Logf("Old Access Token: %s", account.AccessToken)

	token, err := account.RevokeAccessToken()
	if err != nil {
		t.Error(token)
	}

	t.Logf("New Access Token: %s", token.AccessToken)
}
