/* reaplace substring in a string

author: ashraf minhaj
mail  : ahraf_minhaj@yhaoo.com
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	broken := "hey#minhaj#bro"
	replacer := strings.NewReplacer("#", " ")
	fixed := replacer.Replace(broken)

	fmt.Println(fixed)
}
