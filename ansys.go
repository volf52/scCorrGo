package main

import (
	"log"
	"os"
	"strconv"
)

func main() {
	args := os.Args

	if len(args) != 2{
		log.Fatal("Usage: ./{nameOfCliFile} {n}")
	}

	n, err := strconv.ParseInt(args[1], 10, 64)
	intn := int(n)

	if err != nil{
		log.Fatal(err)
	}

	abcdTable := generate_unique_tuples(intn)

	calculateCorrelations(abcdTable, float64(n), intn, true)
}
