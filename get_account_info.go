package telegraph

import (
	"strings"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// GetAccountInfo get information about a Telegraph account. Returns an Account object on success.
func (a *Account) GetAccountInfo(fields ...string) (r *Account, err error) {
	args := http.AcquireArgs()
	defer http.ReleaseArgs(args)
	args.Add("access_token", a.AccessToken) // required
	args.Add("fields", `["`+strings.Join(fields, `","`)+`"]`)

	dst := new(Response)
	dst, err = makeRequest("getAccountInfo", args)
	if err != nil {
		return nil, err
	}

	r = new(Account)
	err = json.Unmarshal(*dst.Result, r)
	return
}
