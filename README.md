# Go Util

A set of common function for golang data structure

## Installation
```shell
go get github.com/vanhtuan0409/goutil
```

## Documentation
Please refer to [godoc](https://godoc.org/github.com/vanhtuan0409/goutil)

Available function
- Slice
    - [func FillSliceByFloat(value float64, length int) []float64](https://godoc.org/github.com/vanhtuan0409/goutil/#FillSliceByFloat)
    - [func FillSliceByInt(value, length int) []int](https://godoc.org/github.com/vanhtuan0409/goutil/#FillSliceByInt)
    - [func IndexOf(arr interface{}, ele interface{}) int](https://godoc.org/github.com/vanhtuan0409/goutil/#IndexOf)
    - [func IsExist(arr interface{}, ele interface{}) bool](https://godoc.org/github.com/vanhtuan0409/goutil/#IsExist)
    - [func Range(start, end, step int) []int](https://godoc.org/github.com/vanhtuan0409/goutil/#Range)
    - [func RangeFloat(start, end, step float64) []float64](https://godoc.org/github.com/vanhtuan0409/goutil/#RangeFloat)
- Number
    - [func Abs(num int) int](https://godoc.org/github.com/vanhtuan0409/goutil/#Abs)
    - [func Max(num1, num2 int) int](https://godoc.org/github.com/vanhtuan0409/goutil/#Max)
    - [func Min(num1, num2 int) int](https://godoc.org/github.com/vanhtuan0409/goutil/#Min)
    - [func Pow(base, exponent int) int](https://godoc.org/github.com/vanhtuan0409/goutil/#Pow)
- JSON
    - [func ReadJSON(s string, result interface{}) error](https://godoc.org/github.com/vanhtuan0409/goutil/#ReadJSON)
    - [func ReadJSONFromFile(filename string, result interface{}) error](https://godoc.org/github.com/vanhtuan0409/goutil/#ReadJSONFromFile)
    - [func ReadJSONIO(r io.ReadCloser, result interface{}) error](https://godoc.org/github.com/vanhtuan0409/goutil/#ReadJSONIO)
    - [func WriteJSONToFile(filename string, data interface{}, indent int) error](https://godoc.org/github.com/vanhtuan0409/goutil/#WriteJSONToFile)
