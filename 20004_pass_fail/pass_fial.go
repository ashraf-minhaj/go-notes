/* pass_fail shows if a grade is passing or failing

author: ashraf minhaj
mail  : ahraf_minhaj@yhaoo.com
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a grade: ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n') // get error value
	// input, _ := reader.ReadString('\n') // use blank identifier to ignor error

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(input)
	// fmt.Println(err)

	fmt.Println(reflect.TypeOf(input))

	cleanInput := strings.NewReplacer("\n", "").Replace(input) // get rid of the \n char
	// fixed := replacer.Replace(broken)
	grade, err := strconv.Atoi(cleanInput)
	fmt.Println(reflect.TypeOf(grade))

	if err != nil {
		log.Fatal(err)
	}

	if grade >= 60 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}

}
