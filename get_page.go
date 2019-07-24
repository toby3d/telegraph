package telegraph

import (
	gopath "path"
	"strconv"

	http "github.com/valyala/fasthttp"
)

// GetPage get a Telegraph page. Returns a Page object on success.
func GetPage(path string, returnContent bool) (*Page, error) {
	args := http.AcquireArgs()
	defer http.ReleaseArgs(args)
	args.Add("path", path) // required
	args.Add("return_content", strconv.FormatBool(returnContent))

	data, err := makeRequest(gopath.Join("getPage", path), args)
	if err != nil {
		return nil, err
	}

	var result Page
	err = parser.Unmarshal(data, &result)
	return &result, err
}
