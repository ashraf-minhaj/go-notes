/* get time and date

author: ashraf minhaj
mail  : ahraf_minhaj@yhaoo.com
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	year := now.Year()

	fmt.Println(now)
	fmt.Println(year)

	// single liner
	fmt.Println(time.Now().Year())
	fmt.Println(time.Now().Second())
}
