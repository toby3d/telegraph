package telegraph_test

import (
	"encoding/json"
	"testing"

	"source.toby3d.me/toby3d/telegraph"
)

func TestNodeElement_MarshalJSON(t *testing.T) {
	t.Parallel()

	result, err := json.Marshal(telegraph.NodeElement{
		Tag:      telegraph.A,
		Attrs:    &telegraph.Attributes{Href: "https://example.com/"},
		Children: []telegraph.Node{{Text: "Sample Link"}},
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(result))
}