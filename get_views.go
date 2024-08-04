package telegraph

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
)

// GetViews get the number of views for a Telegraph article. Returns a
// [PageViews] object on success. By default, the total number of page views
// will be returned.
type GetViews struct {
	// Required. Path to the Telegraph page (in the format Title-12-31,
	// where 12 is the month and 31 the day the article was first
	// published).
	Path string `json:"-"`

	// Required if month is passed. If passed, the number of page views for
	// the requested year will be returned.
	Year uint16 `json:"year,omitempty"` // 2000-2100

	// Required if day is passed. If passed, the number of page views for
	// the requested month will be returned.
	Month uint8 `json:"month,omitempty"` // 1-12

	// Required if hour is passed. If passed, the number of page views for
	// the requested day will be returned.
	Day uint8 `json:"day,omitempty"` // 1-31

	// If passed, the number of page views for the requested hour will be
	// returned.
	Hour uint8 `json:"hour,omitempty"` // 0-24
}

func (params GetViews) Do(ctx context.Context, client *http.Client) (*PageViews, error) {
	data := make(url.Values)

	switch {
	case 0 < params.Year:
		if params.Year < 2000 {
			params.Year = 2000
		} else if 2100 < params.Year {
			params.Year = 2100
		}

		data.Set("year", strconv.FormatUint(uint64(params.Year), 10))

		fallthrough
	case 0 < params.Month:
		if 12 < params.Month {
			params.Month = 12
		}

		data.Set("month", strconv.FormatUint(uint64(params.Month), 10))

		fallthrough
	case 0 < params.Day:
		if 31 < params.Day {
			params.Day = 31
		}

		data.Set("day", strconv.FormatUint(uint64(params.Day), 10))

		fallthrough
	case 0 < params.Hour:
		if 24 < params.Hour {
			params.Hour = 24
		}

		data.Set("hour", strconv.FormatUint(uint64(params.Hour), 10))
	}

	return get[*PageViews](ctx, client, data, "getViews", params.Path)
}