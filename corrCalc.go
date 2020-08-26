package main

import "math"

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
