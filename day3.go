package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day3Part1() {
	file, err := os.Open("03-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var counts []int
	for scanner.Scan() {
		s := scanner.Text()
		if counts == nil {
			log.Printf("building a slice of size: %d", len(s))
			counts = make([]int, len(s))
		}

		line := strings.Split(s, "")

		for i, ch := range line {
			if ch == "0" {
				counts[i] -= 1
			} else if ch == "1" {
				counts[i] += 1
			} else {
				log.Fatal("Something other than a 1 or 0 ->" + ch)
			}
		}
	}
	log.Printf("after scanning: %x", counts)

	gamma := 0
	epsilon := 0
	for _, i := range counts {
		gamma = gamma << 1
		epsilon = epsilon << 1
		if i > 0 {
			// more ones than zeros
			gamma += 1
		} else {
			epsilon += 1
		}
	}

	log.Printf("gamma=%b (%d), epsilon=%b (%d), power consumption=%d", gamma, gamma, epsilon, epsilon, gamma*epsilon)
}

func Day3Part2() {
	file, err := os.Open("03-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	src := []string{}
	for scanner.Scan() {
		s := scanner.Text()
		src = append(src, s)
	}

	var oxyValue int64
	var co2Value int64

	oxyDepth := 0
	oxy := split(MostCommon, oxyDepth, src)
	oxyDepth++

	for {
		if len(oxy) == 1 {
			log.Printf("Found oxygen value: %s\n", oxy[0])
			oxyValue, err = strconv.ParseInt(oxy[0], 2, 64)
			if err != nil {
				log.Fatalf("Could not convert to int64: %v", err)
			}
			break
		}
		oxy = split(MostCommon, oxyDepth, oxy)
		oxyDepth++
	}

	co2Depth := 0
	co2 := split(LeastCommon, co2Depth, src)
	co2Depth++
	for {
		if len(co2) == 1 {
			log.Printf("Found co2 value: %s\n", co2[0])
			co2Value, err = strconv.ParseInt(co2[0], 2, 64)
			if err != nil {
				log.Fatalf("Could not convert to int64: %v", err)
			}
			break
		}
		co2 = split(LeastCommon, co2Depth, co2)
		co2Depth++
	}

	log.Printf("oxy=%d, co2=%d, result=%d\n", oxyValue, co2Value, oxyValue*co2Value)
}

type BitPref int

const (
	MostCommon BitPref = iota
	LeastCommon
)

//
func split(dir BitPref, pos int, src []string) []string {
	zeros := []string{}
	ones := []string{}
	for _, s := range src {
		if s[pos:pos+1] == "0" {
			// log.Printf("%s has leading 0 in pos %d", s, pos)
			zeros = append(zeros, s)
		} else {
			// log.Printf("%s has leading 1 in pos %d", s, pos)
			ones = append(ones, s)
		}
	}
	switch dir {
	case MostCommon:
		if len(zeros) > len(ones) {
			// log.Println("more zeros")
			return zeros
		} else {
			// log.Println("more ones")
			return ones
		}
	case LeastCommon:
		if len(zeros) <= len(ones) {
			// log.Println("less zeros")
			return zeros
		} else {
			// log.Println("less ones")
			return ones
		}
	default:
		log.Fatalf("could not determine")
		return []string{}
	}
}
