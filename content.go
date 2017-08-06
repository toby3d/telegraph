package telegraph

import (
	"bytes"
	"errors"
	"strings"

	html "golang.org/x/net/html"
)

var (
	availableTags = map[string]bool{
		"a":          true,
		"aside":      true,
		"b":          true,
		"blockquote": true,
		"br":         true,
		"code":       true,
		"em":         true,
		"figcaption": true,
		"figure":     true,
		"h3":         true,
		"h4":         true,
		"hr":         true,
		"i":          true,
		"iframe":     true,
		"img":        true,
		"li":         true,
		"ol":         true,
		"p":          true,
		"pre":        true,
		"s":          true,
		"strong":     true,
		"u":          true,
		"ul":         true,
		"video":      true,
	}

	availableAttributes = map[string]bool{
		"href": true,
		"src":  true,
	}
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
	if _, ok := availableTags[strings.ToLower(domNode.Data)]; ok {
		nodeElement.Tag = domNode.Data
		for _, attr := range domNode.Attr {
			if _, ok := availableAttributes[strings.ToLower(attr.Key)]; ok {
				nodeElement.Attrs = map[string]string{
					attr.Key: attr.Val,
				}
			}
		}
	}

	for child := domNode.FirstChild; child != nil; child = child.NextSibling {
		nodeElement.Children = append(nodeElement.Children, domToNode(child))
	}

	return nodeElement
}
