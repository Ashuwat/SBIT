package main

import (
	"fmt"
	"math/rand/v2"
	pwd "test/numbers"
)

func printInfo(n pwd.NumberProps) {
	fmt.Println(n.IsPrime())
	fmt.Println(n.Factors())
	fmt.Println(n.IsEven())
}

func testing(x int) (int, error) {
	if x == 3 {
		return x, nil
	} else {
		return 0, nil
	}
}
func main() {
	// var n = pwd.Number{SomeInt: 12}
	// var s = pwd.BigNumber{LargeNum: []int{1, 2, 3, 4}}
	// printInfo(&n)
	// printInfo(&s)
	// println(s.IsPrime())
	println(rand.Int32())
	println(rand.Int32())
	println(rand.Int32())
	println(rand.Int32())
	testing(4)
}
