/* different things with a function
* author: ashraf minhaj
* mail  : ashraf_minhaj@yahoo.com
 */

package main

import (
	"fmt"
	"log"
)

func multiply(num1 int, num2 int) (int, error) {
	/* handle error, and return using a function */
	if num1 < 0 || num2 < 0 {
		return 0, fmt.Errorf("the value can not be negative, reason My boss does not like them")
	}
	result := num1 * num2
	return result, nil
}

func main() {
	// call the function
	answer, err := multiply(-2, 3)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(answer)
}
