package main

import (
	"bufio"   // Import the bufio package for buffered I/O operations.
	"fmt"     // Import the fmt package for formatted I/O operations.
	"os"      // Import the os package for OS interface operations, including standard input.
	"regexp"  // Import the regexp package for regular expression operations.
	"strconv" // Import the strconv package for string conversion operations.
	"strings" // Import the strings package for string manipulation operations.
)

func main() {
	reader := bufio.NewReader(os.Stdin) // Create a new reader object to read from standard input.

	fmt.Println("Enter a string:")            // Prompt the user to enter a string.
	inputString, _ := reader.ReadString('\n') // Read the user's string input up to the newline character.

	fmt.Println("Enter an integer X:")              // Prompt the user to enter an integer X.
	xInput, _ := reader.ReadString('\n')            // Read the user's integer input as a string.
	x, _ := strconv.Atoi(strings.TrimSpace(xInput)) // Convert the integer string to an int type, removing any leading or trailing whitespace.

	outputString := padNumbers(inputString, x) // Call the padNumbers function with the user's string and integer inputs.
	fmt.Println("Output:", outputString)       // Print the modified string with padded numbers.
}

// padNumbers takes an input string and an integer X, then returns a modified string
// where whole numbers in the input string are left-padded with zeros to X characters.
func padNumbers(input string, x int) string {
	regex := regexp.MustCompile(`\d+`)                                   // Compile a regular expression to match one or more digits, representing whole numbers.
	return regex.ReplaceAllStringFunc(input, func(match string) string { // For each match of the regular expression in the input string...
		number, _ := strconv.Atoi(match)                  // Convert the matched string to an integer.
		formattedNumber := fmt.Sprintf("%0*d", x, number) // Format the number with left-padding zeros to ensure it has a width of X characters.
		return formattedNumber                            // Return the formatted (padded) number as the replacement string.
	})
}
