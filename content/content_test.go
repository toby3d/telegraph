package content_test

import (
	"log"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"

	"source.toby3d.me/toby3d/telegraph"
	"source.toby3d.me/toby3d/telegraph/content"
)

func TestDomToNode(t *testing.T) {
	t.Parallel()

	nodes, err := html.ParseFragment(strings.NewReader(`<p>Hello, world!</p>`), &html.Node{Type: html.ElementNode})
	if err != nil {
		log.Fatalln("cannot parse HTML content:", err)
	}

	expect := &telegraph.Node{Element: &telegraph.NodeElement{
		Tag:      telegraph.P,
		Attrs:    new(telegraph.Attributes),
		Children: []telegraph.Node{{Text: "Hello, world!"}},
	}}

	actual := content.DomToNode(nodes[0])
	if diff := cmp.Diff(expect, actual, cmpopts.EquateComparable(telegraph.Tag{})); diff != "" {
		t.Error(diff)
	}
}

func TestNodeToDom(t *testing.T) {
	t.Parallel()

	expect := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.P,
		Attr:     make([]html.Attribute, 0),
	}
	expect.AppendChild(&html.Node{
		Type:     html.TextNode,
		DataAtom: atom.Plaintext,
		Data:     "Hello, world!",
		Attr:     make([]html.Attribute, 0),
	})

	actual := content.NodeToDom(telegraph.Node{Element: &telegraph.NodeElement{
		Tag:      telegraph.P,
		Attrs:    new(telegraph.Attributes),
		Children: []telegraph.Node{{Text: "Hello, world!"}},
	}})

	if diff := cmp.Diff(expect, actual); diff != "" {
		t.Error(diff)
	}
}

/*
func Example() {
	account, err := (telegraph.CreateAccount{
		ShortName: *util.Must(telegraph.NewShortName("Sandbox")),
	}).Do(context.Background(), http.DefaultClient)
	if err != nil {
		log.Fatalln("cannot create account:", err)
	}

	article, err := html.Parse(strings.NewReader(`<p>Hello, world!</p>`))
	if err != nil {
		log.Fatalln("cannot parse HTML content:", err)
	}

	pageContent := make([]telegraph.Node, 0)
	if node := content.DomToNode(article); node != nil {
		pageContent = append(pageContent, node.Element.Children...)
	}

	pageTitle, err := telegraph.NewTitle("Title of page")
	if err != nil {
		log.Fatalln("cannot parse title:", err)
	}

	page, err := telegraph.CreatePage{
		AccessToken:   account.AccessToken,
		Title:         *pageTitle,
		Content:       pageContent,
		ReturnContent: true,
	}.Do(context.Background(), http.DefaultClient)
	if err != nil {
		log.Fatalln("cannot create page:", err)
	}

	if 0 < len(page.Content) {
		for child := article.FirstChild; child != nil; child = article.NextSibling {
			article.RemoveChild(child)
		}

		article.AppendChild(content.NodeToDom(telegraph.Node{Element: &telegraph.NodeElement{
			Children: page.Content,
		}}))
	}

	var buf bytes.Buffer
	if err = html.Render(&buf, article); err != nil {
		log.Fatalln("cannot render HTML content:", err)
	}

	fmt.Print(buf.String())
	// Output: <p>Hello, world!</p>
}*/