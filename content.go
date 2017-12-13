package telegraph

import (
	"bytes"
	"errors"
	"io"
	"strings"

	"golang.org/x/net/html"
)

type (
	// Node is abstract object represents a DOM Node. It can be a String which represents a DOM text
	// node or a NodeElement object.
	Node interface{}

	// NodeElement represents a DOM element node.
	NodeElement struct {
		// Name of the DOM element.
		// Available tags: a, aside, b, blockquote, br, code, em, figcaption, figure, h3, h4, hr, i,
		// iframe, img, li, ol, p, pre, s, strong, u, ul, video.
		Tag string `json:"tag"`

		// Attributes of the DOM element. Key of object represents name of attribute, value
		// represents value of attribute.
		// Available attributes: href, src.
		Attrs map[string]string `json:"attrs,omitempty"` // optional

		// List of child nodes for the DOM element.
		Children []Node `json:"children,omitempty"` // optional
	}
)

// ErrInvalidDataType is returned when ContentFormat function are passed a data argument of invalid
// type.
var ErrInvalidDataType = errors.New("invalid data type")

// ContentFormat transforms data to a DOM-based format to represent the content of the page.
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
	case io.Reader:
		dom, err := html.Parse(dst)
		if err != nil {
			return nil, err
		}
		doc = *dom
	default:
		return nil, ErrInvalidDataType
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
