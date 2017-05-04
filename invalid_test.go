package telegraph

import "testing"

func TestContentFormatByWTF(t *testing.T) {
	_, err := ContentFormat(42)
	if err == nil {
		t.Error()
	}
	t.Log(err.Error())
}

func TestCreateInvalidAccount(t *testing.T) {
	_, err := CreateAccount("", "", "")
	if err == nil {
		t.Error()
	}
	t.Log(err.Error())
}

func TestCreateInvalidPage(t *testing.T) {
	newPage := &Page{
		AuthorURL: "lolwat",
	}
	_, err := demoAccount.CreatePage(newPage, false)
	if err == nil {
		t.Error()
	}
	t.Log(err.Error())
}

func TestEditInvalidAccountInfo(t *testing.T) {
	var update Account

	_, err := demoAccount.EditAccountInfo(&update)
	if err == nil {
		t.Error()
	}
	t.Log(err.Error())
}

func TestEditInvalidPage(t *testing.T) {
	update := &Page{
		AuthorURL: "lolwat",
	}

	_, err := demoAccount.EditPage(update, false)
	if err == nil {
		t.Error()
	}
	t.Log(err.Error())
}

func TestGetInvalidAccountInfo(t *testing.T) {
	var account Account
	_, err := account.GetAccountInfo("short_name", "page_count")
	if err == nil {
		t.Error()
	}
	t.Log(err.Error())
}

func TestGetInvalidPageList(t *testing.T) {
	var account Account
	_, err := account.GetPageList(0, 3)
	if err == nil {
		t.Error()
	}
	t.Log(err.Error())
}

func TestGetInvalidPageListByOffset(t *testing.T) {
	var account Account
	_, err := account.GetPageList(-42, 3)
	if err == nil {
		t.Error()
	}
	t.Log(err.Error())
}

func TestGetInvalidPageListByLimit(t *testing.T) {
	var account Account
	_, err := account.GetPageList(0, 9000)
	if err == nil {
		t.Error()
	}
	t.Log(err.Error())
}

func TestGetInvalidPage(t *testing.T) {
	_, err := GetPage("lolwat", true)
	if err == nil {
		t.Error()
	}
	t.Log(err.Error())
}

func TestGetInvalidViewsByPage(t *testing.T) {
	_, err := GetViews("lolwat", 2016, 12, 0, -1)
	if err == nil {
		t.Error()
	}
	t.Log(err.Error())
}

func TestGetInvalidViewsByHour(t *testing.T) {
	_, err := GetViews("Sample-Page-12-15", 42, 0, 0, 0)
	if err == nil {
		t.Error()
	}
	t.Log(err.Error())
}

func TestGetInvalidViewsByDay(t *testing.T) {
	_, err := GetViews("Sample-Page-12-15", 23, 42, 0, 0)
	if err == nil {
		t.Error()
	}
	t.Log(err.Error())
}

func TestGetInvalidViewsByMonth(t *testing.T) {
	_, err := GetViews("Sample-Page-12-15", 23, 24, 22, 0)
	if err == nil {
		t.Error()
	}
	t.Log(err.Error())
}

func TestGetInvalidViewsByYear(t *testing.T) {
	_, err := GetViews("Sample-Page-12-15", 23, 24, 12, 1980)
	if err == nil {
		t.Error()
	}
	t.Log(err.Error())
}

func TestRevokeInvalidAccessToken(t *testing.T) {
	var account Account
	_, err := account.RevokeAccessToken()
	if err == nil {
		t.Error()
	}
	t.Log(err.Error())
}
