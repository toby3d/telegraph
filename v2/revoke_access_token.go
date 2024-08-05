package telegraph

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// RevokeAccessToken revoke access_token and generate a new one, for example, if
// the user would like to reset all connected sessions, or you have reasons to
// believe the token was compromised. On success, returns an [Account] object
// with new access_token and auth_url fields.
type RevokeAccessToken struct {
	// Access token of the Telegraph account.
	AccessToken string `json:"access_token"`
}

func (params RevokeAccessToken) Do(ctx context.Context, client *http.Client) (*Account, error) {
	data, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("revokeAccessToken: cannot marshal request parameters: %w", err)
	}

	return post[*Account](ctx, client, bytes.NewReader(data), "revokeAccessToken")
}