package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day4Part1() {
	called, boards := readInput("04-example.txt")
	log.Printf("calledNumbers: %v, num boards: %d", called, len(boards))
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
