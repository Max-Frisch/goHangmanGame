package main

import (
	"fmt"
)

var pl = fmt.Println

/*
 +---+
 0   |
/|\  |
/ \  |
    ===

 Secret Word: HA___A_
 Incorrect Guesses: U

 Guess a Letter: Y

 Sorry, you're dead! The word was HANGMAN
 Congrats! The secret word was HANGMAN

 Please Enter Only One Letter
 Please Enter Only Letters(a-z)
 Please Enter A Letter You Haven't Guessed Yet
*/

var hmArr = [7]string{
	" +---+\n" +
		"     |\n" +
		"     |\n" +
		"     |\n" +
		"    ===\n",
	" +---+\n" +
		" 0   |\n" +
		"     |\n" +
		"     |\n" +
		"    ===\n",
	" +---+\n" +
		" 0   |\n" +
		" |   |\n" +
		"     |\n" +
		"    ===\n",
	" +---+\n" +
		" 0   |\n" +
		"/|   |\n" +
		"     |\n" +
		"    ===\n",
	" +---+\n" +
		" 0   |\n" +
		"/|\\  |\n" +
		"     |\n" +
		"    ===\n",
	" +---+\n" +
		" 0   |\n" +
		"/|\\  |\n" +
		"/    |\n" +
		"    ===\n",
	" +---+\n" +
		" 0   |\n" +
		"/|\\  |\n" +
		"/ \\  |\n" +
		"    ===\n",
}

var wordArr = [14]string{
	"Xylophone", "Election", "Harmonica", "Zodiac", "Endoplasma", "Retina", "Ocean", "Pharmacology", "Randomness", "Elitism", "Riot", "Competition", "Blueberry", "Sausage",
}

var randWord string
var guessedLetters string
var correctLetters []string
var wrongLetters []string

func main() {

	// for i := 0; i < len(hmArr); i++ {
	// 	pl(hmArr[i])
	// }
}
