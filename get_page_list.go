package telegraph

import (
	"strconv"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// GetPageList get a list of pages belonging to a Telegraph account. Returns a PageList object, sorted
// by most recently created pages first.
func (a *Account) GetPageList(offset, limit int) (r *PageList, err error) {
	args := http.AcquireArgs()
	defer http.ReleaseArgs(args)
	args.Add("access_token", a.AccessToken) // required
	args.Add("offset", strconv.Itoa(offset))
	args.Add("limit", strconv.Itoa(limit))

	dst := new(Response)
	dst, err = makeRequest("getPageList", args)
	if err != nil {
		return nil, err
	}

	r = new(PageList)
	err = json.Unmarshal(*dst.Result, r)
	return
}
