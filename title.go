package telegraph

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
	"unicode/utf8"

	"github.com/brianvoe/gofakeit/v7"
)

// Title represent page title.
type Title struct {
	title string // 1-256 characters
}

var ErrTitleLength error = errors.New("unsupported length of the provided string")

// NewTitle returns a new [Title] from string if validation successull, or
// [ErrTitleLength] error if Title is too short or long.
func NewTitle(raw string) (*Title, error) {
	if count := utf8.RuneCountInString(raw); count < 1 || 256 < count {
		return nil, fmt.Errorf("Title: %w: want 1-256 characters, got %d", ErrTitleLength, count)
	}

	return &Title{raw}, nil
}

// IsEmpty returns true if current [Title] is empty.
func (t Title) IsEmpty() bool { return t.title == "" }

func (t *Title) Update(newTitle string) error {
	if count := utf8.RuneCountInString(newTitle); count < 1 || 256 < count {
		return fmt.Errorf("Title: %w: want 1-256 characters, got %d", ErrTitleLength, count)
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

// TestTitle returns valid random generated [Title] for tests.
func TestTitle(tb testing.TB) *Title {
	tb.Helper()

	return &Title{gofakeit.SentenceSimple()}
}