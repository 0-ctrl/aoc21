package parsing

import (
	"bufio"
	"os"
	"strconv"
)

func ReadInput(filename string) (string, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(file), nil
}

func ReadInputToSlice(filename string) ([]string, error) {
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

func StrSliceToInt(str []string) ([]int, error) {
	var result []int
	for _, s := range str {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		result = append(result, num)
	}
	return result, nil
}
