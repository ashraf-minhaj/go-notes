/* A function that takes an int and multiplies it by 2.
*
* author: ashraf minhaj
* mail  : ashraf_minhaj@yahoo.com
 */

package main

import (
	"fmt"
)

// func double(number int) int {
// 	return number * 2
// }

func persistDouble(number *int) {
	*number *= 2
}

func main() {
	amount := 6
	// doubleAmount := double(amount)

	// fmt.Println(doubleAmount)
	// fmt.Println(amount)

	// change by pointer
	// fmt.Println("\nNow trying with pointers")
	persistDouble(&amount)
	fmt.Println(amount)
}
