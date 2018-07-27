package telegraph

import (
	"strconv"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// CreatePage create a new Telegraph page. On success, returns a Page object.
func (a *Account) CreatePage(page *Page, returnContent bool) (r *Page, err error) {
	if page == nil {
		return nil, ErrNoInputData
	}

	var src []byte
	src, err = json.Marshal(page.Content)
	if err != nil {
		return
	}

	args := http.AcquireArgs()
	defer http.ReleaseArgs(args)
	args.Add("access_token", a.AccessToken) // required
	args.Add("title", page.Title)           // required
	args.Add("author_name", a.AuthorName)
	args.Add("author_url", a.AuthorURL)
	args.Add("content", string(src))
	args.Add("return_content", strconv.FormatBool(returnContent))

	dst := new(Response)
	dst, err = makeRequest("createPage", args)
	if err != nil {
		return nil, err
	}

	r = new(Page)
	err = json.Unmarshal(*dst.Result, r)
	return
}
