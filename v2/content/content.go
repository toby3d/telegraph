// The content package contains utilities for converting page content from HTML
// DOM format to Telegraph API supported JSON objects and vice versa.
package content

import (
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"

	"source.toby3d.me/toby3d/telegraph"
)

func DomToNode(domNode *html.Node) *telegraph.Node {
	switch domNode.Type {
	default:
		return nil
	case html.TextNode:
		return &telegraph.Node{Text: domNode.Data}
	case html.ElementNode:
	}

	tag, _ := telegraph.NewTag(domNode.DataAtom)
	nodeElement := telegraph.NewNodeElement(tag)

	for _, attr := range domNode.Attr {
		switch attr.Key {
		case "src":
			nodeElement.Attrs.Src = attr.Val
		case "href":
			nodeElement.Attrs.Href = attr.Val
		}
	}

	for child := domNode.FirstChild; child != nil; child = child.NextSibling {
		if node := DomToNode(child); node != nil {
			nodeElement.Children = append(nodeElement.Children, *node)
		}
	}

	return &telegraph.Node{Element: nodeElement}
}

func NodeToDom(node telegraph.Node) *html.Node {
	if node.Text != "" {
		return &html.Node{
			Type:     html.TextNode,
			DataAtom: atom.Plaintext,
			Data:     node.Text,
			Attr:     make([]html.Attribute, 0),
		}
	}

	domNode := &html.Node{
		Type:     html.ElementNode,
		DataAtom: 0,
		Data:     "",
		Attr:     make([]html.Attribute, 0),
	}

	if node.Element != nil {
		domNode = &html.Node{
			Type:     html.ElementNode,
			DataAtom: node.Element.Tag.Atom(),
			Data:     node.Text,
			Attr:     make([]html.Attribute, 0, 2),
		}

		if node.Element.Attrs != nil {
			for key, val := range map[string]string{
				"href": node.Element.Attrs.Href,
				"src":  node.Element.Attrs.Src,
			} {
				if val == "" {
					continue
				}

				domNode.Attr = append(domNode.Attr, html.Attribute{
					Key: key,
					Val: val,
				})
			}
		}
	}

	for _, child := range node.Element.Children {
		domNode.AppendChild(NodeToDom(child))
	}

	return domNode
}