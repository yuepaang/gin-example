package main

import (
	"github.com/ypeng7/data-microservices/utils"
)

func main() {

	filename := "./test.csv"
	data := [][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}
	utils.SaveToCSV(data, filename)
}
