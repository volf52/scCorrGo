package main

import (
	"sync"
)

func sccWorker(n float64, server *CorrServer, abcdTable *[]dfRow, wg *sync.WaitGroup) {
	defer wg.Done()
	//var scc float64
	var innerWg sync.WaitGroup
	for i, row := range *abcdTable {
		innerWg.Add(1)
		go func(trow dfRow, idx int) {
			defer innerWg.Done()
			scc := sccCalc(trow.a, trow.b, trow.c, trow.d, n)
			server.updateChannel <- UpdateRequest{
				k: scc,
				v: idx,
			}
		}(row, i)
		//scc = sccCalc(row.a, row.b, row.c, row.d, n)
		//table.updateTable(scc, i)
	}

	innerWg.Wait()
}

func pearsonWorker(table *CorrTable, abcdTable *[]dfRow, wg *sync.WaitGroup) {
	defer wg.Done()

	var corr float64

	for i, row := range *abcdTable {
		corr = pearson(row.a, row.b, row.c, row.d)
		table.updateTable(corr, i)
	}
}

type otherFn func(a, b, c float64) float64

func otherCorrWorker(fn otherFn, table *CorrTable, abcdTable *[]dfRow, wg *sync.WaitGroup) {
	defer wg.Done()

	var corr float64

	for i, row := range *abcdTable {
		corr = fn(row.a, row.b, row.c)
		table.updateTable(corr, i)
	}
}

type UpdateRequest struct {
	k float64
	v int
}
type CorrServer struct {
	table         *CorrTable
	updateChannel chan UpdateRequest
}

func NewCorrServer(tbl *CorrTable) *CorrServer {
	server := &CorrServer{
		table:         tbl,
		updateChannel: make(chan UpdateRequest),
	}

	return server
}

func (s *CorrServer) loop(){
	for update := range s.updateChannel{
		s.table.updateTable(update.k, update.v)
	}
}