/* Number Guessing Game using Golang

 This challenges users to guess e random number.

author: ashraf minhaj
mail  : ashraf_minhaj@yahoo.com
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func generateRandomNumber() int {
	seconds := time.Now().Unix()    // get current date and time an integer
	rand.Seed(seconds)              // seed the random number generator
	randomInt := rand.Intn(100 + 1) // generate randome number from 0-100
	// fmt.Println(targetRandomNumber)

	return randomInt
}

func main() {
	guesses := 10
	success := false

	// guess a random number
	targetRandomNumber := generateRandomNumber()
	fmt.Println(targetRandomNumber)

	fmt.Println("I've chosen a random number between 0 - 100, \ncan you guess it?")

	for guesses >= 1 {
		// take user input
		// create a bufio reader which will let us read keynoard input
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Make a guess: ")
		userInput, err := reader.ReadString('\n') // take user input
		if err != nil {
			log.Fatal(err)
		}

		cleanUserInput := strings.TrimSpace(userInput) // clear whitespace
		guess, err := strconv.Atoi(cleanUserInput)     // convert str to int

		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println(userInput)

		// compare guess to the target
		if guess == targetRandomNumber {
			fmt.Println("Congratulations, you've guessed right")
			success = true
			break
		} else if guess < targetRandomNumber {
			fmt.Println("Your guess was low")
		} else {
			fmt.Println("Your guess was high")
		}

		guesses--
		fmt.Println("Guesses left", guesses)
	}

	if !success {
		fmt.Println("You're out of move, the number was", targetRandomNumber)
	}

}
