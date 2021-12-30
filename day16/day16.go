package day16

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var expMap = map[rune]string{
	'0': "0000",
	'1': "0001",
	'2': "0010",
	'3': "0011",
	'4': "0100",
	'5': "0101",
	'6': "0110",
	'7': "0111",
	'8': "1000",
	'9': "1001",
	'A': "1010",
	'B': "1011",
	'C': "1100",
	'D': "1101",
	'E': "1110",
	'F': "1111",
}

func Day16Part1() {
	input, err := readString("day16/input.txt")
	if err != nil {
		log.Fatalf("Could not read input: %s", err)
	}
	log.Printf("Src: %s", input)
	packets := &[]Packet{}
	parsePackets(input, 0, packets)
	log.Printf("Packets Read: %+v", packets)
	total := 0
	for _, n := range *packets {
		log.Printf("looking at %v with version=%d, total=%d", n, n.VersionNum(), total)
		total += n.VersionNum()
		if o, ok := n.(OperatorPacket); ok {
			for _, sub := range o.packets {
				log.Printf("|-- looking at %v, total=%d", sub, total)
				total += sub.VersionNum()
			}
		}
	}
	log.Printf("total packet version sum: %d", total)
}

type Packet interface {
	Value() int
	VersionNum() int
}

type OperatorPacket struct {
	versionNum int
	packets    []Packet
}

func (o OperatorPacket) Value() int {
	return -1
}

func (o OperatorPacket) VersionNum() int {
	return o.versionNum
}

type Type4Packet struct {
	versionNum int
	val        int
}

func (t Type4Packet) Value() int {
	return t.val
}

func (t Type4Packet) VersionNum() int {
	return t.versionNum
}

func readString(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		return expandToBinary(scanner.Text()), nil
	}
	return "", fmt.Errorf("failed to read string")
}

func expandToBinary(s string) string {
	exp := ""
	for _, ch := range s {
		exp += expMap[ch]
	}
	return exp
}

func parsePackets(input string, startingPos int, packets *[]Packet) int {
	if startingPos > len(input) {
		log.Fatalf("something went wrong! trying to parse packets from pos=%d, len=%d", startingPos, len(input))
	}
	log.Printf("Attempting to decode packets from: %s", input[startingPos:])
	x := startingPos
	version := binaryStringToInt(input[x : x+3])
	x += 3
	id := binaryStringToInt(input[x : x+3])
	x += 3
	log.Printf("Reading version=%d, type id: %d", version, id)
	if id == 4 {
		log.Printf("parsing type 4 packet...")
		bits := ""
		for {
			log.Printf("looking at: %s", input[x:x+5])
			cont := input[x : x+1]
			x += 1
			bits += input[x : x+4]
			x += 4
			if cont == "0" {
				break
			}
		}
		val := binaryStringToInt(bits)
		log.Printf("Read the value: %d from the bits at input[%d : %d]", val, startingPos, x)
		*packets = append(*packets, Type4Packet{version, val})
		return x - startingPos
	} else {
		log.Printf("Operator packet")
		lengthType := input[x : x+1]
		x += 1
		if lengthType == "0" {
			lengthOfSubpackets := binaryStringToInt(input[x : x+15])
			x += 15
			log.Printf("Goal Subpackets length: %d", lengthOfSubpackets)
			l := 0
			log.Printf("before looping: len(input)=%d x=%d, L=%d, subpacketLength=%d", len(input), x, l, lengthOfSubpackets)
			subpackets := &[]Packet{}
			for {
				read := parsePackets(input, x, subpackets)
				l += read
				x += read
				log.Printf("x=%d, L=%d, read=%d", x, l, read)
				if l >= lengthOfSubpackets {
					log.Printf("read all subpackets we were supposed to read...")
					*packets = append(*packets, OperatorPacket{version, *subpackets})
					return x
				}
			}
		} else {
			numSubPackets := binaryStringToInt(input[x : x+11])
			x += 11
			log.Printf("Number of subpackets contained: %d", numSubPackets)
			subpackets := &[]Packet{}
			for {
				read := parsePackets(input, x, subpackets)
				x += read
				if len(*subpackets) >= numSubPackets {
					break
				}
			}
			*packets = append(*packets, OperatorPacket{version, *subpackets})
			return x
		}
	}
}

func binaryStringToInt(s string) int {
	i, err := strconv.ParseInt(s, 2, 0)
	if err != nil {
		log.Fatalf("Could not convert %s to an int", s)
	}
	return int(i)
}
