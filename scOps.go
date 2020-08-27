package main

import "math"

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

	return xOnes / n, yOnes / n
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

func (tuple *abcdTuple) RMSE(n float64) float64{
	xVal, yVal := tuple.UpeValue(n)
	andRes := tuple.And() / n

	errVal := (xVal * yVal) - andRes
	errVal *= errVal

	errVal /= n
	errVal = math.Sqrt(errVal)

	return errVal
}

func (table *ABCDTable) RMSE(n float64) float64{
	squaredErr := 0.0

	for _, tuple := range *table{
		xVal, yVal := tuple.UpeValue(n)
		andRes := tuple.And()

		errVal := (xVal * yVal) - andRes
		squaredErr += errVal * errVal
	}

	return math.Sqrt(squaredErr / n)
}