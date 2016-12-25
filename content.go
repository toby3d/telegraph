package telegraph

import (
	"bytes"
	"errors"
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

var availableTags = map[string]bool{
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

var availableAttributes = map[string]bool{
	"href": true,
	"src":  true,
}

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
		nodeElement.Tag = strings.ToLower(domNode.Data)
		for i := range domNode.Attr {
			attr := domNode.Attr[i]
			if _, ok := availableAttributes[strings.ToLower(attr.Key)]; ok {
				switch {
				case attr.Key == "src" && strings.Contains(attr.Val, "vimeo.com"):
					nodeElement.Attrs = parseEmbed("vimeo", attr.Val)
				case attr.Key == "src" && (strings.Contains(attr.Val, "youtube.com") || strings.Contains(attr.Val, "youtu.be")):
					nodeElement.Attrs = parseEmbed("youtube", attr.Val)
				case attr.Key == "src" && strings.Contains(attr.Val, "twitter.com"):
					nodeElement.Attrs = parseEmbed("twitter", attr.Val)
				default:
					nodeElement.Attrs = map[string]string{
						strings.ToLower(attr.Key): strings.ToLower(attr.Val),
					}
				}
			}
		}
	}

	for child := domNode.FirstChild; child != nil; child = child.NextSibling {
		nodeElement.Children = append(nodeElement.Children, domToNode(child))
	}

	return nodeElement
}

func parseEmbed(service string, url string) map[string]string {
	return map[string]string{
		"src":               fmt.Sprint("/embed/", service, "?url=", url),
		"width":             "640",
		"height":            "360",
		"frameborder":       "0",
		"allowtransparency": "true",
		"allowfullscreen":   "true",
		"scrolling":         "no",
	}
}
