/* find types of values

author: ashraf minhaj
mail  : ahraf_minhaj@yhaoo.com
*/

package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf("Yo bro")) // string
	fmt.Println(reflect.TypeOf('a'))      // rune outputs unicode value, 97.
	fmt.Print(reflect.TypeOf(true))       // boolean
	fmt.Println(reflect.TypeOf(1))        // number, integer
	fmt.Println(reflect.TypeOf(420.69))   // number, floating value
}
