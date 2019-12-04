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
	
	valueToFind := 19690720
	numbersLine := strings.Split(line, ",")

	numbers := []int{}
	for _, num := range numbersLine {
		numF, _ := strconv.Atoi(num)
		numbers = append(numbers, numF)
	}

	numbersBackup := []int{}
	numbersBackup = append(numbersBackup, numbers...)

	numRange := []int{}
	for i := 0; i < 100; i++ {
		numRange = append(numRange, i)
	}

	result := 0
	for _, noun := range numRange {
		for _, verb := range numRange {
			numbers[1] = noun
			numbers[2] = verb
			result = getResult(numbers)
			if result == valueToFind {
				println("noun: ", noun, ", verb: ", verb)
				println("Result: ", 100 * noun + verb)
				break
			}
			numbers = append([]int{}, numbersBackup...)
		}
		if result == valueToFind {
			break
		}
	}
}

func getResult(numbers []int) int {
	lenNumbers := len(numbers)
	for i := 0; i < lenNumbers; i += 4 {
		operation := numbers[i]
		if operation == 99 {
			break
		}

		pos1, pos2, targetPos := numbers[i+1], numbers[i+2], numbers[i+3]
		number1 := numbers[pos1]
		number2 := numbers[pos2]

		if targetPos > lenNumbers {
			break
		}

		if operation == 1 {
			numbers[targetPos] = number1 + number2
		}
		if operation == 2 {
			numbers[targetPos] = number1 * number2
		}
	}
	return numbers[0]
}
