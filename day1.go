package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func Day1Part1() {
	file, err := os.Open("01-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	count := 0
	last := math.MaxInt32
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		thisValue, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		if thisValue > last {
			count += 1
		}
		last = thisValue
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("day1 part1 :: found %d increasing values\n", count)
}

func Day1Part2() {
	file, err := os.Open("01-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	ints := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		ints = append(ints, val)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	count := 0
	for i := 3; i < len(ints); i++ {
		if ints[i-2]+ints[i-1]+ints[i] > ints[i-3]+ints[i-2]+ints[i-1] {
			count++
		}
	}
	fmt.Printf("day1 part2 :: found %d increasing values\n", count)
}
