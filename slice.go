package goutil

import (
	"reflect"
)

// IsExist Check if element is existed in slice
func IsExist(arr interface{}, ele interface{}) bool {
	s := reflect.ValueOf(arr)
	if !isSlice(s) {
		return false
	}
	for i := 0; i < s.Len(); i++ {
		iter := s.Index(i)
		if reflect.DeepEqual(iter.Interface(), ele) {
			return true
		}
	}
	return false
}

// IndexOf Find index of element in slice
//
// If not exist, return -1
func IndexOf(arr interface{}, ele interface{}) int {
	s := reflect.ValueOf(arr)
	if !isSlice(s) {
		return -1
	}
	for i := 0; i < s.Len(); i++ {
		iter := s.Index(i)
		if reflect.DeepEqual(iter.Interface(), ele) {
			return i
		}
	}
	return -1
}

// Range Generate a slice of sequence number
//
// If start > end or step < 0, return empty slice
func Range(start, end, step int) []int {
	if step <= 0 || end < start {
		return []int{}
	}
	s := make([]int, 0, 1+(end-start)/step)
	for start <= end {
		s = append(s, start)
		start += step
	}
	return s
}

// RangeFloat Generate a slice of sequence number
//
// If start > end or step < 0, return empty slice
func RangeFloat(start, end, step float64) []float64 {
	if step <= 0 || end < start {
		return []float64{}
	}
	length := 1 + int((end-start)/step)
	s := make([]float64, 0, length)
	for start <= end {
		s = append(s, start)
		start += step
	}
	return s
}
