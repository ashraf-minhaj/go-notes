/* remove whitespace from string

author: ashraf minhaj
mail  : ashraf_minhaj@yahoo.com
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "23\n"
	cleanData := strings.TrimSpace(data)
	fmt.Println(cleanData)
}
