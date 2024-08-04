package telegraph_test

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"

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

	const expect string = `{"attrs":{"href":"https://example.com/"},"children":["Sample Link"],"tag":"a"}`
	if diff := cmp.Diff(expect, string(result)); diff != "" {
		t.Error(diff)
	}
}

func TestNodeElement_String(t *testing.T) {
	t.Parallel()

	f := func(name string, input telegraph.NodeElement, expect string) {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if actual := input.String(); actual != expect {
				t.Errorf("want '%s', got '%s'", expect, actual)
			}
		})
	}

	f("p", telegraph.NodeElement{
		Children: []telegraph.Node{{Text: "Hello, World!"}},
		Tag:      telegraph.P,
	}, `<p>Hello, World!</p>`)
	f("img", telegraph.NodeElement{
		Attrs: &telegraph.Attributes{Src: "https://example.com/photo.jpg"},
		Tag:   telegraph.Img,
	}, `<img src="https://example.com/photo.jpg" />`)
	f("aimg", telegraph.NodeElement{
		Attrs: &telegraph.Attributes{Href: "https://example.com/"},
		Children: []telegraph.Node{{
			Element: &telegraph.NodeElement{
				Attrs: &telegraph.Attributes{Src: "https://example.com/photo.jpg"},
				Tag:   telegraph.Img,
			},
		}},
		Tag: telegraph.A,
	}, `<a href="https://example.com/"><img src="https://example.com/photo.jpg" /></a>`)
}