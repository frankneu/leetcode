package main

/**
https://leetcode.com/problems/jewels-and-stones/
*/
func numJewelsInStones(J string, S string) int {
	set := make(map[uint8]int)
	for i := 0; i < len(J); i++ {
		set[J[i]] = 1
	}
	count := 0
	for i := 0; i < len(S); i++ {
		count += set[S[i]]
	}
	return count
}

func main() {
	println(numJewelsInStones("aA", "aAAbbb"))
}
