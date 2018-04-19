package main

import "fmt"

func main() {
	createInitialDict()
}

type Node struct {
	Bits string
	Subs []*Node
}

func createInitialDict() {
	sliceNodes := make([]*Node, 256)
	for i := range sliceNodes {
		sliceNodes[i] = new(Node)
	}
	root := Node{"", sliceNodes}
	for i := 0; i < 256; i++ {
		fmt.Printf("num: %d - %0.8b\n", i, i)
	}
}
