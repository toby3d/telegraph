package telegraph_test

import (
	"testing"

	"github.com/brianvoe/gofakeit/v7"

	"source.toby3d.me/toby3d/telegraph"
)

func TestNewTitle(t *testing.T) {
	t.Parallel()

	f := func(name, input string) {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual, err := telegraph.NewTitle(input)
			if err != nil {
				t.Fatal(err)
			}

			if actual.String() != input {
				t.Errorf("want '%s', got '%s'", input, actual)
			}
		})
	}

	f("short", "a")
	f("medium", telegraph.TestTitle(t).String())
	f("long", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla eget semper felis, vitae sagittis "+
		"ipsum. Fusce mollis et dui nec malesuada. Integer nisi est, cursus at velit in, aliquam tempor orci. "+
		"Aliquam imperdiet placerat sodales. Integer ac lorem amet.")
}

func TestTitle_Update(t *testing.T) {
	t.Parallel()

	title := telegraph.TestTitle(t)
	before := title.String()

	if err := title.Update(gofakeit.SentenceSimple()); err != nil {
		t.Fatal(err)
	}

	if after := title.String(); before == after {
		t.Errorf("want not equal '%s', got '%s'", before, after)
	}
}