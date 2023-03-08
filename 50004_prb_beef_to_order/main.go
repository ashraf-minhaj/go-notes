/* Take average order and figure out how
* much beef to order for a restaurant
* by averating.
*
* author: ashraf minhaj
* mail  : ashraf_minhaj@yahoo.com
 */

package main

import "fmt"

func main() {
	sum := 0.0

	// order histroy of last n weeks, we'll predict the 7th day
	// based on average order amount
	orderHistory := [3]float64{
		71.8,
		56.2,
		89.5,
	}

	for _, quantity := range orderHistory {
		fmt.Println(quantity)
		sum += quantity
	}
	fmt.Println(sum)

	average := sum / float64(len(orderHistory))

	fmt.Println(average)
}
