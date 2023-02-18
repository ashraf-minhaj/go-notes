/* store values in variables

author: ashraf minhaj
mail  : ahraf_minhaj@yhaoo.com
*/

package main

import (
	"fmt"
)

func main() {
	// declare the variables with type
	var quantity int
	var length, width float64
	var customerName string

	// if you know the value beforehand
	// you can omit the type of the value
	var size int = 6 // hehehe
	var height = 5.7

	// assign values
	quantity = 2
	// length, width = 14.5, 23.5 // this is also fine
	length = 1
	width = 23
	customerName = "Minhaj"

	fmt.Println(quantity)
	fmt.Println(length, width)
	fmt.Println(customerName)

	fmt.Println(size)
	fmt.Println((height))

}
