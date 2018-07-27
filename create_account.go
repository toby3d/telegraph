package telegraph

import (
	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// CreateAccount create a new Telegraph account. Most users only need one
// account, but this can be useful for channel administrators who would like to
// keep individual author names and profile links for each of their channels. On
// success, returns an Account object with the regular fields and an additional
// access_token field.
func CreateAccount(account *Account) (r *Account, err error) {
	if account == nil {
		return nil, ErrNoInputData
	}

	args := http.AcquireArgs()
	defer http.ReleaseArgs(args)
	args.Add("short_name", account.ShortName) // required
	args.Add("author_name", account.AuthorName)
	args.Add("author_url", account.AuthorURL)

	dst := new(Response)
	dst, err = makeRequest("createAccount", args)
	if err != nil {
		return
	}

	r = new(Account)
	err = json.Unmarshal(*dst.Result, r)
	return
}
