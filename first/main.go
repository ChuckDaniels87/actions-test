package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1]
	fmt.Printf("first action says: %s\n", arg)
	fmt.Printf("::set-output name=out::%s", arg + "-first")
}
