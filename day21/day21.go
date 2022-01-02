package day21

import (
	"fmt"
	"log"
)

func Day21Part1() {
	game := Game{[2]int{7, 6}, [2]int{0, 0}, 0}
	die := DiracDie{0, 1}
	for game.scores[0] < 1000 && game.scores[1] < 1000 {
		turn := game.turn % 2
		game.positions[turn] = move(game.positions[turn], die.roll())
		game.scores[turn] += game.positions[turn] + 1
		game.turn++
	}
	log.Printf("Result: %d", min(game.scores[0], game.scores[1])*die.rollCount)
}

// return the resulting position for a given roll
func move(playerPos int, roll int) int {
	newPos := (playerPos + roll) % 10
	return newPos
}

type Game struct {
	positions [2]int
	scores    [2]int
	turn      int
}

func (g *Game) key() string {
	return fmt.Sprintf("%d-%d-%d-%d-%d", g.positions[0], g.positions[1], g.scores[0], g.scores[1], g.turn)
}

func (g *Game) clone() Game {
	return Game{[2]int{g.positions[0], g.positions[1]}, [2]int{g.scores[0], g.scores[1]}, g.turn}
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

var outcomeCache = map[string][2]int{}

func Day21Part2() {
	// log.Printf("444356092776315 + 341960390180808 = %d", 444356092776315+341960390180808)
	// example 4,8 but 0-based
	stats := [2]int{}
	game := Game{[2]int{3, 7}, [2]int{0, 0}, 0}
	game.play(&stats)
	log.Printf("stats: %v", stats)
}

func (game *Game) play(stats *[2]int) {

	for roll, count := range rollPossibilities() {
		player := game.turn % 2
		game.positions[player] += (game.positions[player] + roll) % 10
		game.scores[player] += game.positions[player] + 1
		game.turn++
		if game.scores[player] >= 21 {
			stats[player] += count
		} else {
			if val, ok := outcomeCache[game.key()]; ok {
				log.Printf("already saw this game play out! %s (stats %v)", game.key(), val)
				stats[0] += val[0] * count
				stats[1] += val[1] * count
				return
			}
			quantumGame := game.clone()
			altStats := [2]int{}
			quantumGame.play(&altStats)
			for n := range altStats {
				stats[n] += count * altStats[n]
			}
			outcomeCache[game.key()] = [2]int{altStats[0], altStats[1]}
		}
	}
}

func rollPossibilities() map[int]int {
	freq := map[int]int{}
	// 3 rolls of 3 possibilites
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				freq[i+j+k] += 1
			}
		}
	}
	// log.Printf("%v", freq)
	return freq
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
