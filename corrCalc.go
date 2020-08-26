package main

import (
	"fmt"
	"log"
	"math"
	"sync"
	"time"
)

func calculateCorrelations(abcdTable *[]dfRow, n float64, intn int, writeTables bool) {
	sccTable := MakeCorrTable()
	pearsonTable := MakeCorrTable()
	jacTable := MakeCorrTable()
	diceTable := MakeCorrTable()
	sorTable := MakeCorrTable()
	anderTable := MakeCorrTable()
	ss2Table := MakeCorrTable()
	ochTable := MakeCorrTable()
	ku2Table := MakeCorrTable()

	var wg sync.WaitGroup
	log.Println("Starting...")
	start := time.Now()

	wg.Add(1)
	go sccWorker(n, sccTable, abcdTable, &wg)
	wg.Add(1)
	go pearsonWorker(pearsonTable, abcdTable, &wg)
	wg.Add(1)
	go otherCorrWorker(jac, jacTable, abcdTable, &wg)
	wg.Add(1)
	go otherCorrWorker(dice, diceTable, abcdTable, &wg)
	wg.Add(1)
	go otherCorrWorker(sor, sorTable, abcdTable, &wg)
	wg.Add(1)
	go otherCorrWorker(ander, anderTable, abcdTable, &wg)
	wg.Add(1)
	go otherCorrWorker(ss2, ss2Table, abcdTable, &wg)
	wg.Add(1)
	go otherCorrWorker(och, ochTable, abcdTable, &wg)
	wg.Add(1)
	go otherCorrWorker(ku2, ku2Table, abcdTable, &wg)

	wg.Wait()
	elapsed := time.Since(start)

	fmt.Printf("Took %v to calculate tables\n", elapsed)

	if writeTables {
		var writeWg sync.WaitGroup
		start = time.Now()

		fmt.Println("Writing tables to disk")
		sccTable.syncWrite("scc", intn, &writeWg)
		pearsonTable.syncWrite("pearson", intn, &writeWg)
		jacTable.syncWrite("jaccard", intn, &writeWg)
		diceTable.syncWrite("dice", intn, &writeWg)
		sorTable.syncWrite("sorensen", intn, &writeWg)
		anderTable.syncWrite("anderson", intn, &writeWg)
		ss2Table.syncWrite("ss2", intn, &writeWg)
		ochTable.syncWrite("ochiai", intn, &writeWg)
		ku2Table.syncWrite("ku2", intn, &writeWg)

		writeWg.Wait()
		elapsed = time.Since(start)

		fmt.Printf("Took %v to write tables\n", elapsed)
	}
	log.Println("Done!!!")
}

func sccCalc(a, b, c, d, n float64) float64 {
	numer := a * d
	numer -= b * c

	if numer == 0 {
		return numer
	}

	var denom float64

	apb := a + b
	apc := a + c
	apbIntoApc := apb * apc

	if numer > 0 {
		denom = n * math.Min(apb, apc)
		denom -= apbIntoApc
	} else
	{
		denom = apbIntoApc
		denom -= n * math.Max(a-d, 0)
	}

	numer /= denom

	return numer
}

func pearson(a, b, c, d float64) float64 {
	numer := a * d
	numer -= b * c

	if numer == 0 {
		return 0
	}

	denom := a + b
	denom *= a + c
	denom *= b + d
	denom *= c + d
	denom = math.Sqrt(denom)

	numer /= denom

	return numer
}

func jac(a, b, c float64) float64 {
	result := a
	if result == 0 {
		return 0
	}

	result /= a + b + c

	return result
}

func dice(a, b, c float64) float64 {
	result := a
	if result == 0 {
		return 0
	}

	result *= 2
	result /= result + b + c

	return result
}

func sor(a, b, c float64) float64 {
	result := a
	if result == 0 {
		return 0
	}

	result *= 4
	result /= result + b + c

	return result
}

func ander(a, b, c float64) float64 {
	result := a
	if result == 0 {
		return 0
	}

	result *= 8
	result /= result + b + c

	return result
}

func ss2(a, b, c float64) float64 {
	result := a
	if result == 0 {
		return 0
	}

	result /= result + 2*(b+c)

	return result
}

func och(a, b, c float64) float64 {
	result := a
	if result == 0 {
		return 0
	}

	result /= math.Sqrt(a+b) * math.Sqrt(a+c)

	return result
}

func ku2(a, b, c float64) float64 {
	if a == 0 {
		return 0
	}

	result := (a / (a + b)) + (a / (a + c))
	result /= 2

	return result
}
