package telegraph

import (
	"fmt"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

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
	err = json.Unmarshal(*body.Result, &resp)

	return &resp, err
}
