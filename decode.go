package main

import (
	"fmt"
)

func decodeFile(fileName string) {
	compr := ""

	createInitialDict()
	data := getBytesSlice(fileName)
	fmt.Printf("Data\n\n")
	for _, vl := range data {
		fmt.Printf("%d - %s\n", vl, string(vl))
		compr += fmt.Sprintf("%0.8b", vl)
	}
	compr = removeZerosFromEnd(compr)
	fmt.Printf("Compr: %s\n", compr)
}
