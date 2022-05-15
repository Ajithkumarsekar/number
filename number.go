package number

//IsEven takes int64 and returns true if it is an even number
func IsEven(num int64) bool {
	return num%2 == 0
}

//IsOdd takes int64 and returns true if it is an odd number
func IsOdd(num int64) bool {
	return !IsEven(num)
}
