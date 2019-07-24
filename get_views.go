package telegraph

import (
	gopath "path"
	"strconv"

	http "github.com/valyala/fasthttp"
)

// GetViews get the number of views for a Telegraph article. By default, the total number of page
// views will be returned. Returns a PageViews object on success.
func GetViews(path string, date ...int) (*PageViews, error) {
	args := http.AcquireArgs()
	defer http.ReleaseArgs(args)
	args.Add("path", path) // required
	if len(date) > 0 {
		args.Add("year", strconv.Itoa(date[0]))
	}
	if len(date) > 1 {
		args.Add("month", strconv.Itoa(date[1]))
	}
	if len(date) > 2 {
		args.Add("day", strconv.Itoa(date[2]))
	}
	if len(date) > 3 {
		args.Add("hour", strconv.Itoa(date[3]))
	}

	data, err := makeRequest(gopath.Join("getViews", path), args)
	if err != nil {
		return nil, err
	}

	var result PageViews
	err = parser.Unmarshal(data, &result)
	return &result, err
}
