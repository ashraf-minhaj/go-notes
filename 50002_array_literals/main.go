/* Array Literals
* If you know in advance what values an array should hold,
* you can initialize the array with those values using an array literal.
* author: ashraf minhaj
* mail  : ashraf_minhaj@yahoo.com
 */

package main

import "fmt"

func main() {
	// words := [4]string{} // decalre as empty array

	// words[0] = "Munajat"
	// words[1] = "dhor"
	// words[2] = "Shalar"
	// words[3] = "Putera!"

	// for i := 0; i < len(words); i++ {
	// 	fmt.Print(words[i], " ")
	// }

	// or -
	words := [4]string{
		"Munajat",
		"dhor",
		"shalar",
		"putera",
	}

	fmt.Println(words)
	fmt.Println(len(words))

	for i := 0; i < len(words); i++ {
		fmt.Print(words[i], " ")
	}

	fmt.Println("")

	words[3] = "vaiyera"

	for i := 0; i < len(words); i++ {
		fmt.Print(words[i], " ")
	}

	for i := 0; i < len(words); i++ {
		fmt.Print(words[i], " ")
	}
}
