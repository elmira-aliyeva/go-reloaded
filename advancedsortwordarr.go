package student

//Compare lol
func Compare(b, a string) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

//AdvancedSortWordArr lol
func AdvancedSortWordArr(array []string, f func(a, b string) int) {
	l := len(array)
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if f(array[j], array[i]) > 0 {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
}
