package main

import (
	"fmt"
)

func main() {
	n := 8.0
	intn := int(n)

	stringCorrTbl := readStringCorrTable(fmt.Sprintf("./n%v/scc_%v_go_rfreqs.json", intn, intn))
	abcdTable := generate_unique_tuples(intn)

	errTable := abcdTable.CalculateErrors(stringCorrTbl, n)

	errTable.writeErrorTable(fmt.Sprintf("./scc_%v_rmse.json", intn))
}
