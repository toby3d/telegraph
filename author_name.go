package telegraph

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
	"unicode/utf8"

	"github.com/brianvoe/gofakeit/v7"
)

// AuthorName represent author name used when creating new articles.
type AuthorName struct {
	authorName string // 0-128 characters
}

var ErrAuthorNameLength error = errors.New("unsupported length")

// NewAuthorName parse raw string as AuthorName and validate it's length.
func NewAuthorName(raw string) (*AuthorName, error) {
	if count := utf8.RuneCountInString(raw); 128 < count {
		return nil, fmt.Errorf("AuthorName: %w: want up to 128 characters, got %d", ErrAuthorNameLength, count)
	}

	return &AuthorName{raw}, nil
}

func (an *AuthorName) UnmarshalJSON(v []byte) error {
	unquoted, err := strconv.Unquote(string(v))
	if err != nil {
		return fmt.Errorf("AuthorName: UnmarshalJSON: cannot unquote value '%s': %w", string(v), err)
	}

	result, err := NewAuthorName(unquoted)
	if err != nil {
		return fmt.Errorf("AuthorName: UnmarshalJSON: cannot parse value '%s': %w", string(v), err)
	}

	*an = *result

	return nil
}

func (an AuthorName) MarshalJSON() ([]byte, error) {
	if an.authorName != "" {
		return []byte(strconv.Quote(an.authorName)), nil
	}

	return nil, nil
}

func (an AuthorName) String() string {
	return an.authorName
}

func (an AuthorName) GoString() string {
	return "telegraph.AuthorName(" + an.String() + ")"
}

func TestAuthorName(tb testing.TB) *AuthorName {
	tb.Helper()

	return &AuthorName{gofakeit.FirstName() + " " + gofakeit.LastName()}
}