package main

import (
	"encoding/csv"
	"log"
	"os"
	"sort"
)

type Data struct {
	Roll  string
	Name  string
	Marks string
}

type Datas []Data

// func (s Datas) Len() int {
// 	return len(s)
// }

func (s Datas) Less(i, j int) bool {
	return s[i].Roll < s[j].Roll
}

func (s Datas) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Datas) Len() int {
	return len(s)
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func writeCSV(data1 [][]string, data2 [][]string, data3 [][]string, data4 [][]string) {
	file1, err := os.OpenFile("10AHindi.csv", os.O_CREATE|os.O_WRONLY, 0777)
	defer file1.Close()

	if err != nil {
		os.Exit(1)
	}
	file2, err := os.OpenFile("10BHindi.csv", os.O_CREATE|os.O_WRONLY, 0777)
	defer file2.Close()

	if err != nil {
		os.Exit(1)
	}
	file3, err := os.OpenFile("10CHindi.csv", os.O_CREATE|os.O_WRONLY, 0777)
	defer file3.Close()

	if err != nil {
		os.Exit(1)
	}

	file4, err := os.OpenFile("10DHindi.csv", os.O_CREATE|os.O_WRONLY, 0777)
	defer file4.Close()

	if err != nil {
		os.Exit(1)
	}
	csvWriter := csv.NewWriter(file1)
	csvWriter.WriteAll(data1)
	csvWriter.Flush()
	csvWriter = csv.NewWriter(file2)
	csvWriter.WriteAll(data2)
	csvWriter.Flush()
	csvWriter = csv.NewWriter(file3)
	csvWriter.WriteAll(data3)
	csvWriter.Flush()
	csvWriter = csv.NewWriter(file4)
	csvWriter.WriteAll(data4)
	csvWriter.Flush()
}

func main() {
	records := readCsvFile("name_list_10.csv")
	admNoToSecMap := make(map[string]string)
	admNoToRollMap := make(map[string]string)
	admNoToNameMap := make(map[string]string)

	for i := 1; i < len(records); i++ {
		admNoToSecMap[records[i][6]] = records[i][3]
		admNoToRollMap[records[i][6]] = records[i][4]
		admNoToNameMap[records[i][6]] = records[i][1]
	}

	marks := readCsvFile("marks_30AUG20.csv")
	// res10A := make([]Data, 0)
	// res10B := make([]Data, 0)
	// res10C := make([]Data, 0)
	// res10D := make([]Data, 0)
	// for i := 1; i < len(marks); i++ {
	// 	sec := admNoToSecMap[marks[i][20]]
	// 	roll := admNoToRollMap[marks[i][20]]
	// 	name := admNoToNameMap[marks[i][20]]
	// 	marks := marks[i][5]

	// 	d := Data{
	// 		Roll:  roll,
	// 		Name:  name,
	// 		Marks: marks,
	// 	}

	// 	if sec == "A" {
	// 		res10A = append(res10A, d)
	// 	}
	// 	if sec == "B" {
	// 		res10B = append(res10B, d)
	// 	}
	// 	if sec == "C" {
	// 		res10C = append(res10C, d)
	// 	}
	// 	if sec == "D" {
	// 		res10D = append(res10D, d)
	// 	}

	// }

	// // sort.Slice(res10A, func(i, j int) bool {
	// // 	return res10A[i].roll < res10A[j].roll
	// // })

	// // sort.Sort(res10A)

	// sort.SliceStable(res10A, func(i, j int) bool {
	// 	return res10A[i].Roll < res10A[j].Roll
	// })
	res10A := make([][]string, 0)
	res10B := make([][]string, 0)
	res10C := make([][]string, 0)
	res10D := make([][]string, 0)

	for i := 1; i < len(marks); i++ {
		sec := admNoToSecMap[marks[i][20]]
		roll := admNoToRollMap[marks[i][20]]
		name := admNoToNameMap[marks[i][20]]

		// fmt.Printf("sec: %v, roll: %v, marks: %v", sec, roll, mark)
		temp := make([]string, 0)
		temp = append(temp, roll)
		temp = append(temp, name)
		temp = append(temp, marks[i][5])

		if sec == "A" {
			res10A = append(res10A, temp)
		}
		if sec == "B" {
			res10B = append(res10B, temp)
		}
		if sec == "C" {
			res10C = append(res10C, temp)
		}
		if sec == "D" {
			res10D = append(res10D, temp)
		}
	}

	sort.SliceStable(res10A, func(i, j int) bool {
		return res10A[i][0] < res10A[j][0]
	})
	sort.SliceStable(res10B, func(i, j int) bool {
		return res10B[i][0] < res10B[j][0]
	})
	sort.SliceStable(res10C, func(i, j int) bool {
		return res10C[i][0] < res10C[j][0]
	})
	sort.SliceStable(res10D, func(i, j int) bool {
		return res10D[i][0] < res10D[j][0]
	})

	//fmt.Println(res10A)

	writeCSV(res10A, res10B, res10C, res10D)
}
