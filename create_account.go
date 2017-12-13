package telegraph

import (
	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// Account represents a Telegraph account.
type Account struct {
	// Only returned by the createAccount and revokeAccessToken method. Access token of the Telegraph
	// account.
	AccessToken string `json:"access_token"` // optional

	// URL to authorize a browser on telegra.ph and connect it to a Telegraph account. This URL is
	// valid for only one use and for 5 minutes only.
	AuthURL string `json:"auth_url,omitempty"` // optional

	// Account name, helps users with several accounts remember which they are currently using.
	// Displayed to the user above the "Edit/Publish"  button on Telegra.ph, other users don't see
	// this name.
	ShortName string `json:"short_name"`

	// Default author name used when creating new articles.
	AuthorName string `json:"author_name"`

	// Profile link, opened when users click on the author's name below the title. Can be any link,
	// not necessarily to a Telegram profile or channel.
	AuthorURL string `json:"author_url"`

	// Number of pages belonging to the Telegraph account.
	PageCount int `json:"page_count,omitempty"` // optional
}

// CreateAccount create a new Telegraph account. Most users only need one account, but this can be
// useful for channel administrators who would like to keep individual author names and profile links
// for each of their channels. On success, returns an Account object with the regular fields and an
// additional access_token field.
func CreateAccount(account *Account) (*Account, error) {
	args := http.AcquireArgs()

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

	body, err := request("createAccount", "", args)
	if err != nil {
		return nil, err
	}

	var resp Account
	err = json.Unmarshal(*body.Result, &resp)

	return &resp, err
}
