package goutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsExist(t *testing.T) {
	arr := []string{"a", "b", "c", "d"}
	isExist := IsExist(arr, "b")
	assert.True(t, isExist)

	isExist = IsExist(arr, "f")
	assert.False(t, isExist)

	arr = []string{}
	isExist = IsExist(arr, "b")
	assert.False(t, isExist)

	arr2 := []int{0, 1, 2, 3, 4}
	isExist = IsExist(arr2, 4)
	assert.True(t, isExist)

	arr2 = []int{0, 1, 2, 3, 4}
	isExist = IsExist(arr2, 10)
	assert.False(t, isExist)
}

func TestIndexOf(t *testing.T) {
	arr := []string{"a", "b", "c", "d"}
	index := IndexOf(arr, "b")
	assert.Equal(t, 1, index)

	index = IndexOf(arr, "f")
	assert.Equal(t, -1, index)

	arr = []string{}
	index = IndexOf(arr, "b")
	assert.Equal(t, -1, index)

	arr2 := []int{0, 1, 2, 3, 4}
	index = IndexOf(arr2, 4)
	assert.Equal(t, 4, index)

	arr2 = []int{0, 1, 2, 3, 4}
	index = IndexOf(arr2, 10)
	assert.Equal(t, -1, index)
}

func TestRange(t *testing.T) {
	arr := Range(0, 10, 1)
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	assert.Equal(t, expected, arr)

	arr = Range(5, 10, 2)
	expected = []int{5, 7, 9}
	assert.Equal(t, expected, arr)

	arr = Range(10, 5, 2)
	expected = []int{}
	assert.Equal(t, expected, arr)

	arr = Range(5, 10, -2)
	expected = []int{}
	assert.Equal(t, expected, arr)
}

func TestRangeFloat(t *testing.T) {
	arr := RangeFloat(0.0, 2.0, 0.5)
	expected := []float64{0.0, 0.5, 1.0, 1.5, 2.0}
	assert.Equal(t, expected, arr)

	arr = RangeFloat(2.0, 0.0, 0.5)
	expected = []float64{}
	assert.Equal(t, expected, arr)

	arr = RangeFloat(0.0, 2.0, -0.5)
	expected = []float64{}
	assert.Equal(t, expected, arr)
}

func TestFillSliceByInt(t *testing.T) {
	assert.Equal(t, []int{}, FillSliceByInt(5, -1))
	assert.Equal(t, []int{}, FillSliceByInt(5, 0))
	assert.Equal(t, []int{5}, FillSliceByInt(5, 1))
	assert.Equal(t, []int{-5, -5, -5}, FillSliceByInt(-5, 3))
}

func TestFillSliceByFloat(t *testing.T) {
	assert.Equal(t, []float64{}, FillSliceByFloat(5.0, -1))
	assert.Equal(t, []float64{}, FillSliceByFloat(5.0, 0))
	assert.Equal(t, []float64{5.0}, FillSliceByFloat(5.0, 1))
	assert.Equal(t, []float64{-5.0, -5.0, -5.0}, FillSliceByFloat(-5.0, 3))
}
