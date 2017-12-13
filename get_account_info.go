package telegraph

import (
	"fmt"
	"strings"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// GetAccountInfo get information about a Telegraph account. Returns an
// Account object on success.
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
