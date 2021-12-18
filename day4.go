package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var checks = [][]int{
	{0, 1, 2, 3, 4},
	{5, 6, 7, 8, 9},
	{10, 11, 12, 13, 14},
	{15, 16, 17, 18, 19},
	{20, 21, 22, 23, 24},
	{0, 5, 10, 15, 20},
	{1, 6, 11, 16, 21},
	{2, 7, 12, 17, 22},
	{3, 8, 13, 18, 23},
	{4, 9, 14, 19, 24},
	// {0, 6, 12, 18, 24},
	// {4, 8, 12, 16, 20},
}

func Day4Part1() {
	called, boards := readInput("04-input.txt")
	log.Printf("calledNumbers: %v, num boards: %d", called, len(boards))
	for callCount, number := range called {
		log.Printf("Calling %d", number)
		for _, board := range boards {
			for i, spot := range board {
				if spot == number {
					board[i] = -1
					break
				}
			}
			if callCount < 5 {
				continue
			}
			// see if there's a winning line
			for _, line := range checks {
				// log.Printf("checking board id=%d", boardId)
				win := checkForWin(board, line)
				if win {
					sum := sumBoard(board)
					log.Printf("*** WINNER (via %v) *** sum=%d, last number called=%d, answer=%d", line, sum, number, sum*number)
					return
				}
			}

		}
	}
}

func sumBoard(board []int) int {
	sum := 0
	for _, n := range board {
		if n != -1 {
			sum += n
		}
	}
	return sum
}

func checkForWin(board []int, line []int) bool {
	// log.Printf("checking line %v", line)
	for _, pos := range line {
		if board[pos] != -1 {
			// log.Println("not a winner")
			return false
		}
	}
	return true
}

// called numbers, and an array of arrays (5x5 = 25)
func readInput(filename string) ([]int, [][]int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var called []int
	boards := [][]int{}
	thisBoard := make([]int, 25)
	lineCount := 0
	boardPos := 0
	for scanner.Scan() {
		s := scanner.Text()
		if lineCount == 0 {
			for _, c := range strings.Split(s, ",") {
				num, err := strconv.Atoi(c)
				if err != nil {
					log.Fatalf("Could not handle %s", c)
				}
				called = append(called, num)
			}
			lineCount += 1
			continue
		}

		if s == "" || s == "\n" {
			continue
		}

		// 5 numbers in a row
		fiveNums := strings.Fields(s)
		for _, c := range fiveNums {
			n, err := strconv.Atoi(strings.TrimSpace(c))
			if err != nil {
				log.Fatalf("Failed to convert %s to int: %s", s, err.Error())
			}
			thisBoard[boardPos] = n
			boardPos += 1

			if boardPos == 25 {
				boardPos = 0
				boards = append(boards, thisBoard)
				log.Printf("added board: %+v", thisBoard)
				thisBoard = make([]int, 25)
			}
		}
		lineCount += 1
	}
	return called, boards
}
