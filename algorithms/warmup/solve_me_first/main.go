package main

import (
	"fmt"

	resolution "github.com/willian2s/hackerrank-golang/algorithms/warmup/solve_me_first/resolution"
)

func main() {
	var a, b, res uint32
	fmt.Println("Enter two numbers: ")
	fmt.Scanf("%v\n%v", &a, &b)
	res = resolution.SolveMeFirst(a, b)
	fmt.Println(res)
}
