package telegraph

import (
	gopath "path"
	"time"
)

type getViews struct {
	// Path to the Telegraph page (in the format Title-12-31, where 12 is the month and 31 the day the article was
	// first published).
	Path string `json:"path"`

	// Required if month is passed. If passed, the number of page views for the requested year will be returned.
	Year int `json:"year,omitempty"`

	// Required if day is passed. If passed, the number of page views for the requested month will be returned.
	Month int `json:"month,omitempty"`

	// Required if hour is passed. If passed, the number of page views for the requested day will be returned.
	Day int `json:"day,omitempty"`

	// If passed, the number of page views for the requested hour will be returned.
	Hour int `json:"hour,omitempty"`
}

// GetViews get the number of views for a Telegraph article. By default, the total number of page views will be
// returned. Returns a PageViews object on success.
func GetViews(path string, date time.Time) (*PageViews, error) {
	p := new(getViews)
	p.Path = path

	if !date.IsZero() {
		p.Year = date.Year()
		p.Month = int(date.Month())
		p.Day = date.Day()
		p.Hour = date.Hour()
	}

	data, err := makeRequest(gopath.Join("getViews", path), p)
	if err != nil {
		return nil, err
	}

	result := new(PageViews)
	if err = parser.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}
