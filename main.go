package main

import (
    "fmt"    // Importing fmt for printing
    "os"     // Importing os for file and command line operations
    "strings" // Importing strings for string manipulation
)

// asciiMap is a global variable that stores the ASCII art templates for characters.
// The keys are of type rune (Unicode code points) and the values are arrays of 8 strings.
var asciiMap map[rune][8]string

// getUserInput reads and validates user input from command line arguments.
func getUserInput() string {
    inputArgs := os.Args[1:] // Get the command line arguments, excluding the program name
    if len(inputArgs) != 1 { // Check if exactly one argument is provided
        fmt.Println("Please provide an input") // Print an error message if no input is provided
        os.Exit(1) // Exit the program with an error code
    }
    if inputArgs[0] == "" { // Check if the input argument is an empty string
        os.Exit(0) // Exit the program without an error code
    }
    if inputArgs[0] == "\\n" { // Check if the input argument is a newline character
        return "" // Return an empty string
    }
    return inputArgs[0] // Return the input argument
}

// removeEmptyLines removes empty lines from a slice of strings.
func removeEmptyLines(lines []string) (result []string) {
    for _, line := range lines { // Loop through each line in the input slice
        if line == "" { // Check if the line is empty
            continue // Skip empty lines
        }
        result = append(result, line) // Add non-empty lines to the result slice
    }
    return // Return the result slice with empty lines removed
}

// createAsciiMap takes a slice of strings and builds a map where each rune corresponds
// to an array of 8 strings from the input slice.
func createAsciiMap(inputLines []string) map[rune][8]string {
    asciiMapResult := make(map[rune][8]string) // Initialize an empty map
    for i := 0; i < len(inputLines); i += 8 { // Loop through the input slice in steps of 8 lines
        // Map each set of 8 lines to a character, starting from the space character (' ')
        asciiMapResult[rune(i/8+' ')] = [8]string{
            inputLines[i], inputLines[i+1], inputLines[i+2], inputLines[i+3],
            inputLines[i+4], inputLines[i+5], inputLines[i+6], inputLines[i+7],
        }
    }
    return asciiMapResult // Return the resulting map
}

// printAsciiArt prints a single word in ASCII art format by concatenating the lines from the global asciiMap.
func printAsciiArt(word string) {
    asciiOutput := "" // Initialize an empty string to store the ASCII art
    for i := 0; i < 8; i++ { // Loop through each of the 8 lines of the ASCII art
        for _, char := range word { // Loop through each character in the word
            asciiOutput += asciiMap[char][i] // Add the corresponding line of the character's ASCII art to the output
        }
        asciiOutput += "\n" // Add a newline character after each line of the ASCII art
    }
    fmt.Print(asciiOutput) // Print the final ASCII art for the word
}

// printWords takes a list of words and prints each word as ASCII art.
func printWords(wordList []string) {
    for _, word := range wordList { // Loop through each word in the list
        if word == "" { // Check if the word is empty
            fmt.Println() // Print a blank line if the word is empty
        } else {
            printAsciiArt(word) // Print the ASCII art for each word
        }
    }
}

// main is the entry point of the program. It reads user input, reads a file, builds an ASCII map,
// splits the user input into a list of words, and prints the words as ASCII art.
func main() {
    userInput := getUserInput() // Get the user input from the command line
    fileContent, _ := os.ReadFile("standard.txt") // Read the ASCII art template file
    // Replace Windows-style line endings with Unix-style line endings
    unixFileContent := strings.ReplaceAll(string(fileContent), "\r\n", "\n")
    // Split the file content into lines
    lines := strings.Split(unixFileContent, "\n")
    // Remove empty lines from the file content
    cleanedLines := removeEmptyLines(lines)
    // Build the ASCII map from the cleaned lines
    asciiMap = createAsciiMap(cleanedLines)
    // Split the user input into words
    words := strings.Split(userInput, "\\n")
    // Print each word as ASCII art
    printWords(words)
}