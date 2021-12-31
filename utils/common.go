package utils

import (
	"log"
	"strconv"
)

func ParseInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Could not convert %s to int", s)
	}
	return val
}
