package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

// we'll have to keep track of visited points
var visited = map[string]bool{}

func Day9Part1() {
	heightMap := readHeightMap("09-input.txt")
	// log.Printf("Cave is %+v", heightMap)
	totalRisk := 0
	lowPoints := findLowPoints(heightMap)
	for _, i := range lowPoints {
		totalRisk += i.value + 1
	}
	log.Printf("%d low points. Total risk level %d", len(lowPoints), totalRisk)
}

func Day9Part2() {
	heightMap := readHeightMap("09-input.txt")
	lowPoints := findLowPoints(heightMap)
	// debug = true

	// starting from each low point, try visit all the points you can
	if debug {
		log.Printf("%d low points: %v", len(lowPoints), lowPoints)
	}

	allBasins := [][]string{}
	for _, lp := range lowPoints {
		basin := []string{}
		exploreCave(heightMap, lp.x, lp.y, &basin)
		if debug {
			log.Printf("basin from (%d,%d) is points: %v", lp.x, lp.y, basin)
		}
		allBasins = append(allBasins, basin)
	}

	basinSizes := []int{}
	for _, b := range allBasins {
		basinSizes = append(basinSizes, len(b))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))
	if debug {
		log.Printf("Sizes of basins: %v", basinSizes)
	}
	log.Printf("Total multiplied basin size (of top 3): %d", basinSizes[0]*basinSizes[1]*basinSizes[2])
}

func exploreCave(heightMap [][]int, x int, y int, basin *[]string) {
	thisSquare := fmt.Sprintf("(%d,%d)", x, y)
	if debug {
		log.Printf("Exploring from %s...", thisSquare)
	}
	if x < 0 || y < 0 || x >= len(heightMap[0]) || y >= len(heightMap) {
		if debug {
			log.Printf("%s is beyond the boundaries of cave", thisSquare)
		}
		return
	} else if heightMap[y][x] == 9 {
		if debug {
			log.Printf("%s is a high point", thisSquare)
		}
		return
	}
	if visited[thisSquare] {
		if debug {
			log.Printf("%s was already visited", thisSquare)
		}
		return
	}
	// otherwise, we'll count this as part of the basin, mark it visisted, and see where else we can travel
	visited[thisSquare] = true
	*basin = append(*basin, thisSquare)
	// go left
	if debug {
		log.Println("going left")
	}
	exploreCave(heightMap, x-1, y, basin)

	// go up
	if debug {
		log.Println("going up")
	}
	exploreCave(heightMap, x, y-1, basin)
	// go right
	if debug {
		log.Println("going right")
	}
	exploreCave(heightMap, x+1, y, basin)
	// go down
	if debug {
		log.Println("going down")
	}
	exploreCave(heightMap, x, y+1, basin)
	if debug {
		log.Printf("No where else to travel from %s, basin=%v", thisSquare, basin)
	}
}

func findLowPoints(heightMap [][]int) []LowPoint {
	lowPoints := []LowPoint{}
	for y, rows := range heightMap {
		for x, val := range rows {
			// look left
			if x >= 1 && rows[x-1] <= val {
				continue
			}
			// look up
			if y >= 1 && heightMap[y-1][x] <= val {
				continue
			}
			// look right
			if x < len(rows)-1 && rows[x+1] <= val {
				continue
			}
			// look down
			if y < len(heightMap)-1 && heightMap[y+1][x] <= val {
				continue
			}
			// and if we're here, it's a low point
			if debug {
				log.Printf("(%d,%d) is a low point with value %d + 1", x, y, val)
			}
			lowPoints = append(lowPoints, LowPoint{x, y, val})
		}
	}
	return lowPoints
}

type LowPoint struct {
	x, y  int
	value int
}

func readHeightMap(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	heightMap := [][]int{}
	for scanner.Scan() {
		s := scanner.Text()
		thisRow := []int{}
		for _, c := range strings.Split(s, "") {
			thisRow = append(thisRow, parseInt(c))
		}
		heightMap = append(heightMap, thisRow)
	}
	return heightMap
}
