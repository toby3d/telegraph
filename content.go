package telegraph

import (
	"bytes"
	"errors"
	"strings"

	"golang.org/x/net/html"
)

// ContentFormat transforms data to a DOM-based format to represent the
// content of the page.
func ContentFormat(data interface{}) ([]Node, error) {
	var doc html.Node
	switch dst := data.(type) {
	case string:
		dom, err := html.Parse(strings.NewReader(dst))
		if err != nil {
			return nil, err
		}
		doc = *dom
	case []byte:
		dom, err := html.Parse(bytes.NewReader(dst))
		if err != nil {
			return nil, err
		}
		doc = *dom
	default:
		return nil, errors.New("invalid data type, use []byte or string")
	}

	var content []Node
	content = append(content, domToNode(doc.FirstChild))

	return content, nil
}

func domToNode(domNode *html.Node) interface{} {
	if domNode.Type == html.TextNode {
		return domNode.Data
	}

	if domNode.Type != html.ElementNode {
		return nil
	}

	var nodeElement NodeElement
	switch strings.ToLower(domNode.Data) {
	case "a", "aside", "b", "blockquote", "br", "code", "em", "figcaption", "figure", "h3", "h4", "hr", "i", "iframe", "img", "li", "ol", "p", "pre", "s", "strong", "u", "ul", "video":
		nodeElement.Tag = domNode.Data

		for i := range domNode.Attr {
			switch strings.ToLower(domNode.Attr[i].Key) {
			case "href", "src":
				nodeElement.Attrs = map[string]string{
					domNode.Attr[i].Key: domNode.Attr[i].Val,
				}
			default:
				continue
			}
		}
	}

	for child := domNode.FirstChild; child != nil; child = child.NextSibling {
		nodeElement.Children = append(nodeElement.Children, domToNode(child))
	}

	return nodeElement
}
