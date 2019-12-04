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
	
	valueToFindstr := "19690720"
	numbers := strings.Split(line, ",")
	var numbersBackup []string
	numbersBackup = append(numbersBackup, numbers...)

	numRange := []int{}
	for i := 0; i < 100; i++ {
		numRange = append(numRange, i)
	}

	result := ""
	for _, num1 := range numRange {
		for _, num2 := range numRange {
			noun := strconv.Itoa(num1)
			verb := strconv.Itoa(num2)
			numbers[1] = noun
			numbers[2] = verb
			result = getResult(numbers)
			if result == valueToFindstr {
				println("noun: ", noun, ", verb: ", verb)
				println("Result: ", 100 * num1 + num2)
				break
			}
			numbers = append([]string{}, numbersBackup...)
		}
		if result == valueToFindstr {
			break
		}
	}
}

func getResult(numbers []string) string {
	lenNumbers := len(numbers)
	for i := 0; i < lenNumbers; i += 4 {
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

		if targetPos > lenNumbers {
			break
		}

		if operation == "1" {
			numbers[targetPos] = strconv.Itoa(number1 + number2)
		}
		if operation == "2" {
			numbers[targetPos] = strconv.Itoa(number1 * number2)
		}
	}
	return numbers[0]
}
