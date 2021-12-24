package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func Day14Part1() {
	p := readPolymerInfo("14-input.txt")
	for n := 0; n < 10; n++ {
		p.synthesize()
	}
	_, most, least := p.freqStats()
	fmt.Printf("answer: %d\n", most-least)
}
func Day14Part2() {
	p := readPolymerInfo("14-input.txt")
	// fix me, need a different algorithm
	for n := 0; n < 40; n++ {
		p.synthesize()
	}
	_, most, least := p.freqStats()
	fmt.Printf("answer: %d\n", most-least)
}

func readPolymerInfo(filename string) PolymerInfo {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var formula string
	rules := map[string]string{}
	firstLine := true
	for scanner.Scan() {
		s := scanner.Text()
		if firstLine {
			formula = s
			firstLine = false
			continue
		}
		if s == "" {
			continue
		}
		ruleParts := strings.Split(s, " -> ")
		rules[ruleParts[0]] = ruleParts[1]
	}
	return PolymerInfo{formula, rules}
}

type PolymerInfo struct {
	currentPolymer string
	rules          map[string]string
}

func (p *PolymerInfo) synthesize() {
	newPolymer := ""
	elements := strings.Split(p.currentPolymer, "")
	for n := 0; n < len(elements)-1; n++ {
		if debug {
			log.Printf("rules[%s] -> %s", elements[n]+elements[n+1], p.rules[elements[n]+elements[n+1]])
		}
		newPolymer += elements[n] + p.rules[elements[n]+elements[n+1]]
	}
	newPolymer += elements[len(elements)-1]
	if debug {
		log.Printf("%s -> %s\n", p.currentPolymer, newPolymer)
	}
	p.currentPolymer = newPolymer
}

// m[polymer] -> count, most common count, least common count
func (p *PolymerInfo) freqStats() (map[string]int, int, int) {
	m := map[string]int{}
	for _, ch := range strings.Split(p.currentPolymer, "") {
		m[ch] += 1
	}
	most := 0
	least := math.MaxInt
	for _, v := range m {
		if v > most {
			most = v
		}
		if v < least {
			least = v
		}
	}
	return m, most, least
}
