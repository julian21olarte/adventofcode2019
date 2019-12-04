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
	
	numbers := strings.Split(line, ",")
	numbers[1] = "12"
	numbers[2] = "2"

	for i := 0; i < len(numbers); i += 4 {
		operation := numbers[i]
		if operation == "99" {
			break
		}

		num1, num2, num3 := numbers[i+1], numbers[i+2], numbers[i+3]
		pos1, _ := strconv.Atoi(num1)
		pos2, _ := strconv.Atoi(num2)
		targetPos, _ := strconv.Atoi(num3)
		number1, _ := strconv.Atoi(numbers[pos1])
		number2, _ := strconv.Atoi(numbers[pos2])

		if operation == "1" {
			numbers[targetPos] = strconv.Itoa(number1 + number2)
		}
		if operation == "2" {
			numbers[targetPos] = strconv.Itoa(number1 * number2)
		}
	}
	println(numbers[0])
}
