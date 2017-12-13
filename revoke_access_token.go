package telegraph

import (
	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// RevokeAccessToken revoke access_token and generate a new one, for example,
// if the user would like to reset all connected sessions, or you have reasons
// to believe the token was compromised. On success, returns an Account object
// with new access_token and auth_url fields.
func (account *Account) RevokeAccessToken() (*Account, error) {
	args := http.AcquireArgs()

	// Access token of the Telegraph account.
	args.Add("access_token", account.AccessToken) // required

	body, err := request("revokeAccessToken", "", args)
	if err != nil {
		return nil, err
	}

	var resp Account
	err = json.Unmarshal(*body.Result, &resp)

	return &resp, err
}
