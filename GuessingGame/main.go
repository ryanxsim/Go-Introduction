package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	println("Welcome to the Ultimate Guessing Game")
	println("=====================================")

	for {
		playGame()

		var playAgain string
		fmt.Println("Do you want to play again? (yes/no)")
		fmt.Scanln(&playAgain)

		if playAgain != "yes" {
			break
		}
	}
}

func playGame() {
	var userInputDifficulty string
	var userInputNumber string
	var numberRange int
	guessCount := 1

	println("Please choose the difficulty level")

	println("1. Easy")
	println("2. Medium")
	println("3. Hard")

	for {
		println("Enter your choice: ")

		fmt.Scanln(&userInputDifficulty)
		userInputDifficulty, err := strconv.Atoi(userInputDifficulty)

		if err != nil {
			fmt.Println("invalid input. please enter a number.")
			continue
		}

		switch userInputDifficulty {
		case 1:
			numberRange = 10
		case 2:
			numberRange = 100
		case 3:
			numberRange = 1000
		}
		break
	}

	var generatedNumber = generateNumber(numberRange)

	for {
		text := fmt.Sprintf("Guess a Number between 1 and %d", numberRange)
		fmt.Println(text)
		fmt.Scanln(&userInputNumber)

		guessedNumber, err := strconv.Atoi(userInputNumber)

		if err != nil {
			fmt.Println("invalid input. please enter a number.")
			continue
		}

		if guessedNumber == generatedNumber {
			text := fmt.Sprintf("You guessed it right. Total guesses : %d", guessCount)
			fmt.Println(text)
			break
		} else {
			fmt.Println("You guessed it wrong")
			fmt.Println("The number is " + higherOrLower(guessedNumber, generatedNumber) + " than " + userInputNumber)
			guessCount++
		}
	}
}

func generateNumber(numberRange int) int {
	return rand.Intn(numberRange) + 1
}

func higherOrLower(userNumber int, generatedNumber int) string {
	if userNumber > generatedNumber {
		return "lower"
	} else {
		return "higher"
	}
}
