// Package telegraph has functions and types used for interacting with the
// Telegraph API.
package telegraph

import (
	"errors"

	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
)

const (
	// APIEndpoint should be presented in this for all queries to the Telegraph
	// API must be served over HTTPS.
	APIEndpoint = "https://api.telegra.ph/%s"

	// PathEndpoint used if a path parameter is present.
	PathEndpoint = "https://api.telegra.ph/%s/%s"
)

func request(url string, args *http.Args) (*Response, error) {
	_, body, err := http.Post(nil, url, args)
	if err != nil {
		return nil, err
	}

	var resp Response
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	if !resp.Ok {
		return nil, errors.New(resp.Error)
	}

	return &resp, nil
}
