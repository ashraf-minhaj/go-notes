/* Read data from a file and perform computation
*
* predict how much beef to order
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
	"strconv"
)

func getAverage(arr []float64) float64 {
	sum := 0.0
	length := len(arr)
	for _, value := range arr {
		// fmt.Println(value)
		sum += value
	}
	return sum / float64(length)
}

func main() {
	// an array to hold the values
	numbers := [3]float64{}
	indx := 0
	file, err := os.Open("data.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // close the file and free reources
	// fmt.Println(file)

	// loop until the end of the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		// fmt.Println(reflect.TypeOf(line))
		numbers[indx], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println(numbers, err)
		}
		indx++
	}
	// print the full array
	fmt.Println(numbers)

	// get average
	nextPossibleOrder := getAverage(numbers[:])
	fmt.Println(nextPossibleOrder)
}
