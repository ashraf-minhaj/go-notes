/* go slice basic
*
* author: ashraf minhaj
* mail	: ashraf_minhaj@yahoo.com
 */

package main

import (
	"fmt"
	"log"
)

func main() {
	// mySlice := []string
	var notes []string        // declare the slice
	notes = make([]string, 7) // create a slice with 7 strings

	notes[0] = "mi"
	notes[1] = "nh"
	notes[2] = "aj"

	log.Println(notes)

	for _, value := range notes {
		fmt.Print(value)
	}

}
