package telegraph

import (
	http "github.com/valyala/fasthttp"
)

// EditAccountInfo update information about a Telegraph account. Pass only the
// parameters that you want to edit. On success, returns an Account object with
// the default fields.
func (a *Account) EditAccountInfo(update Account) (*Account, error) {
	args := http.AcquireArgs()
	defer http.ReleaseArgs(args)
	args.Add("access_token", a.AccessToken) // required
	args.Add("short_name", update.ShortName)
	args.Add("author_name", update.AuthorName)
	args.Add("author_url", update.AuthorURL)

	data, err := makeRequest("editAccountInfo", args)
	if err != nil {
		return nil, err
	}

	var result Account
	err = parser.Unmarshal(data, &result)
	return &result, err
}
