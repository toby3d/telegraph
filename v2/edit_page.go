package telegraph

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// EditPage edit an existing Telegraph page. On success, returns a [Page]
// object.
type EditPage struct {
	// Profile link, opened when users click on the author's name below the
	// title. Can be any link, not necessarily to a Telegram profile or
	// channel.
	AuthorURL *URL `json:"author_url,omitempty"` // 0-512 characters

	// Required. Access token of the Telegraph account.
	AccessToken string `json:"access_token"`

	// Required. Path to the page.
	Path string `json:"path"`

	// Required. Page title.
	Title Title `json:"title"` // 1-256 characters

	// Author name, displayed below the article's title.
	AuthorName *AuthorName `json:"author_name,omitempty"` // 0-128 characters

	// Required. [Content] of the page.
	Content []Node `json:"content"` // up to 64 KB

	// If true, a content field will be returned in the [Page] object.
	ReturnContent bool `json:"return_content,omitempty"` // false
}

func (params EditPage) Do(ctx context.Context, client *http.Client) (*Page, error) {
	data, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("editPage: cannot marshal request parameters: %w", err)
	}

	return post[*Page](ctx, client, bytes.NewReader(data), "editPage", params.Path)
}