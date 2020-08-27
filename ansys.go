package main

import (
	"fmt"
	"time"
)

func calcErr(tuple *abcdTuple, toPut *float64, n float64){
	err := tuple.RMSE(n)
	*toPut = err
}

func main() {

	abcdTable := generate_unique_tuples(4)
	errorTable := make([]float64, len(*abcdTable))

	start := time.Now()
	for i, tuple := range *abcdTable {
		calcErr(&tuple, &errorTable[i], 4.0)
	}
	dur := time.Since(start)

	fmt.Printf("RMSE: %v\n", errorTable)
	fmt.Printf("Took %v\n", dur)
}
