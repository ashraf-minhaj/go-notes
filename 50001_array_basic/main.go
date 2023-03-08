/* Array Basic
 * author: ashraf minhaj
 * mail  : ashraf_minhaj@yahoo.com
 */

package main

import "fmt"

func main() {
	var primes [5]int

	// or declare like this way -
	// primes := [5]int{}

	primes[0] = 1
	primes[3] = 3
	fmt.Println(primes[2])
	fmt.Println(primes[1])
	fmt.Println(primes[3])

	// increment value even without assigning values
	primes[1]++
	fmt.Println(primes[1])
}
