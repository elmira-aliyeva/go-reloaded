package student

//StringLength lol
func StringLength(str string) int {
	StringLength := 0
	for range str {
		StringLength++
	}
	return StringLength
}

//Split lol
func Split(str, charset string) []string {
	str += charset
	r1 := []rune(str)
	l1 := StringLength(string(r1))
	r2 := []rune(charset)
	l2 := StringLength(string(r2))
	nbword := 0
	for i := 0; i <= l1-l2; i++ {
		if str[i:i+l2] == charset {
			nbword++
		}
	}
	arr := make([]string, nbword)
	counter := 0
	j := 0
	for i := counter; i <= l1-l2; i++ {
		if str[i:i+l2] == charset {
			arr[j] = str[counter:i]
			counter = i + l2
			j++
		}
	}
	return arr
}
