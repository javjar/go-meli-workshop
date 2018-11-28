package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/javjar/go-meli-workshop/class2/third/node"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("No input given")
		return
	}

	fmt.Println("Arguments:", os.Args[1:])

	root := os.Args[1]
	atoi, _ := strconv.Atoi(root)
	n := &node.Node{Value: atoi}

	if len(os.Args) >= 3 {
		args := os.Args[2:]
		for _, v := range args {
			atoi, _ := strconv.Atoi(v)
			n = node.Add(n, atoi)
		}
	}

	node.PrintAll(n)
}
