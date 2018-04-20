package main

import (
	"fmt"
)

func decodeFile(fileName string) {
	createInitialDict()
	data := getBytesSlice(fileName)
	fmt.Printf("Data\n\n")
	// for _, vl := range data {
	// 	fmt.Printf("%d - %s\n", vl, string(vl))
	// }
	compr := dataToString(data)
	compr = removeZerosLessSignificant(compr)
	// fmt.Printf("Compr: %s\n", compr)
}
