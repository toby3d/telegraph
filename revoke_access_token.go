package telegraph

import (
	http "github.com/valyala/fasthttp"
)

// RevokeAccessToken revoke access_token and generate a new one, for example, if the user would
// like to reset all connected sessions, or you have reasons to believe the token was compromised. On
// success, returns an Account object with new access_token and auth_url fields.
func (a *Account) RevokeAccessToken() (*Account, error) {
	args := http.AcquireArgs()
	defer http.ReleaseArgs(args)
	args.Add("access_token", a.AccessToken)

	resp, err := makeRequest("revokeAccessToken", args)
	if err != nil {
		return nil, err
	}

	var account Account
	err = parser.Unmarshal(resp, &account)
	return &account, err
}
