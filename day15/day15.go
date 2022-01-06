package day15

import (
	"bufio"
	"container/heap"
	"log"
	"math"
	"os"
	"strings"

	"slowteetoe.com/adventofcode2021/utils"
)

var cavern [][]int
var bestPath []Point
var totalRisk int

func Day15Part1() {
	cavern = readCave("day15/input.txt")
	// log.Printf("Cave: %v", cavern)
	a_star(Point{0, 0}, Point{99, 99}, Manhattan)
	log.Printf("Total risk: %d", totalRisk)
	// log.Printf("path: %v", bestPath)
}

func Day15Part2() {}

type Point struct {
	x, y int
}

func readCave(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	m := [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "")
		thisRow := []int{}
		for _, ch := range s {
			thisRow = append(thisRow, utils.ParseInt(ch))
		}
		m = append(m, thisRow)
	}
	return m
}

type Node struct {
	point    Point
	priority int
	index    int
}

func (n Node) neighbors() []Point {
	x := n.point.x
	y := n.point.y
	neighbors := []Point{}
	// right
	if x+1 < len(cavern[0]) {
		neighbors = append(neighbors, Point{x + 1, y})
	}
	// down
	if y+1 < len(cavern) {
		neighbors = append(neighbors, Point{x, y + 1})
	}
	// left
	if x-1 >= 0 {
		neighbors = append(neighbors, Point{x - 1, y})
	}
	// up
	if y-1 >= 0 {
		neighbors = append(neighbors, Point{x, y - 1})
	}
	return neighbors
}

func NewNode(point Point) Node {
	return Node{point, 1, 1}
}

// PriorityQueue implements heap.Interface
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	node := x.(Node)
	// log.Printf("Pushing %v onto queue", node)
	node.index = n
	*pq = append(*pq, &node)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	old[n-1] = nil
	node.index = -1
	*pq = old[0 : n-1]
	return node
}

func a_star(startingPoint Point, goalPoint Point, h func(Node, Node) int) {
	log.Println("Searching for viable path using A* algorithm....")
	pq := PriorityQueue{}
	heap.Init(&pq)

	start := NewNode(startingPoint)
	goal := NewNode(goalPoint)

	// cost of the cheapest path from start to n currently known
	gScore := map[Point]int{}
	gScore[start.point] = 0

	// for node n, fScore := gScore[n] + h(n)
	start.priority = gScore[start.point] + h(start, goal)

	// there's probably a better way than maintaining a map in addition to the priority queue, but this is easy enough
	inOpenSet := map[Point]bool{}
	inOpenSet[start.point] = true
	heap.Push(&pq, start)

	cameFrom := map[Point]Point{}

	for {
		// node in open set with lowest fScore
		current := heap.Pop(&pq).(*Node)
		delete(inOpenSet, current.point)
		// log.Printf("popped: %v (neighbors: %v)", current, current.neighbors())

		if current.point.x == goalPoint.x && current.point.y == goalPoint.y {
			// log.Printf("****************************** HIT THE GOAL **************************************")
			bestPath, totalRisk = reconstructPath(cameFrom, current.point)
			return
		}

		for _, neighborPoint := range current.neighbors() {

			d := d(neighborPoint)

			tentativeGScore := gScore[current.point] + d

			// log.Printf("examining neighbor=%v, tentative g-score: %d  ( gSCore[current.point]=%d + d(neighbor)=%d )", neighborPoint, tentativeGScore, gScore[current.point], d)

			if tentativeGScore < getOrMaxInt(gScore, neighborPoint) {

				// log.Println("this path is better than any previous one, record it")
				cameFrom[neighborPoint] = current.point
				// log.Printf("cameFrom[%v] = %v", neighborPoint, current.point)

				gScore[neighborPoint] = tentativeGScore
				// log.Printf("gScore[%v] = %d", neighborPoint, tentativeGScore)

				// fScore[neighbor] = tentativeGScore + h(neighbor, goal)

				if _, ok := inOpenSet[neighborPoint]; !ok {
					neighbor := NewNode(neighborPoint)
					neighbor.priority = tentativeGScore + h(neighbor, goal)
					// log.Printf("Neigbor was not yet in the openSet, adding: %v", neighbor)
					heap.Push(&pq, neighbor)
					inOpenSet[neighbor.point] = true
					// log.Printf("priority queue now holds: %d", pq.Len())
				}
			} else {
				// log.Printf("tenative g-score of %d does not beat %d", tentativeGScore, gScore[neighborPoint])
			}
			// log.Printf("Open set: %v", inOpenSet)
		}

		if pq.Len() == 0 {
			log.Fatalln("open set is empty, but did not reach goal. FAIL!")
			return
		}
	}

}

func getOrMaxInt(m map[Point]int, p Point) int {
	if val, ok := m[p]; ok {
		return val
	}
	return math.MaxInt

}

func d(p Point) int {
	d := cavern[p.y][p.x]
	// log.Printf("d(%d,%d) is %d", p.x, p.y, d)
	return d
}

// returns path and risk score
func reconstructPath(cameFrom map[Point]Point, current Point) ([]Point, int) {
	totalRisk := 0
	pathTaken := []Point{current}
	// log.Printf("Recreating path from %v using %v", current, cameFrom)
	for {
		if current.x == 0 && current.y == 0 {
			// log.Printf("done")
			return pathTaken, totalRisk
		}
		totalRisk += d(current)
		if val, ok := cameFrom[current]; ok {
			// log.Printf("current backtracks to %+v", val)
			current = val
		} else {
			log.Fatalf("No value for %v in cameFrom!", current)
		}
		// or maybe append and reverse?
		pathTaken = prepend(pathTaken, current)
		// log.Printf("backtracking from %v, total path is: %v", current, pathTaken)
	}
}

func prepend(path []Point, p Point) []Point {
	path = append(path, Point{})
	copy(path[1:], path)
	path[0] = p
	return path
}

func Manhattan(node, goal Node) int {
	dx := abs(node.point.x - goal.point.x)
	dy := abs(node.point.y - goal.point.y)
	return dx + dy
}

func abs(val int) int {
	if val >= 0 {
		return val
	}
	return val * -1
}
