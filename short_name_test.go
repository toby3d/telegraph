package telegraph_test

import (
	"testing"

	"source.toby3d.me/toby3d/telegraph"
)

func TestNewShortName(t *testing.T) {
	t.Parallel()

	f := func(name, input string) {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual, err := telegraph.NewShortName(input)
			if err != nil {
				t.Fatal(err)
			}

			if actual.String() != input {
				t.Errorf("want '%s', got '%s'", input, actual)
			}
		})
	}

	f("short", "a")
	f("medium", telegraph.TestShortName(t).String())
	f("long", "HotDogGroundHogDayMetropolisMars")
}