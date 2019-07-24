package telegraph

import (
	"strconv"

	http "github.com/valyala/fasthttp"
)

// GetPageList get a list of pages belonging to a Telegraph account. Returns a PageList object, sorted
// by most recently created pages first.
func (a *Account) GetPageList(offset, limit int) (*PageList, error) {
	args := http.AcquireArgs()
	defer http.ReleaseArgs(args)
	args.Add("access_token", a.AccessToken) // required
	if offset > 0 {
		args.Add("offset", strconv.Itoa(offset))
	}
	if limit > 0 {
		args.Add("limit", strconv.Itoa(limit))
	}

	data, err := makeRequest("getPageList", args)
	if err != nil {
		return nil, err
	}

	var result PageList
	err = parser.Unmarshal(data, &result)
	return &result, err
}
