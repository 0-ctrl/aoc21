package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/0-ctrl/aoc21/parsing"
)

func convertInput(input []string) (numbers [][]int) {
	for _, line := range input {
		num := make([]int, len(line))
		str := strings.Split(line, "")
		for n, s := range str {
			num[n], _ = strconv.Atoi(s)
		}
		numbers = append(numbers, num)
	}
	return
}

func findCommon(numbers [][]int, pos int) int {
	zeros, ones := 0, 0

	for _, line := range numbers {
		if line[pos] == 0 {
			zeros++
			continue
		}
		ones++
	}
	if zeros > ones {
		return 0
	}
	return 1
}

func destil(numbers [][]int, pos, target int) [][]int {
	var newNumbers [][]int
	for _, line := range numbers {
		if line[pos] == target {
			newNumbers = append(newNumbers, line)
		}
	}
	if len(newNumbers) == 0 {
		return numbers
	}
	return newNumbers
}

func readInput(filename string) ([]string, error) {
	var result []string

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func binSliceToDec(a []int) uint64 {
	if len(a) == 0 {
		return 0
	}

	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.Itoa(v)
	}
	r, _ := strconv.ParseUint(strings.Join(b, ""), 2, len(a))
	return r
}

func extractGamma(numbers [][]int) []string {
	lengths := len(numbers[0])

	result := make([]string, lengths)
	for i := 0; i < lengths; i++ {
		zeros := 0
		ones := 0
		for _, line := range numbers {
			if line[i] == 0 {
				zeros++
				continue
			}
			ones++
		}
		if zeros > ones {
			result[i] = "0"
			continue
		}
		result[i] = "1"
	}
	return result
}

func partOne(data []string, numbers [][]int) {
	extraction := extractGamma(numbers)

	lengths := len(numbers[0])
	mask, _ := strconv.ParseUint(strings.Repeat("1", lengths), 2, lengths)
	gamma, _ := strconv.ParseUint(strings.Join(extraction, ""), 2, lengths)
	epsilon := gamma ^ mask

	fmt.Println("Power consumtion: ", gamma*epsilon)
}
func partTwo(data []string, oxygen [][]int) {
	scrubber := make([][]int, len(oxygen))
	for i := range oxygen {
		scrubber[i] = make([]int, len(oxygen[i]))
		copy(scrubber[i], oxygen[i])
	}

	for i := 0; i < len(oxygen[0]); i++ {
		most := findCommon(oxygen, i)
		least := findCommon(scrubber, i) ^ 1
		oxygen = destil(oxygen, i, most)
		scrubber = destil(scrubber, i, least)
	}

	oxygenInt := binSliceToDec(oxygen[0])
	scrubberInt := binSliceToDec(scrubber[0])
	fmt.Println("Oxygen: \t", oxygenInt)
	fmt.Println("CO2 Scrubber: \t", scrubberInt)

	fmt.Println("Rating: \t", oxygenInt*scrubberInt)
}

func main() {
	input, err := parsing.ReadInputToSlice("input")
	if err != nil {
		log.Fatal(err)
	}

	diagnostics := convertInput(input)
	fmt.Println("Part one:")
	partOne(input, diagnostics)

	fmt.Println("------------------")

	fmt.Println("Part two:")
	partTwo(input, diagnostics)
}
