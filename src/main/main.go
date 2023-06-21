package main

import (
	"fmt"
	"golang/arithmetic/DFA"
)

func main() {
	DFA.Detection("我爱我家", DFA.BuildSearchTree([]string{"爱", "家"}))
	fmt.Println("hello golang")
}
