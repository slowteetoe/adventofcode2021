package main

import (
	"bufio"
	"fmt"
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
			for i, spot := range board.spots {
				if spot == number {
					board.spots[i] = -1
					break
				}
			}
			if callCount < 5 {
				continue
			}
			// see if there's a winning line
			for _, line := range checks {
				// log.Printf("checking board id=%d", boardId)
				win := checkForWin(board.spots, line)
				if win {
					sum := sumBoard(board.spots)
					log.Printf("*** WINNER (via %v) *** sum=%d, last number called=%d, answer=%d", line, sum, number, sum*number)
					return
				}
			}

		}
	}
}

func Day4Part2() {
	called, boards := readInput("04-example.txt")
	remaining := len(boards)
	log.Printf("calledNumbers: %v, num boards: %d", called, remaining)
	for _, number := range called {
		log.Printf("Calling %d", number)
	RecheckBoards:
		for _, board := range boards {
			if board.bingo {
				log.Println("skipping bingo'd board")
				continue
			}
			// otherwise check and mark the spot
			for i, spot := range board.spots {
				if spot == number {
					board.spots[i] = -1
					break
				}
			}

			// see if there's a winning line
			for _, line := range checks {
				// log.Printf("checking board id=%d", boardId)
				win := checkForWin(board.spots, line)
				if win {
					board.bingo = true
					sum := sumBoard(board.spots)

					if remaining == 1 {
						log.Printf("*** ULTIMATE WINNER (via %v) *** sum=%d, last number called=%d, answer=%d", line, sum, number, sum*number)
						log.Println(board.String())
						return
					} else {
						log.Printf("*** BINGO (via %v) *** sum=%d, last number called=%d, answer=%d", line, sum, number, sum*number)

						// take this board out of the running
						log.Printf("Taking board out of the running (curr: %d) ...", remaining)

						// tmp := []Board{}
						// // THIS PART IS BROKEN!
						// for _, b := range boards {
						// 	if !b.bingo {
						// 		tmp = append(tmp, b)
						// 	} else {
						// 		log.Printf("NOT ADDING WINNING BOARD\n%v\n", b)
						// 	}
						// }
						// boards = tmp
						remaining -= 1
						log.Printf("%d non-winning boards remaining.", remaining)
						if remaining <= 5 {
							for _, b := range boards {
								if !b.bingo {
									log.Println(b.String())
								}
							}
						}
						continue RecheckBoards
					}
				}
			}

		}
	}
}

type Board struct {
	bingo bool
	spots []int
}

func NewBoard() Board {
	return Board{
		bingo: false,
		spots: make([]int, 25),
	}
}

func (b *Board) String() string {
	var sb strings.Builder
	if b.bingo {
		sb.WriteString("[BINGO'D]\n")
	}
	for pos, spot := range b.spots {
		if pos%5 == 0 {
			sb.WriteString("\n")
		}
		sb.WriteString("\t" + numOrMark(spot))
	}
	return sb.String()
}

func numOrMark(spot int) string {
	if spot == -1 {
		return "X"
	}
	return fmt.Sprintf("%d", spot)
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
func readInput(filename string) ([]int, []Board) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var called []int
	boards := []Board{}
	thisBoard := NewBoard()
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
			thisBoard.spots[boardPos] = n
			boardPos += 1

			if boardPos == 25 {
				boardPos = 0
				boards = append(boards, thisBoard)
				thisBoard = NewBoard()
			}
		}
		lineCount += 1
	}
	return called, boards
}
