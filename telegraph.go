package telegraph

import (
	"encoding/json"
	"errors"

	jsoniter "github.com/json-iterator/go"
	http "github.com/valyala/fasthttp"
)

// Response contains a JSON object, which always has a Boolean field ok. If ok equals true, the request was
// successful, and the result of the query can be found in the result field. In case of an unsuccessful request, ok
// equals false, and the error is explained in the error field (e.g. SHORT_NAME_REQUIRED).
type Response struct {
	Ok     bool            `json:"ok"`
	Error  string          `json:"error,omitempty"`
	Result json.RawMessage `json:"result,omitempty"`
}

var parser = jsoniter.ConfigFastest //nolint:gochecknoglobals

func makeRequest(path string, payload interface{}) ([]byte, error) {
	src, err := parser.Marshal(payload)
	if err != nil {
		return nil, err
	}

	u := http.AcquireURI()
	defer http.ReleaseURI(u)
	u.SetScheme("https")
	u.SetHost("api.telegra.ph")
	u.SetPath(path)

	req := http.AcquireRequest()
	defer http.ReleaseRequest(req)
	req.SetRequestURIBytes(u.FullURI())
	req.Header.SetMethod(http.MethodPost)
	req.Header.SetUserAgent("toby3d/telegraph")
	req.Header.SetContentType("application/json")
	req.SetBody(src)

	resp := http.AcquireResponse()
	defer http.ReleaseResponse(resp)

	if err := http.Do(req, resp); err != nil {
		return nil, err
	}

	r := new(Response)
	if err := parser.Unmarshal(resp.Body(), r); err != nil {
		return nil, err
	}

	if !r.Ok {
		return nil, errors.New(r.Error)
	}

	return r.Result, nil
}
