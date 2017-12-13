package telegraph

import (
	"strconv"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// PageList represents a list of Telegraph articles belonging to an account. Most recently created
// articles first.
type PageList struct {
	// Total number of pages belonging to the target Telegraph account.
	TotalCount int `json:"total_count"`

	// Requested pages of the target Telegraph account.
	Pages []*Page `json:"pages"`
}

// GetPageList get a list of pages belonging to a Telegraph account. Returns a PageList object, sorted
// by most recently created pages first.
func (account *Account) GetPageList(offset, limit int) (*PageList, error) {
	args := http.AcquireArgs()

	// Access token of the Telegraph account.
	args.Add("access_token", account.AccessToken) // required

	// Sequential number of the first page to be returned.
	args.Add("offset", strconv.Itoa(offset))

	// Limits the number of pages to be retrieved.
	args.Add("limit", strconv.Itoa(limit))

	body, err := request("getPageList", "", args)
	if err != nil {
		return nil, err
	}

	var resp PageList
	err = json.Unmarshal(*body.Result, &resp)

	return &resp, err
}
