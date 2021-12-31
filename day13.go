package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"slowteetoe.com/adventofcode2021/utils"
)

func Day13Part1() {
	o := readOrigami("13-input.txt")
	log.Printf("%s", o.String())
	log.Printf("Processing fold instruction: %s", o.foldInstructions[0])
	o.processFoldN(0)
	log.Printf("%s", o.String())
}

func Day13Part2() {
	o := readOrigami("13-input.txt")
	o.processFoldInstructions()
	o.display()
}

func readOrigami(filename string) OrigamiMap {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	m := map[Point]bool{}
	instr := []string{}
	for scanner.Scan() {
		raw := scanner.Text()
		if strings.HasPrefix(raw, "fold along ") {
			instr = append(instr, strings.ReplaceAll(raw, "fold along ", ""))
		} else if raw == "" {
			// log.Printf("delim line: %s", raw)
		} else {
			s := strings.Split(scanner.Text(), ",")
			m[NewPoint(s[0], s[1])] = true
		}
	}
	return OrigamiMap{m, instr}
}

type OrigamiMap struct {
	m                map[Point]bool
	foldInstructions []string
}

func (o OrigamiMap) String() string {
	return fmt.Sprintf("%d points: %v, \n%v", len(o.m), o.m, o.foldInstructions)
}

func (o OrigamiMap) processFoldN(n int) {
	cmd := strings.Split(o.foldInstructions[n], "=")
	switch cmd[0] {
	case "y":
		o.foldY(utils.ParseInt(cmd[1]))
	case "x":
		o.foldX(utils.ParseInt(cmd[1]))
	}
}

func (o OrigamiMap) processFoldInstructions() {
	for n := range o.foldInstructions {
		o.processFoldN(n)
	}
}

func (o OrigamiMap) display() {
	maxX := 0
	maxY := 0
	for k := range o.m {
		if k.x > maxX {
			maxX = k.x
		}
		if k.y > maxY {
			maxY = k.y
		}
	}
	var sb strings.Builder
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if o.m[Point{x, y}] {
				sb.WriteString("■")
			} else {
				sb.WriteString(" ")
			}
		}
		sb.WriteString("\n")
	}
	log.Printf("\n%s\n", sb.String())
}

func (o OrigamiMap) foldX(n int) {
	for k := range o.m {
		if k.x > n {
			dx := k.x - n
			newX := n - dx
			o.m[Point{newX, k.y}] = true
			delete(o.m, k)
			if debug {
				log.Printf("mapped %v to %d,%d and removed old key", k, newX, k.y)
			}
		}
	}
}

func (o OrigamiMap) foldY(n int) {
	for k := range o.m {
		if k.y > n {
			dy := k.y - n
			newY := n - dy
			o.m[Point{k.x, newY}] = true
			delete(o.m, k)
			if debug {
				log.Printf("mapped %v to %d,%d and removed old key", k, newY, k.y)
			}
		}
	}
}
