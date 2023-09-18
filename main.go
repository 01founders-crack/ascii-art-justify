package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) >= 3 {
		textFromOutside := os.Args[2]
		fileName := ""
		if len(os.Args) == 4 {
			if os.Args[3] == "shadow" || os.Args[3] == "standard" || os.Args[3] == "thinkertoy" {
				fileName = os.Args[3] // if something else standard
			}
		}

		if isASCII(os.Args[2]) {
			// Read Txt File
			fileLines := ReadStandardTxt(fileName)

			// 2d Ascii Art Ascii Array
			asciiTemplates := return2dASCIIArray(fileLines)

			// Ascii Art Array String (Printable)
			abc := returnAllStringArrASCII(textFromOutside, asciiTemplates)

			// max Len Of Ascii Art Array Horizontal
			maxLenOfAscii := returnMaxLenOfAscii(abc)

			//abc2 := returnSpaceSplit2dAsciiArr(textFromOutside, asciiTemplates)
			// lenOfAsciiArtHorizontal := lenOfAsciiArtWithoutSpace(abc2)
			lenOfAsciiArtHorizontalWSpc := lenOfAsciiArtWithSpace(abc)

			for _, v := range abc {
				if maxLenOfAscii < len(v) {
					maxLenOfAscii = len(v)
				}
			}
			lenOfTerminal := returnTerminalCols() - 2

			//lenOfWords := len(abc2)

			abc3 := returnWSpaceSplit2dAsciiArr(textFromOutside,asciiTemplates)
			lenOfWordsabc3 := len(abc3)

			alignType := os.Args[1]
			if checkArgs1(alignType) {
				if lenOfTerminal <= lenOfAsciiArtHorizontalWSpc {
					fmt.Println("Your terminal is not enough for the printing")
				} else {
					alignType = alignType[8:]
					if alignType == "right" {
						// right
						printAlignLeftOrRight(abc, strconv.Itoa(lenOfTerminal))
					} else if alignType == "left" {
						// left
						printAlignLeftOrRight(abc, strconv.Itoa(-lenOfTerminal))
					} else if alignType == "center" {
						// center
						printAlignCenter(abc, lenOfTerminal)
					} else if alignType == "justify" {
						printAlignJustify(abc3, lenOfWordsabc3, lenOfTerminal, lenOfAsciiArtHorizontalWSpc, len(textFromOutside))
					}
				} // if lenOfTerminal <= lenOfAsciiArtHorizontalWSpc
			} else {
				fmt.Println("args problem")
			} // if checkArgs1(alignType)

			fmt.Println()
			//fmt.Println(lenOfAsciiArtHorizontalWSpc, "as", lenOfAsciiArtHorizontal)
		} else {
			fmt.Println("Not Ascii ")
		} // if isASCII(os.Args[2])

	} // if len(os.Args) >= 3
}

func isASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}

func lenOfAsciiArtWithSpace(abc []string) int {
	temp2count := 0
	for _, v := range abc {
		if temp2count < len(v) {
			temp2count = len(v)
		}
	}
	return temp2count
}

func lenOfAsciiArtWithoutSpace(abc2 [][]string) int {
	temp1count := 0
	temp2count := 0
	for _, v := range abc2 {
		for _, y := range v {
			temp2count = 0
			if temp2count < len(y) {
				temp2count = len(y)
			}
			// fmt.Println(len(y))
		}
		// fmt.Println("sadadsad")
		temp1count = temp1count + temp2count
	}
	return temp1count
}

func returnWSpaceSplit2dAsciiArr(text string, asciiTemplates [][]string) [][]string {
	/*
		if ends w \n it gonna print println $
		if you can see text after \n chec;
		before \n
		if yes  println $
		if no println
	*/

	/*
	   func to uses printMultipleCharacters print whole stringfrom outside
	*/
	// Split the input string into an array of strings
	// split the line into words if there is a "\r\n" symbol
	var asciiStringArray [][]string
	// substrings := returnstring2EndlineArray(text)
	substrings := strings.Split(text, " ")
	lenOfsubstrings := len(substrings)
	for index, v := range substrings {
		var tempasciiStrArr []string
		if v == "\\n" {
			var tempStringAddStr string
			// If it is last one
			if index == lenOfsubstrings-1 {
				// tempStringAddStr = tempStringAddStr + fmt.Sprintln("")
			} else if index == 0 {
				// tempStringAddStr = tempStringAddStr + fmt.Sprintln("") // no idea CHECK IT POTENTIAL ERROR
			} else {
				if substrings[index-1] == "\\n" {
					// tempStringAddStr = tempStringAddStr + fmt.Sprintln("")
				} else {
					// "Hello\nWorld"
				}
			}
			tempasciiStrArr = append(tempasciiStrArr, tempStringAddStr)

		} else {
			tempasciiStrArr = append(tempasciiStrArr, printMultipleCharacter(v, asciiTemplates)...)
		}
		asciiStringArray = append(asciiStringArray, tempasciiStrArr)
	}

	return asciiStringArray
}


