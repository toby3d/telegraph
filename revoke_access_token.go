package telegraph

type revokeAccessToken struct {
	// Access token of the Telegraph account.
	AccessToken string `json:"access_token"`
}

// RevokeAccessToken revoke access_token and generate a new one, for example, if the user would
// like to reset all connected sessions, or you have reasons to believe the token was compromised. On
// success, returns an Account object with new access_token and auth_url fields.
func (a *Account) RevokeAccessToken() (*Account, error) {
	resp, err := makeRequest("revokeAccessToken", revokeAccessToken{
		AccessToken: a.AccessToken,
	})
	if err != nil {
		return nil, err
	}

	account := new(Account)
	if err = parser.Unmarshal(resp, &account); err != nil {
		return nil, err
	}

	return account, nil
}
