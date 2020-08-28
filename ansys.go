package main

import "fmt"

func main() {
	n := 2.0
	intn := int(n)
	//streamEncoding := "upe"

	abcdTable := generate_unique_tuples(intn)

	for _, tuple := range *abcdTable {
		xVal, yVal := tuple.BpeValue(n)
		fmt.Printf("%v --> X = %v, Y = %v\n", tuple, xVal, yVal)
	}

	//corrTypes := []string{"scc", "anderson", "dice", "jaccard", "ku2", "ochiai", "pearson", "sorensen", "ss2"}
	//
	//var wg sync.WaitGroup
	//start := time.Now()
	//for _, tp := range corrTypes {
	//	wg.Add(1)
	//	go ErrorWorker(abcdTable, n, intn, tp, streamEncoding, &wg)
	//}
	//wg.Wait()
	//elapsed := time.Since(start)
	//fmt.Printf("Took %v seconds", elapsed)
}