func returnSpaceSplit2dAsciiArr(text string, asciiTemplates [][]string) [][]string {
	/*
		if ends w \n it gonna print println $
		if you can see text after \n chec;
		before \n
		if yes  println $
		if no println
	*/

	/*
	   func to uses printMultipleCharacters print whole stringfrom outside
	*/
	// Split the input string into an array of strings
	// split the line into words if there is a "\r\n" symbol
	var asciiStringArray [][]string
	// substrings := returnstring2EndlineArray(text)
	substrings := strings.Split(text, "")
	lenOfsubstrings := len(substrings)
	for index, v := range substrings {
		var tempasciiStrArr []string
		if v == "\\n" {
			var tempStringAddStr string
			// If it is last one
			if index == lenOfsubstrings-1 {
				// tempStringAddStr = tempStringAddStr + fmt.Sprintln("")
			} else if index == 0 {
				// tempStringAddStr = tempStringAddStr + fmt.Sprintln("") // no idea CHECK IT POTENTIAL ERROR
			} else {
				if substrings[index-1] == "\\n" {
					// tempStringAddStr = tempStringAddStr + fmt.Sprintln("")
				} else {
					// "Hello\nWorld"
				}
			}
			tempasciiStrArr = append(tempasciiStrArr, tempStringAddStr)

		} else {
			tempasciiStrArr = append(tempasciiStrArr, printMultipleCharacter(v, asciiTemplates)...)
		}
		asciiStringArray = append(asciiStringArray, tempasciiStrArr)
	}

	return asciiStringArray
}

func returnTerminalCols() int {
	// Create a new command to run the "tput cols" command.
	cmd := exec.Command("tput", "cols")

	// Set the standard input of the command to os.Stdin (standard input).
	cmd.Stdin = os.Stdin

	// Execute the command and capture its output.
	out, err := cmd.Output()
	// Check if there was an error executing the command.
	if err != nil {
		// If there was an error, log the error message and return 0 as the terminal width.
		log.Fatal(err)
		return 0
	}

	// Convert the output (a byte slice) to a string and remove the trailing newline character.
	// Then, convert the string to an integer (terminal width).
	val, err := strconv.Atoi(string(out[:len(out)-1]))
	// Check if there was an error during the conversion.
	if err != nil {
		// If there was an error, return 0 as the terminal width.
		return 0
	}

	// Return the terminal width as an integer.
	return val
}

func printAlignJustify(abc2 [][]string, lenOfWords, lenOfTerminal, lenOfAsciiArtHorizontalWSpc, lentextFromOutside int) {
	// Iterate through the 8 lines of ASCII art characters
	for i := 0; i < 8; i++ {
		// Calculate the total number of spaces needed for justification.
		totalSpaces := lenOfTerminal - lenOfAsciiArtHorizontalWSpc

		// Calculate the number of gaps between words.
		numGaps := lenOfWords - 1

		// Declare minSpaces and extraSpaces outside the if-else block
		var minSpaces, extraSpaces int

		// Check if numGaps is zero to avoid divide by zero.
		if numGaps == 0 {
			// Handle the case where there are no gaps (e.g., only one word).
			minSpaces = totalSpaces
			extraSpaces = 0
		} else {
			// Calculate the minimum number of spaces between words.
			minSpaces = totalSpaces / numGaps

			// Calculate the number of extra spaces.
			extraSpaces = totalSpaces % numGaps
		}

		// Iterate through each word in the input text
		for j := 0; j < lenOfWords; j++ {
			// Add a "|" character at the beginning of the line (left boundary)
			if j == 0 {
				fmt.Print("|")
			}

			// Print the ASCII art character at position [j][i]
			fmt.Print(abc2[j][i])

			// Add the minimum number of spaces between words.
			if j != lenOfWords-1 {
				for k := 0; k < minSpaces; k++ {
					fmt.Print(" ")
				}
				fmt.Print("      ")
				// Add extra spaces if available.
				if extraSpaces > 0 {
					fmt.Print(" ")
					extraSpaces--
				}
			}
		}

		// Add a "|" character at the end of the line (right boundary)
		fmt.Println("|")
	}
}

func printAlignCenter(abc []string, terminalSize int) {
	for i := range abc {
		s := abc[i]
		padding := (terminalSize - len(s)) / 2
		fmt.Printf("|%*s%*s|\n", padding+len(s), s, terminalSize-padding-len(s), "")
	}
}

