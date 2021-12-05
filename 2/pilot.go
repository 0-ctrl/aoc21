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

type Position struct {
	horizontal int
	depth      int
	aim        int
}

func (p *Position) forward(step int) {
	p.horizontal += step
	p.depth += p.aim * step
}

func (p *Position) up(step int) {
	p.aim -= step
}

func (p *Position) down(step int) {
	p.aim += step
}

func (p Position) sum() int {
	return p.depth * p.horizontal
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

func calculatePosition(data []string, submarine *Position) {
	for _, line := range data {
		s := strings.Split(line, " ")
		step, err := strconv.Atoi(s[1])
		if err != nil {
			log.Fatal(err)
		}
		switch s[0] {
		case "forward":
			submarine.forward(step)
		case "up":
			submarine.up(step)
		case "down":
			submarine.down(step)
		}
	}

}

func main() {

	data, err := parsing.ReadInputToSlice("input")
	if err != nil {
		log.Fatal(err)
	}

	submarine := Position{0, 0, 0}
	calculatePosition(data, &submarine)

	fmt.Println("Part one:")
	fmt.Printf("Pos/depth: %d/%d\n", submarine.horizontal, submarine.aim)
	fmt.Printf("Sum: %d\n", submarine.horizontal*submarine.aim)

	fmt.Println("----------------")

	fmt.Println("Part two:")
	fmt.Printf("Pos/depth: %d/%d\n", submarine.horizontal, submarine.depth)
	fmt.Printf("Sum: %d\n", submarine.sum())
}
