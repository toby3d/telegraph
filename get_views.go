package telegraph

import (
	gopath "path"
	"strconv"
	"time"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// GetViews get the number of views for a Telegraph article. By default, the total number of page
// views will be returned. Returns a PageViews object on success.
func GetViews(path string, date time.Time) (r *PageViews, err error) {
	args := http.AcquireArgs()
	defer http.ReleaseArgs(args)
	args.Add("path", path) // required
	args.Add("year", strconv.Itoa(date.Year()))
	args.Add("month", strconv.Itoa(int(date.Month())))
	args.Add("day", strconv.Itoa(date.Day()))
	args.Add("hour", strconv.Itoa(date.Hour()))

	dst := new(Response)
	dst, err = makeRequest(gopath.Join("getViews", path), args)
	if err != nil {
		return nil, err
	}

	r = new(PageViews)
	err = json.Unmarshal(*dst.Result, r)
	return
}
