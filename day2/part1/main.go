package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
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
	scanner.Scan()
	line := scanner.Text()
	if err != nil {
		log.Fatal(err)
	}
	
	numbersLine := strings.Split(line, ",")
	numbers := []int{}
	for _, num := range numbersLine {
		numF, _ := strconv.Atoi(num)
		numbers = append(numbers, numF)
	}

	numbers[1] = 12
	numbers[2] = 2

	for i := 0; i < len(numbers); i += 4 {
		operation := numbers[i]
		if operation == 99 {
			break
		}

		pos1, pos2, targetPos := numbers[i+1], numbers[i+2], numbers[i+3]
		number1 := numbers[pos1]
		number2 := numbers[pos2]

		if operation == 1 {
			numbers[targetPos] = number1 + number2
		}
		if operation == 2 {
			numbers[targetPos] = number1 * number2
		}
	}
	println(numbers[0])
}
