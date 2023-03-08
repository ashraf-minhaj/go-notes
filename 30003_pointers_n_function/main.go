/* pointers and functions
*
* author: ashraf minhaj
* mail  : ashraf_minhaj@yahoo.com
 */

package main

import (
	"fmt"
	"reflect"
)

func createPointer() *float64 {
	// this function returns a float 64 pointer
	myFloat := 3.1416
	return &myFloat
}

func printPointer(myBoolPointer *bool) {
	// prints the pointer passed in it
	fmt.Println(*myBoolPointer)
}

func main() {
	myFloatPointer := createPointer()

	fmt.Println(reflect.TypeOf(myFloatPointer))
	fmt.Println(myFloatPointer)
	fmt.Println(*myFloatPointer)

	// pas pointer to a function
	myBool := true
	printPointer(&myBool)
}
