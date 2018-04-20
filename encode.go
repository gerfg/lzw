package main

import (
	"fmt"
	"strconv"
)

func encode(fileName string) {
	dict := createInitialDict()
	data := getBytesSlice(fileName)
	compr := encodedString(data, dict)
	// fmt.Printf("Data\n\n")
	// for _, vl := range data {
	// 	fmt.Printf("%d - %s\n", vl, string(vl))
	// }
	fmt.Printf("Dict\n\n")
	// for idx, vl := range dict {
	// 	fmt.Printf("%s - %s\n", idx, vl)
	// }
	fmt.Println("============================")
	// fmt.Printf("Bits: %s\n", compr)
	fileName = fileName[10:]
	createEncodedFile("decoded/"+fileName+".cpr", compr)
}

func encodedString(data []byte, dict map[string]string) (compr string) {
	pos1 := 0
	pos2 := pos1 + 1
	sizeData := len(data)
	dictSize := int64(len(dict))
	compr = ""
	stop := true
	for pos1 != sizeData-1 {
		fmt.Println("")
		for stop == true {
			if pos2 > sizeData {
				pos2 = sizeData
				stop = false
			} else {
				_, ok := dict[getBytesToString(data, pos1, pos2)]
				fmt.Printf("Matching: %s ok: %b - %s -pos1: %d pos2: %d  sizeData: %d\n", getBytesToString(data, pos1, pos2), ok, dict[getBytesToString(data, pos1, pos2)], pos1, pos2, sizeData)
				if ok == false {
					stop = false
				} else {
					pos2++
				}
			}
		}
		fmt.Println("============================\n ")
		if pos2 > sizeData {
			pos2 = sizeData - 1
		}
		str := getBytesToString(data, pos1, pos2)
		compr += dict[getBytesToString(data, pos1, pos2-1)]
		fmt.Printf("> Add to compr: %s - %s\n", getBytesToString(data, pos1, pos2-1), dict[getBytesToString(data, pos1, pos2-1)])
		// fmt.Printf("Compr: %s\n", compr)
		// fmt.Printf("Number: %d\n", dictSize)
		numberBinary := strconv.FormatInt(int64(dictSize), 2)
		dict[str] = numberBinary
		// fmt.Printf("Data(%d, %d): %s\n", pos1, pos2, str)
		// fmt.Printf("dict = %s  -  %s\n", getBytesToString(data, pos1, pos2), numberBinary)
		// dict[string(data[pos1:pos2])]
		pos1 = pos2 - 1
		pos2 = pos1 + 1
		dictSize++
		stop = true
	}
	fmt.Printf("Compr: %s\n", compr)
	return
}
