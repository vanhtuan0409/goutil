package goutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	assert.Equal(t, 5, Max(1, 5))
	assert.Equal(t, 0, Max(0, -1))
	assert.Equal(t, 2, Max(2, 2))
}
func TestMin(t *testing.T) {
	assert.Equal(t, 1, Min(1, 5))
	assert.Equal(t, -1, Min(0, -1))
	assert.Equal(t, 2, Min(2, 2))
}
func TestAbs(t *testing.T) {
	assert.Equal(t, 5, Abs(5))
	assert.Equal(t, 5, Abs(-5))
	assert.Equal(t, 0, Abs(0))
}
func TestPow(t *testing.T) {
	assert.Equal(t, 8, Pow(2, 3))
	assert.Equal(t, 2, Pow(2, 1))
	assert.Equal(t, 1, Pow(2, 0))
	assert.Equal(t, 0, Pow(2, -5))
}
