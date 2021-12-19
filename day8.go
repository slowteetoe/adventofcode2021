package main

import (
	"bufio"
	"log"
	"os"
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

func Day8Part2() {

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
		signals := [][]string{}
		outputs := []OutputSignal{}
		for _, sig := range strings.Fields(fields[0]) {
			signals = append(signals, strings.Split(sig, ""))
		}

		for _, out := range strings.Fields(fields[1]) {
			outputs = append(outputs, NewOutputSignal(strings.Split(out, "")))
		}
		entries = append(entries, NotebookEntry{signals, outputs})
	}
	return entries
}

type NotebookEntry struct {
	signals [][]string
	output  []OutputSignal
}

type OutputSignal struct {
	vals       []string
	decodedVal string
}

func NewOutputSignal(vals []string) OutputSignal {
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
