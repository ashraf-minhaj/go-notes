/* Exercise,
* We’ve written the negate function below,
* which is supposed to update the value of the truth variable to its opposite (false),
* and update the value of the lies variable to its opposite (true).
* But when we call negate on the truth and lies variables and then print their values,
* we see that they’re unchanged! Fix it.
*
* author: ashraf minhaj
* mail  : ashraf_minhaj@yahoo.com
 */

package main

import (
	"fmt"
)

func negate(myBoolean *bool) {
	// fmt.Println(myBoolean)
	// fmt.Println(*myBoolean)
	*myBoolean = !(*myBoolean)
}

func main() {
	truth := true
	negate(&truth)

	lies := false
	negate(&lies)

	fmt.Println(truth)
	fmt.Println(lies)
}
