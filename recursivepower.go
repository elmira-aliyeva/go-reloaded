package student

//RecursivePower returns power of int
func RecursivePower(nb int, power int) int {
	if power > 0 {
		return nb * RecursivePower(nb, power-1)
	} else if power == 0 {
		return 1
	} else {
		return 0
	}
}
