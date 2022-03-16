package bit

// IsOdd tells whether a integer cannot be divided exactly by 2.
func IsOdd(num int64) bool {
	return (num & 1) == 1
}

// IsEven tells whether a integer can be divided exactly by 2.
func IsEven(num int64) bool {
	return (num & 1) == 0
}
