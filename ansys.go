package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	n := 128.0
	intn := int(n)
	pth := fmt.Sprintf("abcd_%v.csv", intn)
	abcdTable, err := parseCsv(pth)

	if err != nil {
		log.Fatal(err)
	}

	sccTable := make(CorrTable)
	pearsonTable := make(CorrTable)
	jacTable := make(CorrTable)
	diceTable := make(CorrTable)
	sorTable := make(CorrTable)
	anderTable := make(CorrTable)
	ss2Table := make(CorrTable)
	ochTable := make(CorrTable)
	ku2Table := make(CorrTable)

	var scc, ps, jacC, diceC, sorC, anderC, ss2C, ochC, ku2C float64
	var a, b, c, d float64
	//var wg sync.WaitGroup
	//
	start := time.Now()
	//
	//wg.Add(1)
	//go sccWorker(n, &sccTable, &abcdTable, &wg)
	//
	//wg.Add(1)
	//go sccWorker(n, &sccTableTwo, &abcdTable, &wg)
	//
	//wg.Wait()
	for i, row := range abcdTable {

		a = row.a
		b = row.b
		c = row.c
		d = row.d

		scc = sccCalc(a, b, c, d, n)
		sccTable.updateTable(scc, i)

		ps = pearson(a, b, c, d)
		pearsonTable.updateTable(ps, i)

		jacC = jac(a, b, c)
		jacTable.updateTable(jacC, i)

		diceC = dice(a, b, c)
		diceTable.updateTable(diceC, i)

		sorC = sor(a, b, c)
		sorTable.updateTable(sorC, i)

		anderC = ander(a, b, c)
		anderTable.updateTable(anderC, i)

		ss2C = ss2(a, b, c)
		ss2Table.updateTable(ss2C, i)

		ochC = och(a, b, c)
		ochTable.updateTable(ochC, i)

		ku2C = ku2(a, b, c)
		ku2Table.updateTable(ku2C, i)
	}

	elapsed := time.Since(start)
	fmt.Printf("Took %v\n", elapsed)
	//fmt.Println("Length is ", len(sccTable))
	//fmt.Println("Length 2 is ", len(sccTableTwo))
	//sccServer.Write(128)
	//sccTable.writeTable("scc", 128)
	//pearsonTable.writeTable("pearson", 128)
	//jacTable.writeTable("jaccard", 128)
	//diceTable.writeTable("dice", 128)
	//sorTable.writeTable("sorensen", 128)
	//anderTable.writeTable("anderson", 128)
	//ss2Table.writeTable("ss2", 128)
	//ochTable.writeTable("ochiai", 128)
	//ku2Table.writeTable("ku2", 128)
}
