/* short variable, declare and if you know it

author: ashraf minhaj
mail  : ahraf_minhaj@yhaoo.com
*/

package main

import (
	"fmt"
)

func main() {
	// declare the variables of which you know the value instantly
	quantity := 2
	length, width := 3.1, 4.2
	customerName := "Minhaj"

	var success bool

	// now I know the value
	success = true

	fmt.Println(quantity)
	fmt.Println(length, width)
	fmt.Println(customerName)
	fmt.Println(success)

	// no new variables error
	// customerName := "ashraf"
	// fmt.Print(customerName)
}
