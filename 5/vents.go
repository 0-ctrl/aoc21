package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/0-ctrl/aoc21/parsing"
)

const MatrixSize = 1024

type Coordinates [][]int

type Matrix struct {
	coordinates Coordinates
	diagonal    bool
}

func (m *Matrix) Add(x1, y1, x2, y2 int) {
	switch {
	case x1 == x2:
		min, max := sorted(y1, y2)
		for i := min; i < max+1; i++ {
			m.AddCoordinate(x1, i)
		}
	case y1 == y2:
		min, max := sorted(x1, x2)
		for i := min; i < max+1; i++ {
			m.AddCoordinate(i, y1)
		}
	default:
		if m.diagonal {
			m.AddCoordinate(x1, y1)
			m.AddCoordinate(x2, y2)

			if x1 < x2 {
				m.AddDiagonals(x1, x2, y1, y2)
			} else {
				m.AddDiagonals(x2, x1, y2, y1)
			}
		} else {
			fmt.Println("Not valid: ", x1, y1, x2, y2)
		}
	}
}

func (m *Matrix) AddDiagonals(i1, i2, j1, j2 int) {
	direction := -1
	if j2 > j1 {
		direction = 1
	}

	c := j1
	for i := i1 + 1; i < i2; i++ {
		c += direction
		m.AddCoordinate(i, c)
	}
}

func (m *Matrix) AddCoordinate(x, y int) {
	m.coordinates[x][y]++
}

func (m *Matrix) SumOverlapping() int {
	sum := 0

	for x := 0; x < len(m.coordinates); x++ {
		for y := 0; y < len(m.coordinates[x]); y++ {
			if m.coordinates[x][y] > 1 {
				sum++
			}
		}
	}

	return sum
}

func CreateMatrix(diagonal bool) *Matrix {
	coordinates := make([][]int, MatrixSize)

	for i := 0; i < MatrixSize; i++ {
		coordinates[i] = make([]int, MatrixSize)
	}

	return &Matrix{coordinates: coordinates, diagonal: diagonal}
}

func sorted(x, y int) (int, int) {
	s := []int{x, y}
	sort.Ints(s)

	return s[0], s[1]
}

func toInt(s string) int {
	r, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return r
}

func main() {
	input, err := parsing.ReadInput("input")
	if err != nil {
		log.Fatal(err)
	}

	diagonal := true
	matrix := CreateMatrix(diagonal)

	for _, line := range strings.Split(input, "\n") {
		s := strings.Split(line, " -> ")

		first := strings.Split(s[0], ",")
		y1, x1 := toInt(first[0]), toInt(first[1])

		second := strings.Split(s[1], ",")
		y2, x2 := toInt(second[0]), toInt(second[1])

		matrix.Add(x1, y1, x2, y2)
	}

	result := matrix.SumOverlapping()
	fmt.Println("Result is: ", result)
}
