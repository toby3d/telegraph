package telegraph

import (
	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

type revokeAccessTokenParameters struct {
	// Access token of the Telegraph account.
	AccessToken string `json:"access_token"` // required
}

// RevokeAccessToken revoke access_token and generate a new one, for example, if the user would
// like to reset all connected sessions, or you have reasons to believe the token was compromised. On
// success, returns an Account object with new access_token and auth_url fields.
func (a *Account) RevokeAccessToken() (r *Account, err error) {
	args := http.AcquireArgs()
	defer http.ReleaseArgs(args)
	args.Add("access_token", a.AccessToken)

	dst := new(Response)
	dst, err = makeRequest("revokeAccessToken", args)
	if err != nil {
		return
	}

	r = new(Account)
	err = json.Unmarshal(*dst.Result, r)
	return
}
