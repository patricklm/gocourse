package basics

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// generate a random number between 1 and 100
	target := random.Intn(100) + 1

	// welcome message
	fmt.Println("Welcome to Guessing Game!")
	fmt.Println("I have chosen a number between 1 and 100")
	fmt.Println("Can you guess what it is?")

	var guess int
	for{
		fmt.Println("Enter your guess: ")
		fmt.Scanln(&guess)

		if guess == target {
			fmt.Println("Congratulations! You guessed the correct number!")
			break
		}else if guess < target {
			fmt.Println("Too low! Try guessing a higher number.")
		}else {
			fmt.Println("Too High! Try guessing a lower number.")
		}
	}
}
