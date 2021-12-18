package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day2Part1() {
	file, err := os.Open("02-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	depth := 0
	pos := 0
	for scanner.Scan() {
		s := strings.Fields(scanner.Text())
		amt, err := strconv.Atoi(s[1])
		if err != nil {
			log.Fatal(err)
		}
		switch s[0] {
		case "forward":
			pos += amt
		case "up":
			depth -= amt
		case "down":
			depth += amt
		default:
			log.Fatal("didn't understand" + s[0])
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("day2 part1 :: depth=%d, hpos=%d (mult=%d)\n", depth, pos, depth*pos)
}

func Day2Part2() {
	file, err := os.Open("02-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	depth := 0
	pos := 0
	aim := 0
	for scanner.Scan() {
		s := strings.Fields(scanner.Text())
		amt, err := strconv.Atoi(s[1])
		if err != nil {
			log.Fatal(err)
		}
		switch s[0] {
		case "forward":
			pos += amt
			depth += (aim * amt)
		case "up":
			aim -= amt
		case "down":
			aim += amt
		default:
			log.Fatal("didn't understand" + s[0])
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("day2 part2 :: depth=%d, hpos=%d (mult=%d)\n", depth, pos, depth*pos)
}
