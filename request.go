package telegraph

import (
	gojson "encoding/json"
	"errors"
	"fmt"
	"net/url"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

// response represents a response from the Telegram API with the result stored raw. If ok equals true,
// the request was successful, and the result of the query can be found in the result field. In case of
// an unsuccessful request, ok equals false, and the error is explained in the error field (e.g.
// SHORT_NAME_REQUIRED).
type response struct {
	Ok     bool               `json:"ok"`
	Error  string             `json:"error"`
	Result *gojson.RawMessage `json:"result"`
}

func request(method, path string, args *http.Args) (*response, error) {
	requestURI := &url.URL{
		Scheme: "https",
		Host:   "api.telegra.ph",
		Path:   method,
	}

	if path != "" {
		requestURI.Path += fmt.Sprint("/", path)
	}

	_, body, err := http.Post(nil, requestURI.String(), args)
	if err != nil {
		return nil, err
	}

	var resp response
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	if !resp.Ok {
		return nil, errors.New(resp.Error)
	}

	return &resp, nil
}
