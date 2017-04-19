package goutil

func ExampleIsExist() {
	arr := []string{"a", "b", "c", "d"}
	_ = IsExist(arr, "b")
	// Return True

	_ = IsExist(arr, "f")
	// Return False
}

func ExampleIndexOf() {
	arr := []string{"a", "b", "c", "d"}
	_ = IndexOf(arr, "b")
	// Return 1

	_ = IndexOf(arr, "f")
	// Return -1
}

func ExampleRange() {
	_ = Range(0, 10, 1)
	// Return [0 ,1 ,2, 3, 4, 5, 6, 7, 8, 9, 10]

	_ = Range(0, 10, 2)
	// Return [0 ,2, 4, 6, 8, 10]
}

func ExampleRangeFloat() {
	_ = RangeFloat(0, 2, 0.5)
	// Return [0, 0.5, 1, 1.5, 2]
}
