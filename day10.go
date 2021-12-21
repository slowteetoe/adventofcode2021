package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Day10Part1() {
	file, err := os.Open("10-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	m := map[string]string{")": "(", "]": "[", "}": "{", ">": "<"}

	scanner := bufio.NewScanner(file)

	score := 0
	for scanner.Scan() {
		stack := Stack{}
		line := scanner.Text()
		log.Println(line)
		for _, c := range strings.Split(line, "") {
			switch c {
			case "(", "[", "{", "<":
				// log.Printf("Pushing %s", c)
				stack.Push(c)
			default:
				// log.Printf("Dealing with %s", c)
				actual, _ := stack.Pop()
				if m[c] != actual {
					log.Printf("Illegal character found, expected %s, but found %s instead.", actual, c)
					score += scoreForIllegalChar(c)
					break
				}
			}
		}
		if !stack.IsEmpty() {
			actual, _ := stack.Pop()
			log.Printf("(Ignore for part 1) Incomplete line, char on stack: %s", actual)
		}
		log.Printf("Final score: %d", score)
	}
}

func scoreForIllegalChar(s string) int {
	switch s {
	case ")":
		return 3
	case "]":
		return 57
	case "}":
		return 1197
	case ">":
		return 25137
	default:
		log.Fatalf("Invalid char, no idea how to handle %s", s)
	}
	return -1
}

func Day10Part2() {}

// LIFO stack
type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // find index last element
		element := (*s)[index] // Obtain the element.
		*s = (*s)[:index]      // Remove it
		return element, true
	}
}
