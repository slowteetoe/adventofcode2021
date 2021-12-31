package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"slowteetoe.com/adventofcode2021/utils"
)

func Day11Part1() {
	m := readOctopusMap("11-input.txt")
	total := 0
	life := OctopusLife{m, map[string]bool{}, &total}
	for x := 0; x < 100; x++ {
		log.Printf("\n%s", life.String())
		life.step()
	}
	log.Printf("\n%s", life.String())
	log.Printf("Total flashes: %d", total)
}

func Day11Part2() {
	m := readOctopusMap("11-input.txt")
	total := 0
	life := OctopusLife{m, map[string]bool{}, &total}
	for x := 1; ; x++ {
		log.Printf("\n%s", life.String())
		syn := life.step()
		if syn {
			log.Printf("All octopuses flashed at same time, step %d", x)
			log.Printf("\n%s", life.String())
			return
		}
	}
}

type OctopusLife struct {
	OctopusMap
	flashed    map[string]bool
	flashCount *int
}

func (o OctopusLife) String() string {
	return fmt.Sprintf("%s\nFlashes: %d", o.OctopusMap.String(), *o.flashCount)
}

func (o OctopusLife) step() bool {
	// bump up energy levels
	for y := range o.OctopusMap {
		for x := range o.OctopusMap[y] {
			o.OctopusMap[x][y] += 1
		}
	}
	// check for flashes
	checkForMoreFlashes := true
	for checkForMoreFlashes {
		checkForMoreFlashes = false
		for y := range o.OctopusMap {
			for x := range o.OctopusMap[y] {
				if o.OctopusMap[x][y] > 9 && !o.flashed[fmt.Sprintf("%d,%d", x, y)] {
					*o.flashCount += 1
					o.flash(x, y)
					checkForMoreFlashes = true
				}
			}
		}
	}
	flashesThisStep := len(o.flashed)
	// reset energies to 0
	for k := range o.flashed {
		xy := strings.Split(k, ",")
		o.OctopusMap[utils.ParseInt(xy[0])][utils.ParseInt(xy[1])] = 0
		delete(o.flashed, k)
	}
	return flashesThisStep == len(o.OctopusMap[0])*len(o.OctopusMap)
}

func (o OctopusLife) flash(x, y int) {
	o.flashed[fmt.Sprintf("%d,%d", x, y)] = true
	//up
	o.absorb(x, y-1)
	//updiagright
	o.absorb(x+1, y-1)
	//right
	o.absorb(x+1, y)
	//downdiagright
	o.absorb(x+1, y+1)
	//down
	o.absorb(x, y+1)
	//downdiagleft
	o.absorb(x-1, y+1)
	//left
	o.absorb(x-1, y)
	//updiagleft
	o.absorb(x-1, y-1)
}

func (o OctopusLife) absorb(x, y int) {
	if x < 0 || y < 0 || y >= len(o.OctopusMap) || x >= len(o.OctopusMap[0]) {
		return
	}
	// log.Printf("%d,%d just received more energy", x, y)
	o.OctopusMap[x][y] += 1
}

type OctopusMap [][]int

func (o OctopusMap) String() string {

	var sb strings.Builder
	for y := range o {
		for x := range o[y] {
			val := o[x][y]
			s := strconv.Itoa(val)
			if val == 0 {
				sb.WriteString("\x1b[43;33;1m0\x1b[0m")
			} else {
				sb.WriteString(s)
			}

		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func readOctopusMap(filename string) OctopusMap {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	m := [][]int{}
	for scanner.Scan() {
		s := scanner.Text()
		thisRow := []int{}
		for _, c := range strings.Split(s, "") {
			thisRow = append(thisRow, utils.ParseInt(c))
		}
		m = append(m, thisRow)
	}
	return m
}
