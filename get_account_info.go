package telegraph

import (
	"strings"

	http "github.com/valyala/fasthttp"
)

// GetAccountInfo get information about a Telegraph account. Returns an Account object on success.
func (a *Account) GetAccountInfo(fields ...string) (*Account, error) {
	args := http.AcquireArgs()
	defer http.ReleaseArgs(args)
	args.Add("access_token", a.AccessToken) // required
	if len(fields) > 0 {
		args.Add("fields", `["`+strings.Join(fields, `","`)+`"]`)
	}

	data, err := makeRequest("getAccountInfo", args)
	if err != nil {
		return nil, err
	}

	var result Account
	err = parser.Unmarshal(data, &result)
	return &result, err
}
