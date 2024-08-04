package telegraph

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
)

// GetPageList get a list of pages belonging to a Telegraph account. Returns a
// [PageList] object, sorted by most recently created pages first.
type GetPageList struct {
	// Required. Access token of the Telegraph account.
	AccessToken string `json:"access_token"`

	// Sequential number of the first page to be returned.
	Offset uint `json:"offset,omitempty"` // 0

	// Limits the number of pages to be retrieved.
	Limit uint16 `json:"limit,omitempty"` // 50 (0-200)
}

func (params GetPageList) Do(ctx context.Context, client *http.Client) (*PageList, error) {
	data := make(url.Values)
	data.Set("access_token", params.AccessToken)
	data.Set("offset", strconv.FormatUint(uint64(params.Offset), 10))

	if params.Limit <= 200 {
		data.Set("limit", strconv.FormatUint(uint64(params.Limit), 10))
	}

	return get[*PageList](ctx, client, data, "getPageList")
}