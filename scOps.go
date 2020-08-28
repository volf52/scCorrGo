package main

import (
	"fmt"
	"math"
)

type ErrorFunc func(tuple *abcdTuple, n float64) float64

func (tuple *abcdTuple) GetOnes() (float64, float64) {
	xOnes := tuple.a + tuple.b
	yOnes := tuple.a + tuple.c

	return xOnes, yOnes
}

func UpeValue(nOnes float64, n float64) float64 {
	return nOnes / n
}

func (tuple *abcdTuple) UpeValue(n float64) (float64, float64) {
	xOnes := tuple.a + tuple.b
	yOnes := tuple.a + tuple.c

	xVal := xOnes / n
	yVal := yOnes / n
	return xVal, yVal
}

//TODO: Complete
func (tuple *abcdTuple) BpeValue(n float64) (float64, float64) {
	xOnes := tuple.a + tuple.b
	yOnes := tuple.a + tuple.c

	xVal := (2*xOnes - n) / n
	yVal := (2*yOnes - n) / n
	return xVal, yVal
}

func (tuple *abcdTuple) And() float64 {
	return tuple.a
}

func (tuple *abcdTuple) Or() float64 {
	return tuple.a + tuple.b + tuple.c
}

func (tuple *abcdTuple) Xor() float64 {
	return tuple.b + tuple.c
}

func (tuple *abcdTuple) Nand() float64 {
	return tuple.b + tuple.c + tuple.d
}

func (tuple *abcdTuple) Nor() float64 {
	return tuple.d
}

func (tuple *abcdTuple) Xnor() float64 {
	return tuple.a + tuple.d
}

func GetErrorUpe(tuple *abcdTuple, n float64) float64 {
	xVal, yVal := tuple.UpeValue(n)
	andResult := tuple.a / n

	return (xVal * yVal) - andResult
}

//TODO: Complete
func GetErrorBpe(tuple *abcdTuple, n float64) float64 {
	xVal, yVal := tuple.UpeValue(n)
	andResult := tuple.a / n

	return (xVal * yVal) - andResult
}

func rmse(idxList *[]int, abcdTable *ABCDTable, n float64, fn ErrorFunc) float64 {
	var errVal float64

	squaredSum := 0.0

	for _, idx := range *idxList {
		errVal = fn(&((*abcdTable)[idx]), n)
		squaredSum += errVal * errVal
	}

	squaredSum /= float64(len(*idxList))

	return math.Sqrt(squaredSum)
}

func (abcdTable *ABCDTable) CalculateErrors(corrTable *StringCorrTable, n float64, tp string) *ErrorTable {
	var fn ErrorFunc
	switch tp {
	case "upe":
		fn = GetErrorUpe
	case "bpe":
		fn = GetErrorBpe
	default:
		panic(fmt.Sprintf("%s is not an error function", tp))
	}
	errorTable := make(ErrorTable, len(*corrTable))

	for corr, idxList := range *corrTable {
		errorTable[corr] = rmse(&idxList, abcdTable, n, fn)
	}

	return &errorTable
}
