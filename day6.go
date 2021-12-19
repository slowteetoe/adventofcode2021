package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Day6Part1() {
	fish := readFishies("06-input.txt")
	days := 80
	log.Printf("There are %d fish: %v", len(fish), fish)

	for n := 1; n <= days; n++ {
		tomorrow := []int{}
		for _, f := range fish {
			if f == 0 {
				tomorrow = append(tomorrow, 6)
				tomorrow = append(tomorrow, 8)
			} else {
				tomorrow = append(tomorrow, f-1)
			}
		}
		fish = tomorrow
		log.Printf("After %2d day(s) there are: %d fish", n, len(fish))
	}
}

func Day6Part2() {
	fish := readFishies("06-input.txt")
	days := 256
	log.Printf("There are %d fish: %v", len(fish), fish)

	// just keep counts, don't need individual fishes
	timer := make([]int, 9)
	for _, f := range fish {
		timer[f] += 1
	}

	for n := 1; n <= days; n++ {
		spawning := timer[0]
		timer[0] = timer[1]
		timer[1] = timer[2]
		timer[2] = timer[3]
		timer[3] = timer[4]
		timer[4] = timer[5]
		timer[5] = timer[6]
		timer[6] = timer[7] + spawning
		timer[7] = timer[8]
		timer[8] = spawning

		sum := int(0)
		for n := range timer {
			sum += timer[n]
		}
		log.Printf("After %2d day(s) there are: %d fish", n, sum)
	}
}

func readFishies(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fish := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), ",")
		for _, f := range fields {
			fish = append(fish, parseInt(f))
		}
	}
	return fish
}
