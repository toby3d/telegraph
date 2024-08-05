package telegraph

import (
	"fmt"
	"strconv"
	"testing"

	"source.toby3d.me/toby3d/telegraph/v2/internal/util"
)

// Title represent page title.
type Title struct {
	title string // 1-256 characters
}

// NewTitle returns a new [Title] from string if validation successull, or
// [ErrTitleLength] error if Title is too short or long.
func NewTitle(raw string) (*Title, error) {
	if err := util.ValidateLength(raw, 1, 256); err != nil {
		return nil, fmt.Errorf("Title: unsupported length: %w", err)
	}

	return &Title{raw}, nil
}

func (t *Title) Update(newTitle string) error {
	if err := util.ValidateLength(newTitle, 1, 256); err != nil {
		return fmt.Errorf("Title: unsupported length: %w", err)
	}

	t.title = newTitle

	return nil
}

func (t *Title) UnmarshalJSON(v []byte) error {
	unquoted, err := strconv.Unquote(string(v))
	if err != nil {
		return fmt.Errorf("Title: UnmarshalJSON: cannot unquote value '%s': %w", string(v), err)
	}

	result, err := NewTitle(unquoted)
	if err != nil {
		return fmt.Errorf("Title: UnmarshalJSON: cannot parse value '%s': %w", string(v), err)
	}

	*t = *result

	return nil
}

func (t Title) MarshalJSON() ([]byte, error) {
	if t.title != "" {
		return []byte(strconv.Quote(t.title)), nil
	}

	return nil, nil
}

func (t Title) String() string {
	return t.title
}

func (t Title) GoString() string {
	return "telegraph.Title(" + t.String() + ")"
}

// TestTitle returns valid [Title] for tests.
func TestTitle(tb testing.TB) *Title {
	tb.Helper()

	return &Title{"Sample Page"}
}