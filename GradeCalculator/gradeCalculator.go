package gradeCalc

import (
	"fmt"

	"github.com/dustin/go-humanize"
)

func averageCalculator(datas *[]data, n uint8) float32 {
	var sum float32
	for _, d := range *datas {
		sum += d.grade
	}
	return sum / float32(n)
}

type data struct {
	name  string
	grade float32
}

func DataCollector() {

	var name string
	var numberOfSubjects uint8

	fmt.Printf("enter your name: ")
	fmt.Scan(&name)
	fmt.Printf("enter the number of subjects: ")
	fmt.Scan(&numberOfSubjects)

	var datas []data = make([]data, numberOfSubjects)

	for i := 0; i < int(numberOfSubjects); i++ {
		fmt.Printf("enter the %s subject name: ", humanize.Ordinal(i+1))
		fmt.Scan(&datas[i].name)
		fmt.Printf("enter the grade for %s: ", datas[i].name)
		fmt.Scan(&datas[i].grade)
	}

	DataPrinter(&datas, numberOfSubjects, name)
}

func DataPrinter(datas *[]data, n uint8, name string) {
	fmt.Println()
	fmt.Printf("Hello, %s here are your grades: \n", name)
	for _, data := range *datas {
		fmt.Printf("	%s  :  %.2f\n", data.name, data.grade)
	}

	fmt.Printf("	Average  :  %.2f\n", averageCalculator(datas, n))
}

func StartProgram() {
	DataCollector()
}
