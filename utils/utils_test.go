package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCapitalizeFirstLetter(t *testing.T) {
	r := CapitalizeFirstLetter("queue")

	assert.Equal(t, "Queue", r)

	r = CapitalizeFirstLetter("QUEUE")

	assert.Equal(t, "Queue", r)
}
