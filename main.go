package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const MAX_NUM_FAILED_GUESSES = 5

func main() {
	dictionary := []string{"hangman", "game over", "fun", "microchip", "breakfast"}
	word := selectWord(dictionary)
	word_progress := createProgressArray(word)

	failedAttemps := 0

	// Intial greeting
	fmt.Println(string(word_progress))
	fmt.Println()

	// Loop until ran out of chances
	for failedAttemps < MAX_NUM_FAILED_GUESSES {
		guess := guessCharacter()
		present, complete := checkGuessAgainstWord(guess, word, word_progress)

		// If word is complete return with success
		if complete {
			fmt.Println()
			fmt.Println("You win!! GG")
			return
		}
		// If user fails increase failures
		if !present {
			failedAttemps += 1
		}
		// Show state again after each character
		showState(word_progress, present, failedAttemps)
	}

	// Fail message after running out of chances
	fmt.Println()
	fmt.Println("You lose...")
}

// createProgressArray: initializes an array of '_'
func createProgressArray(word []byte) []byte {
	word_progress := make([]byte, len(word))
	for i := range word_progress {
		if word[i] != ' ' {
			word_progress[i] = '_'
		} else { // skip space
			word_progress[i] = ' '
		}
	}
	return word_progress
}

// selectWord: pick a word at random from dictionary
func selectWord(dictionary []string) []byte {
	rand.Seed(time.Now().UnixNano())
	pick := rand.Intn(len(dictionary))
	return []byte(dictionary[pick])
}

// showState: give feedback to the user about the state of the game
// !TODO: draw hangman ASCII...
func showState(word_progress []byte, guessCorrect bool, numAttemps int) {
	if !guessCorrect {
		msg := fmt.Sprintf("Noooo... nice try! You still have %d remaining\n", MAX_NUM_FAILED_GUESSES-numAttemps)
		fmt.Println(msg)
	} else {
		fmt.Println("Great guess!")
	}
	fmt.Println("")
	fmt.Println(string(word_progress))
}

// guessCharacter: get user input from prompt
func guessCharacter() byte {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Make a guess: ")
	text, _ := reader.ReadString('\n')

	// get only first character
	return []byte(text)[0]
}

// checkGuessAgainstWord: loop through progress to check if we are done
func checkGuessAgainstWord(guess byte, word []byte, progress []byte) (bool, bool) {
	present := false
	hitCount := 0
	for i, char := range word {
		if char == guess {
			progress[i] = char
			present = true
		}
		if char == progress[i] {
			hitCount += 1
		}
	}
	return present, hitCount == len(word)
}
