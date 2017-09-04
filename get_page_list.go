package telegraph

import (
	"fmt"
	"strconv"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// GetPageList get a list of pages belonging to a Telegraph account. Returns
// a PageList object, sorted by most recently created pages first.
func (account *Account) GetPageList(offset, limit int) (*PageList, error) {
	var args http.Args

	// Access token of the Telegraph account.
	args.Add("access_token", account.AccessToken) // required

	// Sequential number of the first page to be returned.
	args.Add("offset", strconv.Itoa(offset))

	// Limits the number of pages to be retrieved.
	args.Add("limit", strconv.Itoa(limit))

	url := fmt.Sprintf(APIEndpoint, "getPageList")
	body, err := request(url, &args)
	if err != nil {
		return nil, err
	}

	var resp PageList
	err = json.Unmarshal(*body.Result, &resp)

	return &resp, err
}
