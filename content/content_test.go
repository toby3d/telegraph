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