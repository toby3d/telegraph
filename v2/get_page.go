package telegraph

import (
	"context"
	"net/http"
	"net/url"
)

// GetPage get a Telegraph page. Returns a [Page] object on success.
type GetPage struct {
	// Required. Path to the Telegraph page (in the format Title-12-31, i.e.
	// everything that comes after http://telegra.ph/).
	Path string `json:"path"`

	// If true, content field will be returned in [Page] object.
	ReturnContent bool `json:"return_content,omitempty"`
}

func (params GetPage) Do(ctx context.Context, client *http.Client) (*Page, error) {
	data := make(url.Values)

	if params.ReturnContent {
		data.Set("return_content", "true")
	}

	return get[*Page](ctx, client, data, "getPage", params.Path)
}