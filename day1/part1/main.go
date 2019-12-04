package main

import (
	"fmt"
	"log"
	"bufio"
	"os"
	"strconv"
	"math"
)

func main() {

	// open file input
	file, err := os.Open("./input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close() // close file on finish main function

		// generate the new scanner for read line by line the file
    scanner := bufio.NewScanner(file)

	result := 0.0
	// iterate in file opened
    for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text()) // read line and convert string to int
		
		if err != nil {
			log.Fatal(err)
		}
		result += process(float64(number))
	}
	println(fmt.Sprintf("%f", result))
}

func process(mass float64) float64 {
	num := math.Floor(mass / 3)
	num -= 2
	return num
}