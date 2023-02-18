/* get file size using go

author: ashraf minhaj
mail  : ashraf_minhaj@yahoo.com
*/

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileInfo, err := os.Stat("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The size of the file is...", fileInfo.Size(), "bytes")
}
