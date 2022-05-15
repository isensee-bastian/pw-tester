package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const PW_LENGTH = 5

// Strings are defined as a slice of runes.
// Type rune is defined as an alias for int32 (holds 4 bytes).
type password []rune

func (p password) matches(guess []rune) bool {
	if len(p) != len(guess) {
		return false
	}

	for index := 0; index < len(p); index++ {
		if p[index] != guess[index] {
			return false
		}
	}

	return true
}

// Returns a slice of all usable symbols (runes).
func symbols() []rune {
	symbols := make([]rune, 0)

	// Lowercase letters a-z.
	for index := 97; index <= 122; index++ {
		symbols = append(symbols, rune(index))
	}

	// Uppercase letters A-Z.
	for index := 65; index <= 90; index++ {
		symbols = append(symbols, rune(index))
	}

	// Digits 0-9.
	for index := 48; index <= 57; index++ {
		symbols = append(symbols, rune(index))
	}

	return symbols
}

func guess(symbols []rune, pw password) {
	guess := make([]rune, PW_LENGTH)

	if !guessRecursive(symbols, pw, guess, 0) {
		fmt.Println("Not found")
	}
}

// Returns true if recursive logic shall abort (i.e. found match).
func guessRecursive(symbols []rune, pw password, guess []rune, index int) bool {
	// Check if guess has sufficient length for testing.
	if index >= PW_LENGTH {
		if pw.matches(guess) {
			// Found it, exit recursion.
			fmt.Printf("Found it: %c\n", guess)
			return true
		}
	} else {
		// Try next symbols.
		for _, symbol := range symbols {
			guess[index] = symbol

			if guessRecursive(symbols, pw, guess, index+1) {
				return true
			}
		}
	}

	// Could not find match yet. Continue with next symbol at previous position.
	return false
}

func main() {
	symbols := symbols()
	fmt.Printf("%c\n", symbols)

	file, err := os.Open("pw.txt")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		pw := password(line)

		before := time.Now()
		guess(symbols, pw)
		after := time.Now()

		fmt.Printf("  (took %v)\n", after.Sub(before))
	}

	err = scanner.Err()
	check(err)
}

func check(err error) {
	if err != nil {
		fmt.Println("Unexpected error: ", err)
		os.Exit(1)
	}
}
