package telegraph_test

import (
	"bytes"
	"encoding/json"
	"net/url"
	"strings"
	"testing"

	"source.toby3d.me/toby3d/telegraph/v2"
)

func TestURL_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	result := &struct {
		URL telegraph.URL `json:"url"`
	}{}
	if err := json.NewDecoder(strings.NewReader(`{"url":"https://example.com/"}`)).
		Decode(result); err != nil {
		t.Fatal(err)
	}

	const expect string = "https://example.com/"
	if actual := result.URL.String(); actual != expect {
		t.Errorf("got '%s', want '%s'", actual, expect)
	}
}

func TestURL_MarshalJSON(t *testing.T) {
	t.Parallel()

	result := bytes.NewBuffer(nil)
	if err := json.NewEncoder(result).
		Encode(&struct {
			URL telegraph.URL `json:"url"`
		}{URL: telegraph.URL{URL: &url.URL{
			Scheme: "https",
			Host:   "example.com",
			Path:   "/",
		}}}); err != nil {
		t.Fatal(err)
	}

	const expect string = "{\"url\":\"https://example.com/\"}\n"
	if actual := result.String(); actual != expect {
		t.Errorf("got '%s', want '%s'", actual, expect)
	}
}