// take user input using go

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// create a bufio reader which will let us read keynoard input
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter something: ")
	userInput, err := reader.ReadString('\n') // take user input
	if err != nil {
		log.Fatal(err)
	}

	cleanUserInput := strings.TrimSpace(userInput) // clear whitespace
	fmt.Println(cleanUserInput)
}
