package main

import (
	"bufio"
	"log"
	"os"
	"sort"
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

func Day10Part2() {
	file, err := os.Open("10-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scores := []int{}
	for scanner.Scan() {
		score, corrupted := correctLine(scanner.Text())
		if !corrupted {
			scores = append(scores, score)
		}
	}
	sort.Ints(scores)
	log.Printf("Final score: %d  (%v)", scores[len(scores)/2], scores)
}

func correctLine(line string) (int, bool) {
	completions := map[string]string{"(": ")", "[": "]", "{": "}", "<": ">"}
	m := map[string]string{")": "(", "]": "[", "}": "{", ">": "<"}
	log.Printf("attempting to correct: %s", line)
	stack := Stack{}
	for _, c := range strings.Split(line, "") {
		switch c {
		case "(", "[", "{", "<":
			// log.Printf("Pushing %s", c)
			stack.Push(c)
		default:
			// log.Printf("Dealing with %s", c)
			actual, _ := stack.Pop()
			if m[c] != actual {
				log.Printf("Corrupted line, discarding %s", line)
				return 0, true
			}
		}
	}
	log.Printf("%v", stack)
	if stack.IsEmpty() {
		log.Fatalf("Should have had an incomplete line here!")
		return -1, true
	}
	log.Printf("Incomplete line: %s\nstack has %d elements: %v", line, len(stack), stack)
	completion := ""
	for !stack.IsEmpty() {
		c, _ := stack.Pop()
		log.Printf("Looking at %s and closing with %s", c, completions[c])
		completion += completions[c]
	}
	log.Printf("Complete by adding %s", completion)
	return scoreForCompletion(completion), false
}

func scoreForCompletion(s string) int {
	score := 0
	for _, c := range strings.Split(s, "") {
		score *= 5
		switch c {
		case ")":
			score += 1
		case "]":
			score += 2
		case "}":
			score += 3
		case ">":
			score += 4
		}
	}
	return score
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
