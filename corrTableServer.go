package main

import (
	"sync"
)

func sccWorker(n float64, table *CorrTable, abcdTable *[]dfRow, wg *sync.WaitGroup){
	defer wg.Done()
	var scc float64

	for i, row := range *abcdTable {
		scc = sccCalc(row.a, row.b, row.c, row.d, n)
		table.updateTable(scc, i)
	}
}




//type Transactor func(table CorrTable)
//
//type TransactionCommand struct {
//	Transactor Transactor
//	Done       chan bool
//}
//
//type CorrServer struct {
//	commands chan TransactionCommand
//	table    CorrTable
//	name string
//}
//
//func (s *CorrServer) Transact(transactor Transactor) {
//	command := TransactionCommand{
//		Transactor: transactor,
//		Done:       make(chan bool),
//	}
//
//	s.commands <- command
//	<-command.Done
//}
//
//func NewCorrServer(name string) *CorrServer{
//	server := &CorrServer{
//		commands: make(chan TransactionCommand),
//		table:    make(CorrTable),
//		name: name,
//	}
//
//	go server.loop()
//	return server
//}
//
//func (s *CorrServer) loop(){
//	for update := range s.commands{
//		update.Transactor(s.table)
//		update.Done <- true
//	}
//}
//
//func (s *CorrServer) Update(key float64, val int){
//	s.Transact(func(table CorrTable){
//		table.updateTable(key, val)
//	})
//}
//
//func (s *CorrServer) Write(n int){
//	s.table.writeTable(s.name, n)
//}