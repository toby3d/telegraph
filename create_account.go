package telegraph

import (
	"fmt"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// CreateAccount create a new Telegraph account. Most users only need one
// account, but this can be useful for channel administrators who would like
// to keep individual author names and profile links for each of their
// channels. On success, returns an Account object with the regular fields and
// an additional access_token field.
func CreateAccount(account *Account) (*Account, error) {
	var args http.Args

	// Account name, helps users with several accounts remember which they are
	// currently using. Displayed to the user above the "Edit/Publish" button
	// on Telegra.ph, other users don't see this name.
	args.Add("short_name", account.ShortName) // required

	// Default author name used when creating new articles.
	args.Add("author_name", account.AuthorName)

	// Default profile link, opened when users click on the author's name
	// below the title. Can be any link, not necessarily to a Telegram profile
	// or channel.
	args.Add("author_url", account.AuthorURL)

	url := fmt.Sprintf(APIEndpoint, "createAccount")
	body, err := request(url, &args)
	if err != nil {
		return nil, err
	}

	var resp Account
	err = json.Unmarshal(*body.Result, &resp)

	return &resp, err
}
