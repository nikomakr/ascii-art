package main

import (
	"fmt"
	"os"
	"ascii-art/art"
)

func main() {
	args := os.Args[1:] // Get command-line arguments, excluding the program name (os.Args[0] is the program name). This allows us to work with the user-provided input string and optional banner name directly.

	switch len(args) { // Check the number of arguments provided
	case 0: // No arguments provided. Either no input string or banner name provided. Notify user and exit with error code.
		fmt.Fprintln(os.Stderr, "Usage: go run . <string> [banner]")
		os.Exit(1)
	case 1:
		// Only input string provided — default banner is "standard". Render the input string using the default banner and print the result. If an error occurs during rendering, print the error message to standard error and exit with a non-zero status code.
		result, err := art.Render(args[0], "standard")
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}
		fmt.Print(result) // Print the rendered ASCII art to standard output.
	case 2: // Both input string and banner name provided. Render the input string using the specified banner and print the result. If an error occurs during rendering, print the error message to standard error and exit with a non-zero status code.
		// Input string + banner name provided
		result, err := art.Render(args[0], args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}
		fmt.Print(result) // Print the rendered ASCII art to standard output.
	default: // More than 2 arguments provided. Notify user about correct usage and exit with error code. This case handles the scenario where the user provides more than the expected number of arguments, which is considered incorrect usage of the program.
		fmt.Fprintln(os.Stderr, "Usage: go run . <string> [banner]")
		os.Exit(1)
	}
}