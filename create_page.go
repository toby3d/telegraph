package telegraph

type createPage struct {
	// Access token of the Telegraph account.
	AccessToken string `json:"access_token"`

	// Page title.
	Title string `json:"title"`

	// Author name, displayed below the article's title.
	AuthorName string `json:"author_name,omitempty"`

	// Profile link, opened when users click on the author's name below the title. Can be any link, not
	// necessarily to a Telegram profile or channel.
	AuthorURL string `json:"author_url,omitempty"`

	// Content of the page.
	Content []Node `json:"content"`

	// If true, a content field will be returned in the Page object.
	ReturnContent bool `json:"return_content,omitempty"`
}

// CreatePage create a new Telegraph page. On success, returns a Page object.
func (a *Account) CreatePage(page Page, returnContent bool) (*Page, error) {
	data, err := makeRequest("createPage", createPage{
		AccessToken:   a.AccessToken,
		Title:         page.Title,
		AuthorName:    page.AuthorName,
		AuthorURL:     page.AuthorURL,
		Content:       page.Content,
		ReturnContent: returnContent,
	})
	if err != nil {
		return nil, err
	}

	result := new(Page)
	if err = parser.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}
