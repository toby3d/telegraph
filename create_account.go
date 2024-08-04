package telegraph

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// CreateAccount create a new Telegraph account. Most users only need one
// account, but this can be useful for channel administrators who would like to
// keep individual author names and profile links for each of their channels. On
// success, returns an [Account] object with the regular fields and an
// additional access_token field.
type CreateAccount struct {
	// Default profile link, opened when users click on the author's name
	// below the title. Can be any link, not necessarily to a Telegram
	// profile or channel.
	AuthorURL *URL `json:"author_url,omitempty"` // 0-512 characters

	// Default author name used when creating new articles.
	AuthorName *AuthorName `json:"author_name,omitempty"` // 0-128 characters

	// Required.
	ShortName ShortName `json:"short_name"` // 1-32 characters
}

func (params CreateAccount) Do(ctx context.Context, client *http.Client) (*Account, error) {
	data, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("createAccount: cannot marshal request parameters: %w", err)
	}

	return post[*Account](ctx, client, bytes.NewReader(data), "createAccount")
}