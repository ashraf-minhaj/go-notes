/* For Range to access array elements
* and prevent panic.
*
* author: ashraf minhaj
* mail  : ashraf_minhaj@yahoo.com
 */

package main

import "fmt"

func main() {
	dua := [3]string{
		"Bismillahir",
		"Rahmanir",
		"Rahim",
	}

	for indx, value := range dua {
		fmt.Println(indx, value)
	}

	// print just the values
	for _, value := range dua {
		fmt.Print(value, " ")
	}
}
