package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	encode("video3")
}

type Node struct {
	Bits string
	Subs []Node
}

func encode(fileName string) {
	dict := createInitialDict()
	data := getBytesSlice(fileName)
	compr := encodedString(data, dict)
	// fmt.Printf("Data\n\n")
	// for _, vl := range data {
	// 	fmt.Printf("%d - %s\n", vl, string(vl))
	// }
	// fmt.Printf("Dict\n\n")
	// for idx, vl := range dict {
	// 	fmt.Printf("%s - %s\n", idx, vl)
	// }
	createEncodedFile(fileName+".cpr", compr)
}

func createInitialDict() (dict map[string]string) {
	dict = make(map[string]string)
	for i := 0; i < 256; i++ {
		dict[string(i)] = fmt.Sprintf("%0.8b", i)
		// fmt.Printf("Node: %s - %s\n", string(i), dict[string(i)])
	}
	return
}

func getBytesSlice(fileName string) (bytes []byte) {
	bytes, err := ioutil.ReadFile("instances/" + fileName)
	if err != nil {
		panic(err)
	}
	return
}

func encodedString(data []byte, dict map[string]string) (compr string) {
	pos1 := 0
	pos2 := pos1 + 1
	sizeData := len(data)
	dictSize := int64(len(dict))
	compr = ""
	stop := true
	for pos1 != sizeData-1 {
		for stop == true {
			if pos2 > sizeData {
				pos2 = sizeData
				stop = false
			} else {
				_, ok := dict[getBytesToString(data, pos1, pos2)]
				if ok == false {
					stop = false
				} else {
					pos2++
				}
			}
		}
		if pos2 > sizeData {
			pos2 = sizeData - 1
		}
		str := getBytesToString(data, pos1, pos2)
		compr += dict[getBytesToString(data, pos1, pos2-1)]
		// fmt.Printf("Data(%d, %d): %s\n", pos1, pos2, str)
		// fmt.Printf("Compr: %s\n", compr)
		// fmt.Printf("Number: %d\n", dictSize)
		numberBinary := strconv.FormatInt(int64(dictSize), 2)
		dict[str] = numberBinary
		// fmt.Printf("> dict = %s  -  %s\n", dict[string(dictSize)], numberBinary)
		// dict[string(data[pos1:pos2])]
		pos1 = pos2 - 1
		pos2 = pos1 + 1
		dictSize++
		stop = true
	}
	// fmt.Printf("Compr: %s\n", compr)
	return
}

func getBytesToString(bt []byte, pos1 int, pos2 int) (stg string) {
	stg = ""
	for i := pos1; i < pos2; i++ {
		stg += string(bt[i])
	}
	return
}

func createEncodedFile(fileName string, compress string) {
	var bt2 uint8
	var bitsBuffer = 0

	out, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("error creating file: %v", err)
		return
	}
	defer out.Close()

	bytesCreated := 0
	lastBits := len(compress) % 8

	var bytesToWrite []byte

	buf := new(bytes.Buffer)

	bytesToWrite = buf.Bytes()

	for _, vl := range compress {
		if vl == '0' {
			bt2 = bt2 << 1
		}
		if vl == '1' {
			bt2 = bt2<<1 + 1
		}
		bitsBuffer++
		if bitsBuffer == 8 {
			bytesCreated++
			bytesToWrite = append(bytesToWrite, bt2)
			bitsBuffer = 0
			bt2 = 0
		}
	}
	for i := 0; i < (8 - lastBits); i++ {
		bt2 = bt2 << 1
	}
	bytesCreated++
	bytesToWrite = append(bytesToWrite, bt2)

	err = ioutil.WriteFile(fileName, bytesToWrite, 0644)
	if err != nil {
		panic(err)
	}

}
