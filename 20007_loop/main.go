/* loop in golang

author: ashraf minhaj
mail  : ashraf_minhaj@yahoo.com
*/

package main

import (
	"fmt"
)

func main() {
	// initialization statement; condition; post statement
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// init and post are optional
	count := 10
	for count <= 15 {
		fmt.Println(count)
		count++
	}
}
