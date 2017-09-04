package telegraph

import (
	"fmt"
	"strconv"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// CreatePage create a new Telegraph page. On success, returns a Page object.
func (account *Account) CreatePage(page *Page, returnContent bool) (*Page, error) {
	var args http.Args

	// Access token of the Telegraph account.
	args.Add("access_token", account.AccessToken) // required

	// Page title.
	args.Add("title", page.Title) // required

	if page.AuthorName != "" {
		// Author name, displayed below the article's title.
		args.Add("author_name", page.AuthorName)
	}

	if page.AuthorURL != "" {
		// Profile link, opened when users click on the author's name below
		// the title. Can be any link, not necessarily to a Telegram profile
		// or channel.
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

	url := fmt.Sprintf(APIEndpoint, "createPage")
	body, err := request(url, &args)
	if err != nil {
		return nil, err
	}

	var resp Page
	err = json.Unmarshal(*body.Result, &resp)

	return &resp, err
}
