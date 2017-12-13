package telegraph

import (
	"fmt"
	"strings"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

const (
	// FieldShortName used as GetAccountInfo argument for getting account name.
	FieldShortName = "short_name"

	// FieldAuthorName used as GetAccountInfo argument for getting author name.
	FieldAuthorName = "author_name"

	// FieldAuthorURL used as GetAccountInfo argument for getting profile link.
	FieldAuthorURL = "author_url"

	// FieldAuthURL used as GetAccountInfo argument for getting URL to authorize a browser on
	// telegra.ph.
	FieldAuthURL = "auth_url"

	// FieldPageCount used as GetAccountInfo argument for getting number of pages belonging to the
	// Telegraph account.
	FieldPageCount = "page_count"
)

// GetAccountInfo get information about a Telegraph account. Returns an Account object on success.
func (account *Account) GetAccountInfo(fields ...string) (*Account, error) {
	args := http.AcquireArgs()

	// Access token of the Telegraph account.
	args.Add("access_token", account.AccessToken) // required

	// List of account fields to return. Available fields: short_name,
	// author_name, author_url, auth_url, page_count.
	args.Add("fields", fmt.Sprint(`["`, strings.Join(fields, `","`), `"]`))

	body, err := request("getAccountInfo", "", args)
	if err != nil {
		return nil, err
	}

	var resp Account
	err = json.Unmarshal(*body.Result, &resp)

	return &resp, err
}
