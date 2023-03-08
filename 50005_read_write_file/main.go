/* Read from a file
*
* author: ashraf minhaj
* mail  : ashraf_minhaj@yahoo.com
 */

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
)

func main() {
	file, err := os.Open("data.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(file)

	// loop until the end of the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		fmt.Println(reflect.TypeOf(line))
	}

	// close the file and free reources
	defer file.Close()

	if scanner.Err() != nil {
		log.Fatal(err)
	}

	// finally write some data in the file
	file, err = os.OpenFile("data.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // opening...
	if err != nil {
		fmt.Printf("error creating file: %v", err)
		return
	}
	_, err = file.WriteString("\n24")
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(file.Read())
}
