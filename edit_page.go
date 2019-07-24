package telegraph

import (
	"path"
	"strconv"

	http "github.com/valyala/fasthttp"
)

// EditPage edit an existing Telegraph page. On success, returns a Page object.
func (a *Account) EditPage(update Page, returnContent bool) (*Page, error) {
	src, err := parser.Marshal(update.Content)
	if err != nil {
		return nil, err
	}

	args := http.AcquireArgs()
	defer http.ReleaseArgs(args)
	args.Add("access_token", a.AccessToken) // required
	args.Add("path", update.Path)           // required
	args.Add("title", update.Title)         // required
	args.AddBytesV("content", src)          // required
	args.Add("author_name", update.AuthorName)
	args.Add("author_url", update.AuthorURL)
	args.Add("return_content", strconv.FormatBool(returnContent))

	data, err := makeRequest(path.Join("editPage", update.Path), args)
	if err != nil {
		return nil, err
	}

	var result Page
	err = parser.Unmarshal(data, &result)
	return &result, err
}