func printAlignLeftOrRight(abc []string, terminalSize string) {
	// if use  "-" terminalSize gonna be left otherwise it will be right side
	aaaa := "%" + terminalSize + "v"
	for i := range abc {
		fmt.Printf("|"+aaaa+"|\n", abc[i])
	}
}

func returnMaxLenOfAscii(abc []string) int {
	maxLenOfAscii := 0
	for _, v := range abc {
		if maxLenOfAscii < len(v) {
			maxLenOfAscii = len(v)
		}
	}
	return maxLenOfAscii
}

func checkArgs1(str string) bool {
	if len(str) > 7 {
		if str[:8] == "--align=" {
			return true
		}
	}
	return false
}

func ReadStandardTxt(nameOfFile string) []string {
	if nameOfFile == "" {
		nameOfFile = "standard"
	}
	readFile, err := os.Open("artstyles/" + nameOfFile + ".txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()
	return fileLines
}

func return2dASCIIArray(fileLines []string) [][]string {
	var asciiTemplates [][]string
	counter := 0
	var tempAsciArray []string
	for _, line := range fileLines {
		counter++
		// fmt.Println(index, line)
		if counter != 1 {
			tempAsciArray = append(tempAsciArray, line)
		}
		if counter == 9 {
			// fmt.Println("add to list") // but dont include the first line because it is empty line
			asciiTemplates = append(asciiTemplates, tempAsciArray)
			counter = 0
			tempAsciArray = nil
		} else {
		}
	}
	return asciiTemplates
}

// problem '\n' logic when we have 2 '\n' or 1 '\n' is different
func printMultipleCharacter(s string, asciiTemplates [][]string) []string {
	var mystring string
	var mStrArr []string
	// for ex 'hello'
	// we gonna write all letters index 0 after $ after \n after index 1  after $ after \n
	tempIntArrLetter := returnAsciiCodeInt(s)
	for i := 0; i < 8; i++ {
		mystring = ""
		for _, v := range tempIntArrLetter {
			mystring = mystring + fmt.Sprint(asciiTemplates[v][i])
		}
		// mystring = mystring + fmt.Sprintln("")

		mStrArr = append(mStrArr, mystring)
	}
	return mStrArr
}

func returnAsciiCodeInt(s string) []int {
	var tempIntArrLetter []int
	for _, v := range s {
		tempIntArrLetter = append(tempIntArrLetter, (int(v) - 32))
	}
	return tempIntArrLetter
}

func returnAllStringArrASCII(text string, asciiTemplates [][]string) []string {
	/*
		if ends w \n it gonna print println $
		if you can see text after \n chec;
		before \n
		if yes  println $
		if no println
	*/

	/*
	   func to uses printMultipleCharacters print whole stringfrom outside
	*/
	// Split the input string into an array of strings
	// split the line into words if there is a "\r\n" symbol
	var asciiStringArray []string
	substrings := returnstring2EndlineArray(text)
	lenOfsubstrings := len(substrings)
	for index, v := range substrings {
		if v == "\\n" {
			var tempStringAddStr string
			// If it is last one
			if index == lenOfsubstrings-1 {
				// tempStringAddStr = tempStringAddStr + fmt.Sprintln("")
			} else if index == 0 {
				// tempStringAddStr = tempStringAddStr + fmt.Sprintln("") // no idea CHECK IT POTENTIAL ERROR
			} else {
				if substrings[index-1] == "\\n" {
					// tempStringAddStr = tempStringAddStr + fmt.Sprintln("")
				} else {
					// "Hello\nWorld"
				}
			}
			asciiStringArray = append(asciiStringArray, tempStringAddStr)

		} else {
			asciiStringArray = append(asciiStringArray, printMultipleCharacter(v, asciiTemplates)...)
		}
	}

	return asciiStringArray
}

func returnstring2EndlineArray(text string) []string {
	substrings := make([]string, 0)
	escapedN := "\\n"
	newline := "\n"

	for {
		idx := strings.Index(text, escapedN)
		if idx == -1 {
			substrings = append(substrings, text)
			break
		}

		substrings = append(substrings, text[:idx])

		if idx+len(escapedN) < len(text) && text[idx+len(escapedN)] == 'n' {
			substrings = append(substrings, newline)
			text = text[idx+len(escapedN)+1:]
		} else {
			substrings = append(substrings, escapedN)
			text = text[idx+len(escapedN):]
		}
	}
	// fmt.Printf("%#v\n", substrings)
	var mysubstring2 []string
	for _, mysub := range substrings {
		if mysub != "" {
			mysubstring2 = append(mysubstring2, mysub)
		}
	}
	// fmt.Printf("%#v\n", mysubstring2)
	return mysubstring2
}
