package telegraph

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// GetAccountInfo get information about a Telegraph account. Returns an
// [Account] object on success.
type GetAccountInfo struct {
	// Required. Access token of the Telegraph account.
	AccessToken string `json:"access_token"`

	// List of account fields to return.
	Fields []AccountField `json:"fields,omitempty"` // ["short_name","author_name","author_url"]
}

func (params GetAccountInfo) Do(ctx context.Context, client *http.Client) (*Account, error) {
	data := make(url.Values)
	params.populate(data)

	return get[*Account](ctx, client, data, "getAccountInfo")
}

func (p GetAccountInfo) populate(dst url.Values) {
	dst.Set("access_token", p.AccessToken)

	if len(p.Fields) == 0 {
		return
	}

	values := make([]string, 0, len(p.Fields))

	for i := range p.Fields {
		if p.Fields[i].accountField == "" {
			continue
		}

		values = append(values, strconv.Quote(p.Fields[i].accountField))
	}

	dst.Set("fields", "["+strings.Join(values, ",")+"]")
}