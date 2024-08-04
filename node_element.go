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