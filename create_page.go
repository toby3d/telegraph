package telegraph

import (
	"strconv"

	http "github.com/valyala/fasthttp"
)

// CreatePage create a new Telegraph page. On success, returns a Page object.
func (a *Account) CreatePage(page Page, returnContent bool) (*Page, error) {
	src, err := parser.Marshal(page.Content)
	if err != nil {
		return nil, err
	}

	args := http.AcquireArgs()
	defer http.ReleaseArgs(args)
	args.Add("access_token", a.AccessToken) // required
	args.Add("title", page.Title)           // required
	args.Add("author_name", page.AuthorName)
	args.Add("author_url", page.AuthorURL)
	args.AddBytesV("content", src)
	args.Add("return_content", strconv.FormatBool(returnContent))

	data, err := makeRequest("createPage", args)
	if err != nil {
		return nil, err
	}

	var result Page
	err = parser.Unmarshal(data, &result)
	return &result, err
}
