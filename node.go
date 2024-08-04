package telegraph

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Node represent abstract object represents a DOM Node. It can be a String
// which represents a DOM text node or a [NodeElement] object.
type Node struct {
	Element *NodeElement `json:"-"`
	Text    string       `json:"-"`
}

func (n *Node) UnmarshalJSON(v []byte) error {
	switch v[0] {
	case '{':
		nodeElement := new(NodeElement)
		if err := json.Unmarshal(v, nodeElement); err != nil {
			return fmt.Errorf("Node: UnmarshalJSON: cannot unmarshal NodeElement: %w", err)
		}

		n.Element = nodeElement
	case '"':
		unquoted, err := strconv.Unquote(string(v))
		if err != nil {
			return fmt.Errorf("Node: UnmarshalJSON: cannot unquote string: %w", err)
		}

		n.Text = unquoted
	}

	return nil
}

func (n Node) MarshalJSON() ([]byte, error) {
	switch {
	default:
		return nil, nil
	case n.Text != "":
		return []byte(strconv.Quote(n.Text)), nil
	case n.Element != nil:
		result, err := json.Marshal(n.Element)
		if err != nil {
			return nil, fmt.Errorf("Node: MarshalJSON: cannot encode as Element: %w", err)
		}

		return result, nil
	}
}

func (n Node) String() string {
	switch {
	default:
		return ""
	case n.Text != "":
		return n.Text
	case n.Element != nil:
		return n.Element.String()
	}
}

func (n Node) GoString() string {
	if n.Text == "" && n.Element == nil {
		return "telegraph.Node(und)"
	}

	return "telegraph.Node(" + n.String() + ")"
}