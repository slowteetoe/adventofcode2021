package day21

import (
	"log"
)

func Day21Part1() {
	p1Score := 0
	p2Score := 0
	p1 := 7 // zero based, so pos 4 -> 3
	p2 := 6
	turns := 0
	die := DiracDie{0, 1}
	for p1Score < 1000 && p2Score < 1000 {
		roll := die.roll()
		if turns%2 == 0 {
			// move p1
			p1 = move(p1, roll)
			p1Score += p1 + 1
			// log.Printf("player 1 new position: %d for a total score of %d", p1, p1Score)
		} else {
			// move p2
			p2 = move(p2, roll)
			p2Score += p2 + 1
			// log.Printf("player 2 new position: %d for a total score of %d", p1, p1Score)
		}
		turns++
	}
	// log.Printf("Score [p1=%d | p2=%d] in %d rolls", p1Score, p2Score, die.rollCount)
	var result int
	if p1Score < p2Score {
		result = p1Score * die.rollCount
	} else {
		result = p2Score * die.rollCount
	}

	log.Printf("Result: %d", result)
}

// return the resulting position for a given roll
func move(playerPos int, roll int) int {
	newPos := (playerPos + roll) % 10
	return newPos
}

type DiracDie struct {
	rollCount int
	curr      int
}

func (d *DiracDie) roll() int {
	d.rollCount += 3
	// log.Printf("Player rolls %d+%d+%d", d.curr, d.curr+1, d.curr+2)
	val := d.curr + d.curr + 1 + d.curr + 2
	d.curr = d.curr + 3
	return val
}

func Day21Part2() {}
