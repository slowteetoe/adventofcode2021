package day12

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var cavePaths = [][]string{}
var caveSystem = map[string]Cave{}

func Day12Part1() {
	caveSystem = readCaveMap("day12/input.txt")
	spelunk(caveSystem["start"], []string{})
	log.Printf("Number of Paths: %d", len(cavePaths))
}

func spelunk(node Cave, visited []string) {
	// log.Printf("(%v -> %s)", visited, node.symbol)
	if strings.ToLower(node.symbol) == node.symbol {
		for _, s := range visited {
			if s == node.symbol {
				// log.Printf("Invalid route, lowercase cave %s already visited in %v", node.symbol, visited)
				return
			}
		}
	}
	visited = append(visited, node.symbol)
	if node.symbol == "end" {
		log.Printf("journey complete: %s", strings.Join(visited, "->"))
		cavePaths = append(cavePaths, visited)
		return
	}
	for _, c := range node.connections {
		// log.Printf("about to visit %s (from %s)", c, node.symbol)
		spelunk(caveSystem[c], visited)
	}
}

type Cave struct {
	symbol      string
	connections []string
}

func NewCave(s string) Cave {
	return Cave{s, []string{}}
}

func readCaveMap(filename string) map[string]Cave {
	caveSystem := map[string]Cave{}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		caves := strings.Split(scanner.Text(), "-")
		log.Printf("Joining %s and %s", caves[0], caves[1])
		var cave0 Cave
		var cave1 Cave
		if c0, ok := caveSystem[caves[0]]; !ok {
			// log.Printf("haven't seen %s yet, creating...", caves[0])
			cave0 = NewCave(caves[0])
		} else {
			// log.Printf("using %s", c0)
			cave0 = c0
		}
		if c1, ok := caveSystem[caves[1]]; !ok {
			// log.Printf("haven't seen %s yet, creating...", caves[1])
			cave1 = NewCave(caves[1])
		} else {
			// log.Printf("using %s", c1)
			cave1 = c1
		}
		cave0.connections = append(cave0.connections, cave1.symbol)
		cave1.connections = append(cave1.connections, cave0.symbol)
		caveSystem[cave0.symbol] = cave0
		caveSystem[cave1.symbol] = cave1
		// log.Printf("cavesystem=%v", caveSystem)
	}
	return caveSystem
}
