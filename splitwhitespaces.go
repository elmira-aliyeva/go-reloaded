package student

//CountWords lol
func CountWords(str string, lens int) int {
	count := 0
	for i, l := range str {
		if (i != 0 && (l == ' ' || l == 10 || l == 9) && (str[i-1] != ' ' && str[i-1] != 10 && str[i-1] != 9)) || ((l != ' ' && l != 10 && l != 9) && i == lens-1) {
			count++
		}
	}
	return count
}

//SplitWhiteSpaces lol
func SplitWhiteSpaces(str string) []string {
	k, lens := 0, 0
	for range str {
		lens++
	}
	words := CountWords(str, lens)
	arr := make([]string, words)
	first, last, rep := -1, -1, 0
	for i, l := range str {
		if (l != ' ' && l != 10 && l != 9) && (i == 0 || (rune(str[i-1]) == ' ' || rune(str[i-1]) == 10 || rune(str[i-1]) == 9)) {
			first = i
		}
		if (((l == ' ' || l == 10 || l == 9) || i == lens-1) && i != 0 && first != -1 && (str[i-1] != ' ' && str[i-1] != 10 && str[i-1] != 9)) || (i == lens-1 && (l != ' ' && l != 10 && l != 9)) {
			last = i
		}
		if first != -1 && last != -1 && last != rep && (last > first || (last == first && last == lens-1)) {
			arr[k] = str[first:last]
			if (str[last] != ' ' && str[last] != 10 && str[last] != 9) || last == first {
				arr[k] += string(str[last])
			}
			k++
		}
		rep = last
	}
	return arr
}
