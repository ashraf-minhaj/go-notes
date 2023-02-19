/* Takes argument variable and runs a loop

 run: go run . arg

author: ashraf minhaj
mail  : ashraf_minhaj@yahoo.com
*/

package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
)

func main() {
	// unpack argument variable
	arg := os.Args[1]
	fmt.Println(arg)
	fmt.Println(reflect.TypeOf(arg))

	// cleanArg := strings.TrimSpace(arg)
	count, err := strconv.Atoi(arg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(count)

	// init and post are optional
	// count = 10
	for count <= 5 {
		fmt.Println(count)
		count++
	}
}
