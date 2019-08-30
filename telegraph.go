//go:generate ffjson $GOFILE
package telegraph

import (
	gojson "encoding/json"
	"errors"
	"github.com/valyala/fasthttp/fasthttpproxy"

	json "github.com/json-iterator/go"
	http "github.com/valyala/fasthttp"
)

// Response contains a JSON object, which always has a Boolean field ok. If ok
// equals true, the request was successful, and the result of the query can be
// found in the result field. In case of an unsuccessful request, ok equals
// false, and the error is explained in the error field (e.g.
// SHORT_NAME_REQUIRED).
type Response struct {
	Ok     bool              `json:"ok"`
	Error  string            `json:"error"`
	Result gojson.RawMessage `json:"result,omitempty"`
}

var (
	parser        = json.ConfigFastest //nolint:gochecknoglobals
	socks  string = ""
)

func makeRequest(path string, args *http.Args) ([]byte, error) {
	u := http.AcquireURI()
	defer http.ReleaseURI(u)
	u.SetScheme("https")
	u.SetHost("api.telegra.ph")
	u.SetPath(path)
	args.CopyTo(u.QueryArgs())

	req := http.AcquireRequest()
	defer http.ReleaseRequest(req)
	req.Header.SetMethod(http.MethodGet)
	req.SetRequestURIBytes(u.FullURI())
	req.Header.SetUserAgent("toby3d/telegraph")
	req.Header.SetContentType("application/json; charset=utf-8")

	resp := http.AcquireResponse()
	defer http.ReleaseResponse(resp)
	var client http.Client

	if (socks != "") {
		client = http.Client{
			Dial: fasthttpproxy.FasthttpSocksDialer(socks),
		}
	} else {
		client = http.Client{}
	}

	if err := client.Do(req, resp); err != nil {
		return nil, err
	}

	var r Response
	if err := parser.Unmarshal(resp.Body(), &r); err != nil {
		return nil, err
	}

	if !r.Ok {
		return nil, errors.New(r.Error)
	}

	return r.Result, nil
}
