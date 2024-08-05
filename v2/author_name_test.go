package telegraph_test

import (
	"testing"

	"source.toby3d.me/toby3d/telegraph/v2"
)

func TestNewAuthorName(t *testing.T) {
	t.Parallel()

	f := func(name, input string) {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual, err := telegraph.NewAuthorName(input)
			if err != nil {
				t.Fatal(err)
			}

			if actual.String() != input {
				t.Errorf("want '%s', got '%s'", input, actual)
			}
		})
	}

	f("empty", "")
	f("short", "L")
	f("medium", telegraph.TestAuthorName(t).String())
	f("long", "Pablo Diego José Francisco de Paula Juan Nepomuceno María de los Remedios Cipriano de la Santísima "+
		"Trinidad Ruiz y Picasso")
}

func TestAuthorName_IsEmpty(t *testing.T) {
	t.Parallel()

	if actual, expect := (telegraph.AuthorName{}).IsEmpty(), true; actual != expect {
		t.Errorf("want %t, got %t", expect, actual)
	}
}