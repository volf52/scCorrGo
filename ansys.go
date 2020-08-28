package main

import (
	"fmt"
)

func main() {
	n := 8.0
	intn := int(n)

	abcdTable := generate_unique_tuples(intn)
	corrTypes := []string{"scc", "anderson", "dice", "jaccard", "ku2", "ochiai", "pearson", "sorensen", "ss2"}

	for _, tp := range corrTypes {
		stringCorrTbl := readStringCorrTable(fmt.Sprintf("./n%v/%s_%v_go_rfreqs.json", intn, tp, intn))

		errTable := abcdTable.CalculateErrors(stringCorrTbl, n, GetErrorUpe)

		errTable.writeErrorTable(fmt.Sprintf("./n%v/%s_%v_rmse.json", intn, tp, intn))
	}

}
