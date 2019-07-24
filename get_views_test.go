package telegraph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetViews(t *testing.T) {
	t.Run("invalid", func(t *testing.T) {
		t.Run("path", func(t *testing.T) {
			_, err := GetViews("wtf")
			assert.Error(t, err)
		})
		t.Run("year", func(t *testing.T) {
			_, err := GetViews("Sample-Page-12-15", 1980)
			assert.Error(t, err)
		})
		t.Run("month", func(t *testing.T) {
			_, err := GetViews("Sample-Page-12-15", 2000, 22)
			assert.Error(t, err)
		})
		t.Run("day", func(t *testing.T) {
			_, err := GetViews("Sample-Page-12-15", 2000, 2, 42)
			assert.Error(t, err)
		})
		t.Run("hour", func(t *testing.T) {
			_, err := GetViews("Sample-Page-12-15", 2000, 2, 12, 65)
			assert.Error(t, err)
		})
	})
	t.Run("valid", func(t *testing.T) {
		stats, err := GetViews("Sample-Page-12-15")
		assert.NoError(t, err)
		if !assert.NotNil(t, stats) {
			t.FailNow()
		}
		assert.NotZero(t, stats.Views)
	})
}
