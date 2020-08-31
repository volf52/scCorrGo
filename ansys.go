package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	n := 8.0
	intn := int(n)
	streamEncoding := "bpe"

	abcdTable := generateUniqueTuples(intn)
	corrTypes := []string{"scc", "anderson", "dice", "jaccard", "ku2", "ochiai", "pearson", "sorensen", "ss2"}

	var wg sync.WaitGroup
	start := time.Now()
	for _, tp := range corrTypes {
		wg.Add(1)
		go ErrorWorker(abcdTable, n, intn, tp, streamEncoding, &wg)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Took %v\n", elapsed)
}
