package goutil

func ExampleMax() {
	_ = Max(1, 5)
	// return 5

	_ = Max(0, -1)
	// return 0

	_ = Max(2, 2)
	// return 2
}
func ExampleMin() {
	_ = Min(1, 5)
	// return 1

	_ = Min(0, -1)
	// return -1

	_ = Min(2, 2)
	// return 2
}
func ExampleAbs() {
	_ = Abs(5)
	// return 5

	_ = Abs(-5)
	// return 5

	_ = Abs(0)
	// return 0
}
func ExamplePow() {
	_ = Pow(2, 0)
	// return 1

	_ = Pow(2, 3)
	// return 8

	_ = Pow(2, -5)
	// return 0
}
