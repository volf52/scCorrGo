package main

func main() {
	n := 256.0
	intn := int(n)

	abcdTable := generate_unique_tuples(intn)

	calculateCorrelations(abcdTable, n, intn, true)
}
