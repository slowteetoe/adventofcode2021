package day18

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/bradleyjkemp/memviz"
)

func Day18Part1() {
	snailfish := readSnailfish("day18/example.txt")
	log.Printf("%v", snailfish)
	l1 := parseSnailfish(snailfish[0])

	var sb strings.Builder
	print(l1, &sb)
	log.Printf("tree: %s", sb.String())

	// l2 := parseSnailfish(snailfish[1])
	// sb.Reset()
	// print(l2, &sb)

	// l3 := add(l1, l2)
	// sb.Reset()
	// print(l3, &sb)
	// log.Printf("result of adding l1 + l2: %+v", sb.String())

	buf := &bytes.Buffer{}
	memviz.Map(buf, &l1)
	err := ioutil.WriteFile("snailfish-tree", buf.Bytes(), 0644)
	if err != nil {
		panic(err)
	}

	// l1.reduce()
}

func Day18Part2() {}

func readSnailfish(filename string) []string {
	snailfish := []string{}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		snailfish = append(snailfish, scanner.Text())
	}
	return snailfish
}

func (n Node) reduce() {
	for {
		if n.explode() {
			log.Printf("something exploded")
			continue
		}
		if n.split() {
			log.Printf("something split")
			continue
		}
		log.Printf("done reducing")
		return
	}
}

func (n *Node) explode() bool {
	if n != nil {
		n.explodeRecursively(0)
		log.Printf("tree is now: %+v", n)
	}
	return false
}

func (n *Node) explodeRecursively(depth int) bool {
	if depth >= 4 && numericPair(n.left, n.right) {
		left, _ := n.left.(Num)
		right, _ := n.right.(Num)
		log.Printf("We're at depth %d, exploding %d, %d, parent=%+v", depth, left.val, right.val, n.parent)
		newNode := &Node{}
		if leftNum, ok := n.parent.left.(Num); ok {
			log.Printf("Adding %d to %d", left.val, leftNum.val)
			leftNum.val += left.val
			newNode.left = &leftNum
			newNode.right = &Num{0}
			n.parent.right = newNode
			return true
		} else {
			log.Printf("left was not a regular number, ignoring")
		}
		if rightNum, ok := n.parent.right.(Num); ok {
			log.Printf("Adding %d to %d", right.val, rightNum.val)
			rightNum.val += right.val
			newNode.left = &Num{0}
			newNode.right = &rightNum
			n.parent.left = *newNode
			return true
		} else {
			log.Printf("right was not a regular number, ignoring")
		}
		log.Printf("Something went horribly wrong!")
		return true
	}
	if left, ok := n.left.(Node); ok {
		log.Printf("going left to see if we can explode")
		result := left.explodeRecursively(depth + 1)
		if result {
			return true
		}
	}
	if right, ok := n.right.(Node); ok {
		log.Printf("going right to see if we can explode")
		result := right.explodeRecursively(depth + 1)
		if result {
			return true
		}
	}
	return false
}

func numericPair(left, right Any) bool {
	if left == nil || right == nil {
		return false
	}
	if _, ok := left.(Num); !ok {
		return false
	}
	if _, ok := right.(Num); !ok {
		return false
	}
	return true
}

func (n Node) split() bool {
	return false
}

type Any = interface{}

type Num struct {
	val int
}

type Node struct {
	parent      *Node
	left, right *Any
}

func add(a, b Node) Node {
	newRoot := Node{}
	a.parent = &newRoot
	newRoot.left = a
	b.parent = &newRoot
	newRoot.right = b
	return newRoot
}

func parseSnailfish(line string) Node {
	var stack Stack
	for _, c := range strings.Split(line, "") {
		log.Printf("next char is %s, stack is %v", c, stack)
		switch c {
		case ",":
			continue // eat commas
		case "]":
			log.Printf("should be poppin stuff...")
			r, _ := stack.Pop()
			l, _ := stack.Pop()
			pair := Node{}
			pair.left = l
			pair.right = r
			expected, _ := stack.Pop()
			if v, ok := expected.(string); !ok {
				log.Printf("Expected %s to be a left brace, but it wasn't!", v)
			}
			log.Printf("putting %+v back onto the stack", pair)
			stack.Push(pair)
			// continue //
		default:
			if c == "[" {
				stack.Push("[")
			} else {
				i, err := strconv.Atoi(c)
				if err != nil {
					log.Fatalf("attempted to convert %s to an int: %s", c, err.Error())
				}
				stack.Push(Num{i})
			}
		}
	}
	finalTree, _ := stack.Pop()
	if v, ok := finalTree.(Node); ok {
		log.Printf("Ok, returning %+v", v)
		return v
	}
	log.Fatalf("What I want to return is actually %+v, but something went horribly wrong!", finalTree)
	return Node{}
}

func print(root Node, sb *strings.Builder) {
	log.Printf("descending to %+v", root)
	if root.left != nil {
		if val, ok := root.left.(Num); ok {
			sb.WriteString(fmt.Sprintf("%d", val))
			log.Printf("%d", val)
		} else {
			if val, ok := root.left.(Node); ok {
				print(val, sb)
			}
		}
	} else {
		log.Printf("skipping left node, was nil")
	}
	if root.right != nil {
		if val, ok := root.right.(Num); ok {
			sb.WriteString(fmt.Sprintf("%d", val))
			log.Printf("%d", val)
		} else {
			if val, ok := root.right.(Node); ok {
				print(val, sb)
			}
		}
	} else {
		log.Printf("skipping right node, was nil")
	}
}

type Stack []Any

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str Any) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (Any, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}
