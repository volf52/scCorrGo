package main

import (
	"sync"
)

func sccWorker(n float64, table *CorrTable, abcdTable *ABCDTable, wg *sync.WaitGroup) {
	defer wg.Done()
	var scc float64
	for i, row := range *abcdTable {

		scc = sccCalc(row.a, row.b, row.c, row.d, n)
		table.updateTable(scc, i)
	}
}

func pearsonWorker(table *CorrTable, abcdTable *ABCDTable, wg *sync.WaitGroup) {
	defer wg.Done()

	var corr float64

	for i, row := range *abcdTable {
		corr = pearson(row.a, row.b, row.c, row.d)
		table.updateTable(corr, i)
	}
}

type otherFn func(a, b, c float64) float64

func otherCorrWorker(fn otherFn, table *CorrTable, abcdTable *ABCDTable, wg *sync.WaitGroup) {
	defer wg.Done()

	var corr float64

	for i, row := range *abcdTable {
		corr = fn(row.a, row.b, row.c)
		table.updateTable(corr, i)
	}
}

func (table *CorrTable) syncWrite(name string, n int, wg *sync.WaitGroup) {
	wg.Add(1)

	go func() {
		defer wg.Done()

		table.writeTable(name, n)
	}()
}
