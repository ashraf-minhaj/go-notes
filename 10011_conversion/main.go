/* conversion and comparison requrie same type.

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
	length := 3.23

	fmt.Println(float64(quantity) * length)

	fmt.Println(quantity > int(length))
}
