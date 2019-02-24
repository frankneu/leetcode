package main

import (
	"fmt"
	"math/big"
)

/**
https://www.hackerrank.com/challenges/extra-long-factorials/problem
*/

func extraLongFactorials(n int32) {
	ans := big.NewInt(1)
	for i := 2; i <= int(n); i++ {
		ans = ans.Mul(ans, big.NewInt(int64(i)))
	}
	fmt.Println(ans.String())
}

func main() {
	extraLongFactorials(100)
}
