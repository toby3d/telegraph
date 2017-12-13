package telegraph

import (
	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// EditAccountInfo update information about a Telegraph account. Pass only the
// parameters that you want to edit. On success, returns an Account object
// with the default fields.
func (account *Account) EditAccountInfo(update *Account) (*Account, error) {
	args := http.AcquireArgs()

	// Access token of the Telegraph account.
	args.Add("access_token", account.AccessToken) // required

	// New account name.
	args.Add("short_name", update.ShortName)

	// New default author name used when creating new articles.
	args.Add("author_name", update.AuthorName)

	// New default profile link, opened when users click on the author's name
	// below the title. Can be any link, not necessarily to a Telegram profile
	// or channel.
	args.Add("author_url", update.AuthorURL)

	body, err := request("editAccountInfo", "", args)
	if err != nil {
		return nil, err
	}

	var resp Account
	err = json.Unmarshal(*body.Result, &resp)

	return &resp, err
}
