package main

import (
	"fmt"
	"log"
)

func main() {
	n := 128.0
	intn := int(n)
	pth := fmt.Sprintf("abcd_%v.csv", intn)
	abcdTable, err := parseCsv(pth)

	if err != nil {
		log.Fatal(err)
	}

	calculateCorrelations(&abcdTable, n, intn,true)
}
