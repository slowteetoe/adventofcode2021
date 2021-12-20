package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"sort"
	"strings"
)

func Day8Part1() {
	entries := readSignals("08-input.txt")
	// log.Printf("Notebook: %+v", entries)
	count := 0
	for _, e := range entries {
		for _, sig := range e.output {
			if sig.decodedVal == "1" || sig.decodedVal == "4" || sig.decodedVal == "7" || sig.decodedVal == "8" {
				count++
			}
		}
	}
	log.Printf("found %d knowns outputs", count)

}

var SevenSegmentPinLayout = []int{
	0b1110111, // 0
	0b0100100, // 1
	0b1011101, // 2
	0b1101101, // 3
	0b0101110, // 4
	0b1101011, // 5
	0b1111011, // 6
	0b0100101, // 7
	0b1111111, // 8
	0b1101111, // 9
}

func Day8Part2() {
	//  aaaa
	// b    c
	// b    c
	//  dddd
	// e    f
	// e    f
	//  gggg

	allCombinations := generateMappings([]string{"a", "b", "c", "d", "e", "f", "g"})

	log.Printf("generated %d combinations", len(allCombinations))
	// for n, combo := range allCombinations {
	// 	log.Printf("%d -> %v", n, combo)
	// }

	// we'll figure out the mapping of signal letter (e.g. "a") to the segment it turns on (e.g. ), then validate that the mapping is valid for all samples
	// example solution was:
	// decoder := Decoder{map[string]int{
	// 	"a": 2,
	// 	"b": 5,
	// 	"c": 6,
	// 	"d": 0,
	// 	"e": 1,
	// 	"f": 3,
	// 	"g": 4,
	// }}

	notebooks := readSignals("08-input.txt")

	// instead of searching 7! (5040) potential combinations, we could be smarter and figure out the top segment by looking at 7 and 1,
	// cutting it down to 6! or even figuring out the two rightmost segments, and really lowering the
	// search space, but tbh it's just easier to brute force it

	total := 0

	for _, notebook := range notebooks {
		for _, combination := range allCombinations {
			// log.Printf("%d trying out %v", i, combination)
			m := map[string]int{}
			for n, char := range combination {
				m[char] = n
			}
			decoder := Decoder{m}
			valid := decoder.validate(notebook.signals)
			if !valid {
				// log.Printf("Rejecting %v", m)
				continue
			}
			log.Printf("Found valid mapping for signal patterns in this notebook entry %v, now decoding observed values", m)
			thisVal := 0
			for i, obs := range notebook.output {
				thisVal += decoder.decode(obs.vals) * int(math.Pow10(3-i))
			}
			log.Printf("observed number: %d", thisVal)
			total += thisVal
		}
	}
	log.Printf("Total: %d", total)
	// valid := decoder.validate(notebook.signals)
	// if !valid {
	// 	log.Printf("Try again!")
	// } else {
	// 	log.Println("Valid mapping, now decoding the output signals")
	// 	for _, s := range notebook.output {
	// 		output := decoder.decode(s.vals)
	// 		if output == -1 {
	// 			log.Printf("Failed to decode the output %v", s.vals)
	// 		} else {
	// 			log.Printf("%v -> %d", strings.Join(s.vals, ""), output)
	// 		}
	// 	}
	// }
}

func generateMappings(arr []string) [][]string {
	var helper func([]string, int)
	res := [][]string{}

	helper = func(arr []string, n int) {
		if n == 1 {
			tmp := make([]string, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

type Decoder struct {
	m map[string]int
}

func (d Decoder) decode(s []string) int {
	mask := d.bitMaskFor(s)
	for i := range SevenSegmentPinLayout {
		if SevenSegmentPinLayout[i]^mask == 0 {
			return i
		}
	}
	return -1
}

func (d Decoder) bitMaskFor(chars []string) int {
	mask := 0
	for _, c := range chars {
		mask |= 1 << d.m[c]
	}
	return mask
}

func (d Decoder) validate(signals []ObservedValue) bool {
	seen := map[int]int{}
	for x := 0; x < 10; x++ {
		seen[x] = 0
	}

	valid := true

	for _, signal := range signals {
		mask := d.bitMaskFor(signal.vals)
		// log.Printf("mask for '%s' is %b", signal.original, mask)
		found := false
		for i := range checks {
			if SevenSegmentPinLayout[i]^mask == 0 {
				// log.Printf("Found it! It's a %d", i)
				seen[i] += 1
				found = true
			}
		}
		if !found {
			// log.Printf("no match for %s, invalid mapping", signal.original)
			valid = false
		}
	}

	for _, v := range seen {
		if v != 1 {
			// log.Printf("validation failed: saw %d instances of the number %d, expected 1", v, k)
			valid = false
		}
	}
	return valid
}

func readSignals(filename string) []NotebookEntry {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	entries := []NotebookEntry{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), "|")
		signals := []ObservedValue{}
		outputs := []OutputSignal{}
		for _, sig := range strings.Fields(fields[0]) {
			signals = append(signals, NewObservedValue(sig))
		}

		for _, out := range strings.Fields(fields[1]) {
			outputs = append(outputs, NewOutputSignal(out))
		}
		entries = append(entries, NewNotebookEntry(signals, outputs))
	}
	return entries
}

type NotebookEntry struct {
	signals []ObservedValue
	output  []OutputSignal
}

func NewNotebookEntry(signals []ObservedValue, outputs []OutputSignal) NotebookEntry {
	sort.Slice(signals, func(i, j int) bool {
		return len(signals[i].vals) < len(signals[j].vals)
	})
	return NotebookEntry{signals, outputs}
}

type ObservedValue struct {
	vals     []string
	original string
}

func NewObservedValue(sig string) ObservedValue {
	vals := strings.Split(sig, "")
	return ObservedValue{vals, sig}
}

type OutputSignal struct {
	vals       []string
	decodedVal string
}

func NewOutputSignal(val string) OutputSignal {
	vals := strings.Split(val, "")
	decodedVal := ""
	switch len(vals) {
	case 2:
		decodedVal = "1"
	case 3:
		decodedVal = "7"
	case 4:
		decodedVal = "4"
	case 7:
		decodedVal = "8"
	default:
		decodedVal = ""
	}
	return OutputSignal{vals, decodedVal}
}
