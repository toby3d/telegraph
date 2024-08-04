package telegraph

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// URL represent (un)marshling domain for Telegraph JSON objects.
type URL struct {
	*url.URL `json:"-"`
}

func NewURL(u *url.URL) *URL {
	return &URL{URL: u}
}

func (u *URL) UnmarshalJSON(v []byte) error {
	// WARN(toby3d): very strange string escaping only for auth_url, maybe
	// that can be removed later
	unquoted, err := strconv.Unquote(strings.ReplaceAll(string(v), `\/`, "/"))
	if err != nil {
		return fmt.Errorf("URL: UnmarshalJSON: cannot unquote value '%s': %w", string(v), err)
	}

	if len(unquoted) == 0 {
		return nil
	}

	result, err := url.ParseRequestURI(unquoted)
	if err != nil {
		return fmt.Errorf("URL: UnmarshalJSON: cannot parse value '%s': %w", unquoted, err)
	}

	u.URL = result

	return nil
}

func (u URL) MarshalJSON() ([]byte, error) {
	if u.URL != nil {
		return []byte(strconv.Quote(u.String())), nil
	}

	return nil, nil
}

func (u URL) GoString() string {
	if u.URL != nil {
		return "telegraph.URL(" + u.URL.String() + ")"
	}

	return "telegraph.URL(und)"
}