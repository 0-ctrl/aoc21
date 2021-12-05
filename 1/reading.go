package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/0-ctrl/aoc21/parsing"
)

func readInput(filename string) ([]int, error) {
	var result []int

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		reading, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("Unable to convert %v (%T)\n", scanner.Text(), scanner.Text())
		}
		result = append(result, reading)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func partOne(readings []int) int {
	increases := 0

	for i := 1; i < len(readings); i++ {
		if readings[i] > readings[i-1] {
			increases++
		}
	}
	return increases
}
func partTwo(readings []int) int {
	var slidingWindow []int
	readingLen := len(readings)

out:
	for i := 0; i < readingLen; i++ {
		sum := 0
		for j := i; j < i+3; j++ {
			if j >= readingLen {
				break out
			}
			sum += readings[j]
		}
		slidingWindow = append(slidingWindow, sum)
	}

	increases := 0
	for i := 1; i < len(slidingWindow); i++ {
		if slidingWindow[i] > slidingWindow[i-1] {
			increases++
		}
	}
	return increases
}

func main() {
	inputStr, err := parsing.ReadInputToSlice("input")
	if err != nil {
		log.Fatal(err)
	}
	readings, err := parsing.StrSliceToInt(inputStr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part one: Increases: ", partOne(readings))
	fmt.Println("Part two: Increases: ", partTwo(readings))
}
