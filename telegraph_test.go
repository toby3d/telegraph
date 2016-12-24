package telegraph

import (
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"testing"
)

var (
	demoAccount Account
	demoPage    Page
	demoContent = `<p>Hello, world!<p>`
	demoURL     = "https://galyonk.in/whats-with-weak-aaa-sales-dcd7744ef205"
)

func parse(url string) ([]Node, error) {
	dom, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}

	article := dom.Find("article").Children()

	var content []Node
	for i := range article.Nodes {
		content = append(content, domToNode(article.Nodes[i]))
	}

	return content, nil
}

func domToNode(domNode *html.Node) interface{} {
	if domNode.Type == html.TextNode && domNode.Data != "" {
		return domNode.Data
	}

	if domNode.Type != html.ElementNode {
		return nil
	}

	var nodeElement NodeElement

	allowTags := map[string]bool{"a": true, "aside": true, "b": true, "blockquote": true, "br": true, "code": true, "em": true, "figcaption": true, "figure": true, "h3": true, "h4": true, "hr": true, "i": true, "iframe": true, "img": true, "li": true, "ol": true, "p": true, "pre": true, "s": true, "strong": true, "u": true, "ul": true, "video": true}
	if _, ok := allowTags[domNode.Data]; ok {
		nodeElement.Tag = domNode.Data

		for i := range domNode.Attr {
			attr := domNode.Attr[i]
			if attr.Key == "href" || attr.Key == "src" {
				if nodeElement.Attrs == nil {
					break
				}
				nodeElement.Attrs[0].Val = attr.Val
			}
		}
	}

	for child := domNode.FirstChild; child != nil; child = child.NextSibling {
		nodeElement.Children = append(nodeElement.Children, domToNode(child))
	}

	return nodeElement
}

func TestCreateAccount(t *testing.T) {
	acc, err := CreateAccount("Sandbox", "Anonymous", "")
	if err != nil {
		t.Error(err.Error())
	}
	demoAccount = *acc
	t.Logf("New account created!\n%#v", acc)
}

func TestCreatePage(t *testing.T) {
	content, err := parse("https://blog.toby3d.ru/five-sentences/")
	if err != nil {
		t.Error(err.Error())
	}

	newPage := &Page{
		Title:      "5 sentences",
		AuthorName: "toby3d",
		Content:    content,
	}

	page, err := demoAccount.CreatePage(newPage, true)
	if err != nil {
		t.Error(err.Error())
	}

	demoPage = *page
	t.Logf("%#v", page)
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
	content, err := parse(demoURL)
	if err != nil {
		t.Error(err)
	}

	update := &Page{
		Path:       demoPage.Path,
		Title:      "AAA Games",
		AuthorName: "Galyonkin",
		Content:    content,
	}

	page, err := demoAccount.EditPage(update, true)
	if err != nil {
		t.Error(err.Error())
	}

	t.Logf("%#v", page)
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
	page, err := GetPage("Sample-Page-12-15", true)
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
