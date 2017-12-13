package telegraph

import (
	"strconv"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// Page represents a page on Telegraph.
type Page struct {
	// Path to the page.
	Path string `json:"path"`

	// URL of the page.
	URL string `json:"url"`

	// Title of the page.
	Title string `json:"title"`

	// Description of the page.
	Description string `json:"description"`

	// Name of the author, displayed below the title.
	AuthorName string `json:"author_name,omitempty"` // optional

	// Profile link, opened when users click on the author's name below the title. Can be any link,
	// not necessarily to a Telegram profile or channel.
	AuthorURL string `json:"author_url,omitempty"` // optional

	// Image URL of the page.
	ImageURL string `json:"image_url,omitempty"` // optional

	// Content of the page.
	Content []Node `json:"content,omitempty"` // optional

	// Number of page views for the page.
	Views int `json:"views"`

	// Only returned if access_token passed. True, if the target Telegraph account can edit the page.
	CanEdit bool `json:"can_edit,omitempty"` // optional
}

// CreatePage create a new Telegraph page. On success, returns a Page object.
func (account *Account) CreatePage(page *Page, returnContent bool) (*Page, error) {
	args := http.AcquireArgs()

	// Access token of the Telegraph account.
	args.Add("access_token", account.AccessToken) // required

	// Page title.
	args.Add("title", page.Title) // required

	if page.AuthorName != "" {
		// Author name, displayed below the article's title.
		args.Add("author_name", page.AuthorName)
	}

	if page.AuthorURL != "" {
		// Profile link, opened when users click on the author's name below the title. Can be any
		// link, not necessarily to a Telegram profile or channel.
		args.Add("author_url", page.AuthorURL)
	}

	// If true, a content field will be returned in the Page object.
	args.Add("return_content", strconv.FormatBool(returnContent))

	content, err := json.Marshal(page.Content)
	if err != nil {
		return nil, err
	}

	// Content of the page.
	args.Add("content", string(content)) // required

	body, err := request("createPage", "", args)
	if err != nil {
		return nil, err
	}

	var resp Page
	err = json.Unmarshal(*body.Result, &resp)

	return &resp, err
}
