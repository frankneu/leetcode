package main

/**
https://www.hackerrank.com/challenges/sock-merchant/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=warmup
*/

func sockMerchant(n int32, ar []int32) int32 {
	set := make(map[int32]int32)
	for i := 0; i < len(ar); i++ {
		set[ar[i]]++
	}
	var result int32
	for _, v := range set {
		result += v / 2
	}
	return result
}
