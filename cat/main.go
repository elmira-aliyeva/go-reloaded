package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

//PrintString lol
func PrintString(str string) {
	for _, j := range str {
		z01.PrintRune(j)
	}
	// z01.PrintRune('\n')
}

func main() {
	lenargs := 0
	for range os.Args {
		lenargs++
	}
	if lenargs > 1 {
		for i := 1; i < lenargs; i++ {
			file, err := ioutil.ReadFile(os.Args[i])
			if err != nil {
				// msg := string(err.Error())
				msg := "cat: " + os.Args[i] + ": No such file or directory"
				PrintString(msg)
				z01.PrintRune('\n')
				continue
			}
			PrintString(string(file))
		}
	}
	if lenargs == 1 {
		io.Copy(os.Stdin, os.Stdout)
	}
}



// Run the program without arguments and then write: Hello
// Hello
// Does the program returns the value above?

// Write: jaflsdfj
// jaflsdfj
// Does the program returns the value above?

// Write: Computer science (sometimes called computation science or computing science
// Computer science (sometimes called computation science or computing science
// Does the program returns the value above?

// Run the program passing the file: quest8.txt as an argument of execution of the program (as shown in the subject)
// ./cat quest8.txt
// Does the the program displays the same output as in the readme?
// Run the program with the file: quest8T.txt
// ./cat quest8T.txt
// Does the program displays the content of the file?
// Run the program with the arguments: quest8T.txt quest8.txt
// ./cat quest8.txt quest8T.txt 
// Does the program displays the content of the file in order?
// Run the program with a diferent file and then run the system program cat with the same file.
// Are the outputs identical?
// Run both this program and the system program cat passing as an argument a random string that is not the name of a file
// ./cat scddf
// cat: scddf: No such file or directory
// Are the outputs identical?