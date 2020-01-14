package telegraph

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetViews(t *testing.T) {
	t.Run("invalid", func(t *testing.T) {
		t.Run("path", func(t *testing.T) {
			_, err := GetViews("wtf", time.Time{})
			assert.Error(t, err)
		})
		t.Run("year", func(t *testing.T) {
			dt := time.Date(1980, 0, 0, 0, 0, 0, 0, time.UTC)
			_, err := GetViews("Sample-Page-12-15", dt)
			assert.Error(t, err)
		})
		t.Run("month", func(t *testing.T) {
			dt := time.Date(2000, 22, 0, 0, 0, 0, 0, time.UTC)
			result, err := GetViews("Sample-Page-12-15", dt)
			assert.NoError(t, err)
			assert.NotNil(t, result)
		})
		t.Run("day", func(t *testing.T) {
			dt := time.Date(2000, time.February, 42, 0, 0, 0, 0, time.UTC)
			result, err := GetViews("Sample-Page-12-15", dt)
			assert.NoError(t, err)
			assert.NotNil(t, result)
		})
		t.Run("hour", func(t *testing.T) {
			dt := time.Date(2000, time.February, 12, 65, 0, 0, 0, time.UTC)
			result, err := GetViews("Sample-Page-12-15", dt)
			assert.NoError(t, err)
			assert.NotNil(t, result)
		})
	})
	t.Run("valid", func(t *testing.T) {
		dt := time.Date(2016, time.December, 31, 0, 0, 0, 0, time.UTC)
		stats, err := GetViews("Sample-Page-12-15", dt)
		assert.NoError(t, err)
		if !assert.NotNil(t, stats) {
			t.FailNow()
		}

		assert.NotZero(t, stats.Views)
	})
}
