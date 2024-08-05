package util

import (
	"fmt"
	"unicode/utf8"
)

func ValidateLength(str string, min, max int) error {
	count := utf8.RuneCountInString(str)
	if min != -1 && count < min {
		return fmt.Errorf("want minimum length in %d characters, got %d", min, count)
	}

	if max < count {
		return fmt.Errorf("want maximum length in %d characters, got %d", max, count)
	}

	return nil
}