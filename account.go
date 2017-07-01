package telegraph

import (
	"fmt"
	"strings"

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
	if err := json.Unmarshal(body.Result, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// EditAccountInfo update information about a Telegraph account. Pass only the
// parameters that you want to edit. On success, returns an Account object
// with the default fields.
func (account *Account) EditAccountInfo(update *Account) (*Account, error) {
	var args http.Args

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

	url := fmt.Sprintf(APIEndpoint, "editAccountInfo")
	body, err := request(url, &args)
	if err != nil {
		return nil, err
	}

	var resp Account
	if err := json.Unmarshal(body.Result, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetAccountInfo get information about a Telegraph account. Returns an
// Account object on success.
func (account *Account) GetAccountInfo(fields ...string) (*Account, error) {
	var args http.Args

	// Access token of the Telegraph account.
	args.Add("access_token", account.AccessToken) // required

	// List of account fields to return. Available fields: short_name,
	// author_name, author_url, auth_url, page_count.
	args.Add("fields", fmt.Sprintf(`["%s"]`, strings.Join(fields, `","`)))

	url := fmt.Sprintf(APIEndpoint, "getAccountInfo")
	body, err := request(url, &args)
	if err != nil {
		return nil, err
	}

	var resp Account
	if err := json.Unmarshal(body.Result, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RevokeAccessToken revoke access_token and generate a new one, for example,
// if the user would like to reset all connected sessions, or you have reasons
// to believe the token was compromised. On success, returns an Account object
// with new access_token and auth_url fields.
func (account *Account) RevokeAccessToken() (*Account, error) {
	var args http.Args

	// Access token of the Telegraph account.
	args.Add("access_token", account.AccessToken) // required

	url := fmt.Sprintf(APIEndpoint, "revokeAccessToken")
	body, err := request(url, &args)
	if err != nil {
		return nil, err
	}

	var resp Account
	if err := json.Unmarshal(body.Result, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
