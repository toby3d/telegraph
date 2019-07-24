package telegraph

import (
	http "github.com/valyala/fasthttp"
)

// CreateAccount create a new Telegraph account. Most users only need one
// account, but this can be useful for channel administrators who would like to
// keep individual author names and profile links for each of their channels. On
// success, returns an Account object with the regular fields and an additional
// access_token field.
func CreateAccount(account Account) (*Account, error) {
	args := http.AcquireArgs()
	defer http.ReleaseArgs(args)
	args.Add("short_name", account.ShortName) // required
	args.Add("author_name", account.AuthorName)
	args.Add("author_url", account.AuthorURL)

	data, err := makeRequest("createAccount", args)
	if err != nil {
		return nil, err
	}

	var result Account
	err = parser.Unmarshal(data, &result)
	return &result, err
}
