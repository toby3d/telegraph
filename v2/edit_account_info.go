package telegraph

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// EditAccountInfo update information about a Telegraph account. Pass only the
// parameters that you want to edit. On success, returns an [Account] object
// with the default fields.
type EditAccountInfo struct {
	// New default profile link, opened when users click on the author's
	// name below the title. Can be any link, not necessarily to a Telegram
	// profile or channel.
	AuthorURL *URL `json:"author_url,omitempty"` // 0-512 characters

	// New account name.
	ShortName *ShortName `json:"short_name,omitempty"` // 1-32 characters

	// New default author name used when creating new articles.
	AuthorName *AuthorName `json:"author_name,omitempty"` // 0-128 characters

	// Required. Access token of the Telegraph account.
	AccessToken string `json:"access_token"`
}

func (params EditAccountInfo) Do(ctx context.Context, client *http.Client) (*Account, error) {
	data, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("editAccountInfo: cannot marshal request parameters: %w", err)
	}

	return post[*Account](ctx, client, bytes.NewReader(data), "editAccountInfo")
}