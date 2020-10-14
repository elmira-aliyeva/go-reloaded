package main

import (
	student ".."
	"fmt"
	"strings"
)

//Compare lol
func Compare(a, b string) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

func main() {

	result := []string{"The", "earliest", "computing", "device", "undoubtedly", "consisted", "of", "the", "five", "fingers", "of", "each", "hand"}
	student.AdvancedSortWordArr(result, strings.Compare)
	fmt.Println(result)
	result = []string{"The", "word", "digital", "comesfrom", "\"digits\"", "or", "fingers"}
	student.AdvancedSortWordArr(result, strings.Compare)
	fmt.Println(result)
	result = []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	student.AdvancedSortWordArr(result, Compare)
	fmt.Println(result)
	result = []string{"The", "computing", "consisted", "device", "each", "earliest", "fingers", "five", "hand", "of", "of", "the", "undoubtedly"}
	student.AdvancedSortWordArr(result, student.Compare)
	fmt.Println(result) // b, a
	result = []string{"\"digits\"", "The", "comesfrom", "digital", "fingers", "or", "word"}
	student.AdvancedSortWordArr(result, student.Compare)
	fmt.Println(result) // b, a
	result = []string{"1", "2", "3", "A", "B", "C", "a", "b", "c"}
	student.AdvancedSortWordArr(result, student.Compare)
	fmt.Println(result) // b, a

}

// [The computing consisted device each earliest fingers five hand of of the undoubtedly]
// ["digits" The comesfrom digital fingers or word]
// [1 2 3 A B C a b c]
// [undoubtedly the of of hand five fingers earliest each device consisted computing The]
// [word or fingers digital comesfrom The "digits"]
// [c b a C B A 3 2 1]
