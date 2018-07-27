package telegraph

import (
	gojson "encoding/json"
	"errors"
	"net/url"

	"github.com/kirillDanshin/dlog"
	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// Response contains a JSON object, which always has a Boolean field ok. If ok
// equals true, the request was successful, and the result of the query can be
// found in the result field. In case of an unsuccessful request, ok equals
// false, and the error is explained in the error field (e.g.
// SHORT_NAME_REQUIRED).
type Response struct {
	Ok     bool               `json:"ok"`
	Error  string             `json:"error"`
	Result *gojson.RawMessage `json:"result,omitempty"`
}

var defaultURL = url.URL{
	Scheme: "https",
	Host:   "api.telegra.ph",
}

func makeRequest(path string, args *http.Args) (r *Response, err error) {
	requestURL := defaultURL
	requestURL.Path = path
	requestURL.RawQuery = args.String()

	req := http.AcquireRequest()
	defer http.ReleaseRequest(req)
	req.Header.SetMethod("GET")
	req.SetRequestURI(requestURL.String())
	req.Header.SetUserAgent("toby3d/telegraph")
	req.Header.SetContentType("application/json;charset=utf-8")

	dlog.Ln("request:")
	dlog.D(req)

	resp := http.AcquireResponse()
	defer http.ReleaseResponse(resp)
	err = http.Do(req, resp)
	if err != nil {
		dlog.Ln(err.Error())
		return
	}

	dlog.Ln("response:")
	dlog.D(resp)

	r = new(Response)
	if err = json.Unmarshal(resp.Body(), r); err != nil {
		return
	}

	if !r.Ok {
		err = errors.New(r.Error)
	}

	return
}
