package student

// ActiveBits returns number of active bits
func ActiveBits(n int) uint {
	var count uint = 0
	if n < 0 {
		n = n * -1
		if n == -9223372036854775808 {
			return 1
		}
	}
	for i := n; i > 0; i /= 2 {
		if i%2 == 1 {
			count++
		}
	}
	return count
}
