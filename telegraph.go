package telegraph // import "source.toby3d.me/toby3d/telegraph"

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
)

// response is a JSON object, which always has a Boolean field ok. If ok equals
// true, the request was successful, and the result of the query can be found in
// the result field. In case of an unsuccessful request, ok equals false, and
// the error is explained in the error field (e.g. SHORT_NAME_REQUIRED).
type response[T any] struct {
	Result T      `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
	OK     bool   `json:"ok"`
}

var DefaultEndpoint *url.URL = &url.URL{
	Scheme: "https",
	Host:   "api.telegra.ph",
	Path:   "/",
}

func get[T any](ctx context.Context, client *http.Client, data url.Values, method string, pagePath ...string) (T, error) {
	u, _ := url.ParseRequestURI(DefaultEndpoint.String())
	u.Path = path.Join("/", method)
	u.RawQuery = data.Encode()

	if 0 < len(pagePath) {
		u.Path = path.Join("/", method, pagePath[0])
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		var result T
		return result, fmt.Errorf("%s: cannot initialize request: %w", method, err)
	}

	return do[T](client, req)
}

func post[T any](ctx context.Context, client *http.Client, r io.Reader, method string, pagePath ...string) (T, error) {
	u, _ := url.ParseRequestURI(DefaultEndpoint.String())
	u.Path = path.Join("/", method)

	if 0 < len(pagePath) {
		u.Path = path.Join("/", method, pagePath[0])
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u.String(), r)
	if err != nil {
		var result T
		return result, fmt.Errorf("%s: cannot initialize request: %w", method, err)
	}

	req.Header.Set("Content-Type", "application/json")

	return do[T](client, req)
}

func do[T any](client *http.Client, req *http.Request) (T, error) {
	result := new(response[T])

	resp, err := client.Do(req)
	if err != nil {
		return result.Result, fmt.Errorf("cannot make request: %w", err)
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(result); err != nil {
		return result.Result, fmt.Errorf("cannot decode response: %w", err)
	}

	if result.OK {
		return result.Result, nil
	}

	return result.Result, fmt.Errorf("error response: %s", result.Error)
}