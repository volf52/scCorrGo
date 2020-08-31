package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type abcdTuple struct {
	a float64
	b float64
	c float64
	d float64
}

type ABCDTable []abcdTuple
type CorrTable map[float64][]int
type StringCorrTable map[string][]int
type ErrorTable map[string]float64

func MakeCorrTable() *CorrTable {
	table := make(CorrTable)

	return &table
}

func getNumOfUniqueTuples(n int) int {
	cs := 0
	ts := 0

	for i := 0; i < n+2; i++ {
		cs += i
		ts += cs
	}

	return ts
}

func generateUniqueTuples(N int) *ABCDTable {
	arr := make(ABCDTable, getNumOfUniqueTuples(N))

	idx := 0
	n := float64(N)
	i := n
	var j, k float64

	for i > -1 {
		j = n - i
		for j > -1 {
			k = n - i - j
			for k > -1 {
				arr[idx].a = i
				arr[idx].b = j
				arr[idx].c = k
				arr[idx].d = n - i - j - k

				k--
				idx++
			}
			j--
		}
		i--
	}

	return &arr
}

func readCsv(filepth string) ([][]string, error) {
	csvfile, err := os.Open(filepth)

	if err != nil {
		return nil, err
	}

	defer csvfile.Close()

	reader := csv.NewReader(csvfile)
	fields, err := reader.ReadAll()

	return fields, nil
}

func parseFloat64(s string) float64 {
	t, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatal("Error parsing ", s, "exiting")
	}

	return t
}

func parseCsv(pth string) (ABCDTable, error) {
	records, err := readCsv(pth)

	if err != nil {
		return nil, err
	}

	arr := make(ABCDTable, len(records))

	for i := 0; i < len(records); i++ {
		arr[i].a = parseFloat64(records[i][0])
		arr[i].b = parseFloat64(records[i][1])
		arr[i].c = parseFloat64(records[i][2])
		arr[i].d = parseFloat64(records[i][3])
	}

	return arr, nil
}

func (table CorrTable) updateTable(key float64, val int) {
	table[key] = append(table[key], val)
}

func (table StringCorrTable) updateTable(key string, val []int) {
	table[key] = append(table[key], val...)
}

func (table *StringCorrTable) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	buf.Write([]byte{'{'})

	var (
		b    []byte
		err  error
		i    = 0
		last = len(*table)
	)
	for k, v := range *table {
		b, err = json.Marshal(v)
		if err != nil {
			return nil, err
		}

		buf.WriteString(fmt.Sprintf("%q:", k))
		buf.Write(b)

		i++
		if i < last {
			buf.Write([]byte{','})
		}
	}

	buf.Write([]byte{'}'})
	return buf.Bytes(), nil
}

func (table *CorrTable) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	buf.Write([]byte{'{'})

	var (
		b    []byte
		err  error
		i    = 0
		last = len(*table)
	)
	for k, v := range *table {
		b, err = json.Marshal(v)
		if err != nil {
			return nil, err
		}

		buf.WriteString(fmt.Sprintf("%q:", fmt.Sprintf("%0.5f", k)))
		buf.Write(b)

		i++
		if i < last {
			buf.Write([]byte{','})
		}
	}

	buf.Write([]byte{'}'})
	return buf.Bytes(), nil
}

func (table *CorrTable) writeTable(name string, n int) {
	pth := fmt.Sprintf("n%v/%v_%v_go_rfreqs.json", n, name, n)

	var tmpStr string

	stringTable := make(StringCorrTable, len(*table))
	for k, v := range *table {
		tmpStr = fmt.Sprintf("%.5f", k)
		stringTable.updateTable(tmpStr, v)
	}

	b, err := json.MarshalIndent(stringTable, "", "\t")

	if err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile(pth, b, 0644)
}

func readStringCorrTable(pth string) *StringCorrTable {
	data, err := ioutil.ReadFile(pth)

	if err != nil {
		panic(err)
	}

	table := make(StringCorrTable)

	err = json.Unmarshal(data, &table)
	if err != nil {
		panic(err)
	}

	return &table
}

func (errTable *ErrorTable) writeErrorTable(pth string) {
	s, err := json.MarshalIndent(errTable, "", "\t")
	if err != nil {
		panic(err)
	}

	ioutil.WriteFile(pth, s, 0644)
}
