package telegraph

import (
	"path"
)

type editPage struct {
	// Access token of the Telegraph account.
	AccessToken string `json:"access_token"`

	// Path to the page.
	Path string `json:"path"`

	// Page title.
	Title string `json:"title"`

	// Content of the page.
	Content []Node `json:"content"`

	// Author name, displayed below the article's title.
	AuthorName string `json:"author_name,omitempty"`

	// Profile link, opened when users click on the author's name below the title. Can be any link, not
	// necessarily to a Telegram profile or channel.
	AuthorURL string `json:"author_url,omitempty"`

	// If true, a content field will be returned in the Page object.
	ReturnContent bool `json:"return_content,omitempty"`
}

// EditPage edit an existing Telegraph page. On success, returns a Page object.
func (a *Account) EditPage(update Page, returnContent bool) (*Page, error) {
	data, err := makeRequest(path.Join("editPage", update.Path), editPage{
		AccessToken:   a.AccessToken,
		Path:          update.Path,
		Title:         update.Title,
		Content:       update.Content,
		AuthorName:    update.AuthorName,
		AuthorURL:     update.AuthorURL,
		ReturnContent: returnContent,
	})
	if err != nil {
		return nil, err
	}

	result := new(Page)
	if err = parser.Unmarshal(data, result); err != nil {
		return nil, err
	}

	return result, nil
}
