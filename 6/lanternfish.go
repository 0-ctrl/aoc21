package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/0-ctrl/aoc21/parsing"
)

type Lanternfish map[int]int

func (l Lanternfish) Sum() int {
	sum := 0
	for _, num := range l {
		sum += num
	}

	return sum
}

func (l Lanternfish) Days(days int) {
	for i := 0; i < days; i++ {
		l.day()
	}
}

// Moves fishes from n to n-1
// Lanternfish at 0 will spawn the same amount of new at 8
// Move those at 0 to 6
func (l Lanternfish) day() {
	spawning := l[0] // Conserv fishes at 0

	for i := 1; i <= 8; i++ {
		l[i-1] = l[i]
	}

	l[8] = spawning
	l[6] = l[6] + spawning
}

func main() {
	input, err := parsing.ReadInput("input")
	if err != nil {
		log.Fatal(err)
	}

	fishes := make(Lanternfish)
	for i := 0; i <= 8; i++ {
		fishes[i] = 0
	}

	for _, num := range strings.Split(input, ",") {
		fish, err := strconv.Atoi(num)
		if err != nil {
			log.Fatal(err)
		}

		fishes[fish]++
	}

	fishes.Days(256)

	fmt.Println("Result: ", fishes.Sum())
}
