package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

type Node struct {
	Bits string
	Subs []Node
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
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
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

// Decoding...

func dataToString(data []byte) (compr string) {
	for _, vl := range data {
		compr += fmt.Sprintf("%0.8b", vl)
	}
	return
}

func removeZerosLessSignificant(compr string) string {
	for _ = range compr {
		if compr[len(compr)-1] == '0' {
			compr = compr[:len(compr)-2]
		} else {
			break
		}
	}
	return compr
}
