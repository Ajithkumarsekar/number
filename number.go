package number

//IsEven takes int64 and returns true if it is an even number
func IsEven(num int) bool {
	return num%2 == 0
}

//IsOdd takes int64 and returns true if it is an odd number
func IsOdd(num int) bool {
	return !IsEven(num)
}

func IsPrime(num int) bool {
	if num < 2 {
		return false
	}

	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
