package dto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptString(t *testing.T) {
	notSet := OptString{}
	assert.Equal(t, "", notSet.GetValue())
	assert.False(t, notSet.IsSet())

	set := NewOptString("value")
	assert.Equal(t, "value", set.GetValue())
	assert.True(t, set.IsSet())
}
