package day17

import (
	"log"
	"math"
	"os"
	"strings"

	"slowteetoe.com/adventofcode2021/utils"
)

// Parts 1 and 2
func Day17() {
	target := readTargetCoords("day17/input.txt")
	log.Printf("target: %+v", target)
	maxY := -1 * math.MaxInt
	var bestx int
	var besty int
	count := 0
	for i := 0; i < 5000; i++ {
	NextY:
		for j := -10000; j < 10000; j++ {
			vx := i
			vy := j
			x, y := 0, 0
			thisY := -1 * math.MaxInt
			// log.Printf("Checking vx=%d, vy=%d", vx, vy)
			for {
				// x increases by vx
				x += vx
				// y increases by vy
				y += vy

				if y > thisY {
					thisY = y
				}

				if target.inBounds(x, y) {
					log.Printf("BooM! %d, %d is within the target area", x, y)
					count++
					if thisY > maxY {
						bestx = i
						besty = j
						maxY = thisY
						log.Printf("%d, %d resulted in y of %d", x, y, maxY)
					}
					continue NextY
				}

				// vx goes toward 0 by 1, no change if 0
				if vx > 0 {
					vx -= 1
				} else if vx < 0 {
					vx += 1
				}
				// vy decreases by 1
				vy -= 1

				if x > target.maxX || y < target.minY {
					// overshot, not valid
					// log.Printf("checked and rejected!")
					continue NextY
				}
			}
		}
	}
	log.Printf("found %d, %d to be height %d, total valid initial velocities: %d", bestx, besty, maxY, count)
}

type Coords struct {
	minX int
	maxX int
	minY int
	maxY int
}

func (c Coords) inBounds(x, y int) bool {
	// log.Printf("checking %d,%d", x, y)
	return x >= c.minX && x <= c.maxX && y >= c.minY && y <= c.maxY
}

func readTargetCoords(filename string) Coords {
	b, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Couldn't read file: %s", err)
	}
	s := string(b)
	xidx := strings.Index(s, "x=")
	yidx := strings.Index(s, ", y=")
	xs := strings.Split(s[xidx+2:yidx], "..")
	ys := strings.Split(s[yidx+4:], "..")
	log.Printf("x coords: [%s], y coords: [%s]", xs, ys)
	return Coords{utils.ParseInt(xs[0]), utils.ParseInt(xs[1]), utils.ParseInt(ys[0]), utils.ParseInt(ys[1])}
}
