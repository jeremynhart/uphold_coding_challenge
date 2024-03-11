package main

import (
	"testing" // Import the testing package for writing unit tests.
)

// TestPadNumbers is a test function for the padNumbers function.
func TestPadNumbers(t *testing.T) {
	// Define a slice of test cases, each with an input string, an integer X, and the expected output string.
	cases := []struct {
		input  string
		x      int
		output string
	}{
		{"James Bond 7", 3, "James Bond 007"}, // Test case 1: Expect "007" padding for the number 7.
		{"PI=3.14", 2, "PI=03.14"},            // Test case 2: Expect "03" padding for the number 3.
		{"It's 3:13pm", 2, "It's 03:13pm"},    // Test case 3: Expect "03" padding for the number 3.
		{"It's 12:13pm", 2, "It's 12:13pm"},   // Test case 4: No padding needed, numbers already meet the width requirement.
		{"99UR1337", 6, "000099UR001337"},     // Test case 5: Expect "000099" and "001337" padding for the numbers 99 and 1337.
	}

	for _, c := range cases { // Iterate over each test case.
		result := padNumbers(c.input, c.x) // Call padNumbers with the test case's input string and integer X.
		if result != c.output {            // If the result does not match the expected output...
			t.Errorf("padNumbers(%q, %d) == %q, want %q", c.input, c.x, result, c.output) // ...report an error.
		}
	}
}
