package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/0-ctrl/aoc21/parsing"
)

type Board struct {
	matrix [][]*Content
	hasWon bool
}

type Content struct {
	num    int
	marked bool
}

func (b *Board) mark(num int) {
	for _, row := range b.matrix {
		for _, col := range row {
			if col.num == num {
				col.marked = true
			}
		}
	}
}

func (b *Board) sum() int {
	// Sums numbers not marked
	sum := 0
	for _, row := range b.matrix {
		for _, col := range row {
			if !col.marked {
				sum += col.num
			}
		}
	}
	return sum
}

func (b *Board) print() {
	for _, row := range b.matrix {
		result := ""
		for _, col := range row {
			if col.marked {
				result += fmt.Sprintf("[%2d]", col.num)
			} else {
				result += fmt.Sprintf(" %2d ", col.num)
			}
		}
		fmt.Println(result)
	}
	fmt.Println()
}

func (b *Board) checkHorizontal() bool {
	for _, row := range b.matrix {
		win := true
		for _, col := range row {
			if !col.marked {
				win = false
				break
			}
		}
		if win {
			return true
		}
	}
	return false
}

func (b *Board) checkVertical() bool {

	length := len(b.matrix[0])
	for i := 0; i < length; i++ {
		win := true

		for j := 0; j < length; j++ {
			if !b.matrix[j][i].marked {
				win = false
				break
			}
		}
		if win {
			return true
		}
	}
	return false
}

func (b *Board) checkWin() bool {
	if win := b.checkHorizontal(); win {
		return true
	}
	return b.checkVertical()
}

func createBoard(nums []string) Board {
	var board Board

	for _, line := range nums {
		var contents []*Content

		for _, n := range strings.Fields(line) {
			num, err := strconv.Atoi(n)
			if err != nil {
				log.Fatal(err)
			}

			contents = append(contents, &Content{num: num, marked: false})
		}

		board.matrix = append(board.matrix, contents)
	}
	return board
}

func parseBoard(filename string) ([]int, []*Board) {
	input, err := parsing.ReadInput(filename)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(input, "\n")

	// Handle draw numbers
	var draw []int
	for _, drawStr := range strings.Split(lines[0], ",") {
		num, err := strconv.Atoi(drawStr)
		if err != nil {
			log.Fatal(err)
		}
		draw = append(draw, num)
	}

	// Handle the boards
	inputBoards := lines[2:]
	numBoards := len(inputBoards)

	var boards []*Board
	for i := 0; i < numBoards; i += 5 {
		board := createBoard(inputBoards[0:5])
		boards = append(boards, &board)

		if len(inputBoards) > 6 {
			inputBoards = inputBoards[6:]
		}
	}

	return draw, boards
}

func partOne(draws []int, boards []*Board) {
	for i := 0; i < len(draws); i++ {
		draw := draws[i]
		for _, board := range boards {
			board.mark(draw)
			if board.checkWin() {
				sum := board.sum()
				result := sum * draw
				fmt.Println("Winner score: ", result)
				return
			}
		}
	}
}

func partTwo(draws []int, boards []*Board) {
	lastBoardToWin := 0

	for i := 0; i < len(draws); i++ {
		draw := draws[i]

		for _, board := range boards {
			// Skip boards that already won
			if board.hasWon {
				continue
			}

			board.mark(draw)
			if board.checkWin() {
				sum := board.sum()
				board.hasWon = true
				lastBoardToWin = sum * draw
			}
		}
	}

	fmt.Println("Last board to win, score: ", lastBoardToWin)
}

func main() {
	draws, boards := parseBoard("input")

	fmt.Println("Part One:")
	partOne(draws, boards)

	fmt.Println("-----------------------------")

	fmt.Println("Part Two:")
	partTwo(draws, boards)
}
