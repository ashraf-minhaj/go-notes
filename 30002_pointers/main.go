/* pointer example with go
* author: ashraf minhaj
* mail  : ashraf_minhaj@yahoo.com
 */

package main

import (
	"fmt"
	"reflect"
)

func main() {
	numberOfApples := 23
	fmt.Println(numberOfApples)

	// get pointer to numberOfApples
	fmt.Println(&numberOfApples)

	// get pointer type
	fmt.Println(reflect.TypeOf(&numberOfApples))

	// declare a variable to store pointer
	pointer := &numberOfApples
	fmt.Println(pointer)

	// print the values at a pointer
	fmt.Println(*pointer)

	// change value at pointer
	*pointer = 33

	// new value at pointer
	fmt.Println(*pointer)

	// **** change a value and pointer at the same time
	orange := 23
	fmt.Println(orange)

	pointerToOrange := &orange
	fmt.Println(pointerToOrange) // print pointer value
	fmt.Println(*pointer)        // get pointer value

	// update var
	*pointerToOrange = 44
	numberOfApples = 55
	fmt.Println(*pointerToOrange) // print pointer value
	fmt.Println(*pointer)         // get pointer value

	// change variable and pointers too
	fmt.Println("Change vars and pointers")
	myInt := 4
	fmt.Println("myint value:", myInt)
	fmt.Println("myint pointer:", &myInt)

	myPointer := &myInt
	*myPointer = 8
	fmt.Println("myint", myInt)
	fmt.Println("pointer value", *myPointer)
	fmt.Println("mypointer pointer:", myPointer)
}
