package main

func main() {
	n := 64.0
	intn := int(n)

	abcdTable := generate_unique_tuples(intn)

	calculateCorrelations(abcdTable, n, intn, true)
}
