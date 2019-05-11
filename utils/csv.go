package utils

import (
	"encoding/csv"
	"fmt"
	"os"
)

func SaveToCSV(data [][]string, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// UTF-8 BOM
	f.WriteString("\xEF\xBB\xBF")
	w := csv.NewWriter(f)
	w.WriteAll(data)
	fmt.Println("Done!")

}

// func main() {
//     f, err := os.Create("./test.csv")
//     if err != nil {
//         panic(err)
//     }
//     defer f.Close()

//     // UTF-8 BOM
//     f.WriteString("\xEF\xBB\xBF")
//     w := csv.NewWriter(f)
//     data := [][]string{
//         {"1", "2", "3"},
//         {"4", "5", "6"},
//         {"7", "8", "9"},
//     }
//     w.WriteAll(data)
//     fmt.Println("Done!")
// }
