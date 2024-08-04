package telegraph

import "strconv"

// Node represent abstract object represents a DOM Node. It can be a String
// which represents a DOM text node or a [NodeElement] object.
type Node struct {
	Element *NodeElement `json:"-"`
	Text    string       `json:"-"`
}

func (n Node) MarshalJSON() ([]byte, error) {
	if n.Text != "" {
		return []byte(strconv.Quote(n.Text)), nil
	}

	return nil, nil
}