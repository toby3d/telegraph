package telegraph

import (
	"fmt"
	"net/url"
	"strconv"
)

// URL represent (un)marshling domain for Telegraph JSON objects.
type URL struct {
	*url.URL `json:"-"`
}

func (u *URL) UnmarshalJSON(v []byte) error {
	unquoted, err := strconv.Unquote(string(v))
	if err != nil {
		return fmt.Errorf("URL: UnmarshalJSON: cannot unquote value '%s': %w", string(v), err)
	}

	result, err := url.ParseRequestURI(unquoted)
	if err != nil {
		return fmt.Errorf("URL: UnmarshalJSON: cannot parse value '%s': %w", string(v), err)
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