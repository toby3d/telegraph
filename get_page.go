package telegraph

import (
	gopath "path"
	"strconv"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// GetPage get a Telegraph page. Returns a Page object on success.
func GetPage(path string, returnContent bool) (r *Page, err error) {
	args := http.AcquireArgs()
	defer http.ReleaseArgs(args)
	args.Add("path", path) // required
	args.Add("return_content", strconv.FormatBool(returnContent))

	dst := new(Response)
	dst, err = makeRequest(gopath.Join("getPage", path), args)
	if err != nil {
		return nil, err
	}

	r = new(Page)
	err = json.Unmarshal(*dst.Result, r)
	return
}
