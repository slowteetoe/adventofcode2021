package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day5Part2() {
	// Not yet implemented
}

func Day5Part1() {
	lines, maxDimension := readLines("05-input.txt")
	gridSize := maxDimension + 1 // since grid is zero-based, we need a 8x8 grid to hold (0,7)
	log.Printf("Need a %dx%d grid to hold the input", gridSize, gridSize)
	grid := newGrid(gridSize)
	for _, line := range lines {
		if line.Diagonal() {
			// log.Printf("Ignoring diagonal line %s", line.String())
			continue
		}
		for _, p := range line.Points() {
			grid.mark(p)
		}
	}
	if gridSize < 120 { // otherwise too wide to display
		log.Printf("\n%s\n", grid.String())
	}
	log.Printf("answer: %d", grid.OverlapCount())
}

func readLines(filename string) ([]Line, int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := []Line{}
	scanner := bufio.NewScanner(file)

	// find the highest dimension of the grid
	maxx := 0
	maxy := 0
	for scanner.Scan() {
		s := scanner.Text()
		fields := strings.Fields(s)
		if len(fields) != 3 {
			log.Fatalf("expected 3 fields, but got %d", len(fields))
		}
		start := strings.Split(fields[0], ",")
		end := strings.Split(fields[2], ",")
		maxx = Max(maxx, start[0], end[0])
		maxy = Max(maxy, start[1], end[1])
		lines = append(lines, NewLine(NewPoint(start[0], start[1]), NewPoint(end[0], end[1])))
	}
	return lines, max(maxx, maxy)
}

type Grid struct {
	dim    int
	points []int
}

func newGrid(gridSize int) Grid {
	return Grid{gridSize, make([]int, gridSize*gridSize)}
}

func (g Grid) mark(p Point) {
	cell := p.y*g.dim + p.x
	g.points[cell] += 1
}

func (g Grid) String() string {
	var sb strings.Builder
	for i := range g.points {
		if i > 0 && i%g.dim == 0 {
			sb.WriteString("\n")
		}
		displayChar := fmt.Sprintf("%d", g.points[i])
		if displayChar == "0" {
			displayChar = "."
		}
		sb.WriteString(displayChar)
	}
	return sb.String()
}

func (g Grid) OverlapCount() int {
	overlap := 0
	for _, n := range g.points {
		if n > 1 {
			overlap++
		}
	}
	return overlap
}

type Line struct {
	start, end Point
}

func (l Line) String() string {
	return fmt.Sprintf("Line from %s to %s", &l.start, &l.end)
}

func (l Line) Diagonal() bool {
	return l.start.x != l.end.x && l.start.y != l.end.y
}

func (l Line) Points() []Point {
	points := []Point{}
	if l.Diagonal() {
		log.Printf("I don't know how to deal with diagonals right now")
	} else if l.start.x == l.end.x {
		// horizontal line
		for dy := min(l.start.y, l.end.y); dy <= max(l.start.y, l.end.y); dy++ {
			points = append(points, Point{l.start.x, dy})
		}
	} else if l.start.y == l.end.y {
		// vertical line
		for dx := min(l.start.x, l.end.x); dx <= max(l.start.x, l.end.x); dx++ {
			points = append(points, Point{dx, l.start.y})
		}
	}
	return points
}

type Point struct {
	x, y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

func NewPoint(x string, y string) Point {
	return Point{parseInt(x), parseInt(y)}
}

func parseInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Could not convert %s to int", s)
	}
	return val
}

func NewLine(start Point, end Point) Line {
	return Line{start, end}
}

func Max(x int, y string, z string) int {
	return max(max(x, parseInt(y)), parseInt(z))
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
