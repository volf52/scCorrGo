package main

import (
	"fmt"
	"log"
	"sync"
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
	sccServer := NewCorrServer(&sccTable)
	//pearsonTable := make(CorrTable)
	//jacTable := make(CorrTable)
	//diceTable := make(CorrTable)
	//sorTable := make(CorrTable)
	//anderTable := make(CorrTable)
	//ss2Table := make(CorrTable)
	//ochTable := make(CorrTable)
	//ku2Table := make(CorrTable)

	var wg sync.WaitGroup
	start := time.Now()

	wg.Add(1)
	go sccWorker(n, sccServer, &abcdTable, &wg)
	go sccServer.loop()


	//wg.Add(1)
	//go pearsonWorker(&pearsonTable, &abcdTable, &wg)
	//wg.Add(1)
	//go otherCorrWorker(jac, &jacTable, &abcdTable, &wg)
	//wg.Add(1)
	//go otherCorrWorker(dice, &diceTable, &abcdTable, &wg)
	//wg.Add(1)
	//go otherCorrWorker(sor, &sorTable, &abcdTable, &wg)
	//wg.Add(1)
	//go otherCorrWorker(ander, &anderTable, &abcdTable, &wg)
	//wg.Add(1)
	//go otherCorrWorker(ss2, &ss2Table, &abcdTable, &wg)
	//wg.Add(1)
	//go otherCorrWorker(och, &ochTable, &abcdTable, &wg)
	//wg.Add(1)
	//go otherCorrWorker(ku2, &ku2Table, &abcdTable, &wg)

	wg.Wait()
	close(sccServer.updateChannel)
	elapsed := time.Since(start)

	fmt.Printf("Took %v to calculate tables\n", elapsed)
	sccServer.table.writeTable("newSCC", 128)
	//start = time.Now()
	//fmt.Println("Writing tables to disk")
	//sccTable.writeTable("scc", 128)
	//pearsonTable.writeTable("pearson", 128)
	//jacTable.writeTable("jaccard", 128)
	//diceTable.writeTable("dice", 128)
	//sorTable.writeTable("sorensen", 128)
	//anderTable.writeTable("anderson", 128)
	//ss2Table.writeTable("ss2", 128)
	//ochTable.writeTable("ochiai", 128)
	//ku2Table.writeTable("ku2", 128)
	//elapsed = time.Since(start)
	//fmt.Printf("Took %v to write tables\n", elapsed)

}
