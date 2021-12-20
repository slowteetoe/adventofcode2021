package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Day9Part1() {
	heightMap := readHeightMap("09-input.txt")
	// log.Printf("Cave is %+v", heightMap)
	numLowPoints := 0
	totalRisk := 0
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
			log.Printf("(%d,%d) is a low point with value %d + 1", x, y, val)
			numLowPoints += 1
			totalRisk += val + 1
		}
	}
	log.Printf("%d low points. Total risk level %d", numLowPoints, totalRisk)
}

func Day9Part2() {}

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
