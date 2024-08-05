package telegraph_test

import (
	"bytes"
	"encoding/json"
	"net/url"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"source.toby3d.me/toby3d/telegraph/v2"
)

func TestNode_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	result := make([]telegraph.NodeElement, 0)
	if err := json.NewDecoder(strings.NewReader(`[{"tag":"p","children":["Hello, world!"]}]`)).
		Decode(&result); err != nil {
		t.Fatal(err)
	}

	expect := []telegraph.NodeElement{{
		Tag:      telegraph.P,
		Children: []telegraph.Node{{Text: "Hello, world!"}},
	}}

	if diff := cmp.Diff(expect, result, cmpopts.EquateComparable(telegraph.Tag{})); diff != "" {
		t.Error(diff)
	}
}

func TestNode_MarshalJSON(t *testing.T) {
	t.Parallel()

	result := bytes.NewBuffer(nil)
	if err := json.NewEncoder(result).
		Encode(&struct {
			URL telegraph.URL `json:"url"`
		}{
			URL: telegraph.URL{URL: &url.URL{
				Scheme: "https",
				Host:   "example.com",
				Path:   "/",
			}},
		}); err != nil {
		t.Fatal(err)
	}

	const expect string = "{\"url\":\"https://example.com/\"}\n"
	if got := result.String(); got != expect {
		t.Errorf("got '%s', want '%s'", got, expect)
	}
}