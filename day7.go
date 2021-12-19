package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"sort"
	"strings"
)

func Day7Part1() {
	crabs := readCrabs("07-input.txt")
	median := medianIndex(crabs)
	log.Printf("Using the median: %d", median)

	bestCost := math.MaxInt
	bestPos := median

	// flail around aimlessly, but the best should be somewhere around the median, I think?
	for n := median - len(crabs)/3; n < median+len(crabs)/3; n++ {
		totalCost := 0
		for c := range crabs {
			cost := abs(crabs[n] - crabs[c])
			totalCost += cost
		}
		if totalCost < bestCost {
			// log.Printf("found a candidate, moving to horizontal pos %d cost=%d", crabs[n], totalCost)
			bestCost = totalCost
			bestPos = crabs[n]
		}
	}
	log.Printf("Best cost: %d, achieved with horizontal position: %d", bestCost, bestPos)
}

func Day7Part2() {
	crabs := readCrabs("07-input.txt")

	sort.Ints(crabs)

	maxPos := crabs[len(crabs)-1]

	var bestCost int64 = math.MaxInt64
	var bestPos int

	for n := 0; n < maxPos; n++ {
		var totalCost int64 = 0
		for c := range crabs {
			absDistance := abs(n - crabs[c])
			// log.Printf("For crab %d to move from %d to %d is distance %d", c, n, crabs[c], absDistance)
			// 1+2+3+4...
			cost := int64(absDistance * (absDistance + 1) / 2)
			totalCost += cost
		}
		if totalCost < bestCost {
			// log.Printf("found a candidate, moving to horizontal pos %d cost=%d", n, totalCost)
			bestCost = totalCost
			bestPos = n
		}
	}
	log.Printf("Best cost: %d, achieved with horizontal position: %d", bestCost, bestPos)
}

func medianIndex(arr []int) int {
	sort.Ints(arr)
	odd := len(arr)%2 == 1
	midpoint := len(arr) / 2
	var median int
	if odd {
		median = arr[midpoint]
	} else {
		median = arr[midpoint-1] + arr[midpoint]/2
	}
	return median
}

func readCrabs(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	crabs := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), ",")
		for _, f := range fields {
			crabs = append(crabs, parseInt(f))
		}
	}
	return crabs
}
