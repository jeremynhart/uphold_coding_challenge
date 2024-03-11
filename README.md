
README.md FILE ************************************************

03/08/2024: 
Friday March 8, 2024

GO CODING CHALLENGE: 

String Padding Program **************************************************************************

This GO program takes an input of any string along with an integer X and returns a string where any whole numbers found in the inputted string are left-padded with zeros to X characters.

How to Run the GO Program ***

1. Ensure Go is installed: Your system must have Go installed. If not, download and install it from the official Go website.
2. Set up your Go workspace: Clone this repository into your workspace or a preferred directory.
3. Navigate to the project directory: 
- cd path/to/project
4. Run the Program: go run main.go
- Follow the on-screen prompts to enter a string and an integer X.

Running Unit Tests ***

1. Run the following command in the project directory to execute the unit tests: go test


Time and Space Complexity Analysis *************************************************************

Time Complexity ***
The time complexity of this solution is primarily dictated by the regex matching and the iteration over the matches to apply padding. Regex matching over the string has a worst-case time complexity of O(n*m), where n is the length of the input string and m is the length of the pattern (in this case, the digits we're matching). However, for practical purposes and given the fixed pattern \d+, we can approximate the complexity of regex operations as O(n), where n is the length of the input string. For each match, we perform an operation that involves string conversion and formatting, which are both linear in the size of the output. However, since the size of each matched substring and the number of digits in the output string are relatively small, these operations can also be considered to have a constant time complexity. Therefore, the overall time complexity of the solution is approximately O(n), where n is the length of the input string.

Space Complexity ***
The space complexity of the solution is O(n), where n is the length of the output string. This takes into account the storage needed for the modified string with zero-padded numbers.
gex library's internal storage and the temporary variables used for processing each match contribute a smaller, additional amount of space, but this does not significantly affect the overall space complexity.

Jeremy Hart
Senior Engineer 
DevOps / CloudOps
720-421-0600
jeremy.hart@awde.us

