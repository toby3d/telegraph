package telegraph

import (
	"path"
	"strconv"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// EditPage edit an existing Telegraph page. On success, returns a Page object.
func (a *Account) EditPage(update *Page, returnContent bool) (r *Page, err error) {
	if update == nil {
		return nil, ErrNoInputData
	}

	var src []byte
	src, err = json.Marshal(update.Content)
	if err != nil {
		return
	}

	args := http.AcquireArgs()
	defer http.ReleaseArgs(args)
	args.Add("access_token", a.AccessToken) // required
	args.Add("path", update.Path)           // required
	args.Add("title", update.Title)         // required
	args.Add("content", string(src))        // required
	args.Add("author_name", a.AuthorName)
	args.Add("author_url", a.AuthorURL)
	args.Add("return_content", strconv.FormatBool(returnContent))

	dst := new(Response)
	dst, err = makeRequest(path.Join("editPage", update.Path), args)
	if err != nil {
		return nil, err
	}

	r = new(Page)
	err = json.Unmarshal(*dst.Result, r)
	return
}
