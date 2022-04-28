package utils_test

import (
	"testing"

	"github.com/nurcahyaari/kite/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestCapitalizeFirstLetter(t *testing.T) {
	r := utils.CapitalizeFirstLetter("queue")

	assert.Equal(t, "Queue", r)

	r = utils.CapitalizeFirstLetter("QUEUE")

	assert.Equal(t, "Queue", r)
}

func TestConcatDirPath(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		r := utils.ConcatDirPath("test", "test")
		assert.Equal(t, "test/test", r)
	})
	t.Run("test2", func(t *testing.T) {
		r := utils.ConcatDirPath("test/", "test")
		assert.Equal(t, "test/test", r)
	})
}
