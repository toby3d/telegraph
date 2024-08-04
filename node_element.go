package telegraph

// NodeElement object represents a DOM element node.
type NodeElement struct {
	// Optional. Attributes of the DOM element. Key of object represents
	// name of attribute, value represents value of attribute. Available
	// attributes: href, src.
	Attrs *Attributes `json:"attrs,omitempty"`

	// Optional. List of child nodes for the DOM element.
	Children []Node `json:"children,omitempty"`

	// Name of the DOM element. Available tags: a, aside, b, blockquote, br,
	// code, em, figcaption, figure, h3, h4, hr, i, iframe, img, li, ol, p,
	// pre, s, strong, u, ul, video.
	Tag Tag `json:"tag"`
}

func NewNodeElement(tag Tag) *NodeElement {
	return &NodeElement{
		Attrs:    new(Attributes),
		Children: make([]Node, 0),
		Tag:      tag,
	}
}

func (ne NodeElement) String() string {
	result := "<" + ne.Tag.String()

	if ne.Attrs != nil {
		switch {
		case ne.Attrs.Href != "":
			result += ` href="` + ne.Attrs.Href + `"`
		case ne.Attrs.Src != "":
			result += ` src="` + ne.Attrs.Src + `"`
		}
	}

	if len(ne.Children) == 0 {
		return result + " />"
	}

	result += ">"

	for _, n := range ne.Children {
		result += n.String()
	}

	return result + "</" + ne.Tag.String() + ">"
}

func (ne NodeElement) GoString() string {
	return "telegraph.NodeElement(" + ne.String() + ")"
}