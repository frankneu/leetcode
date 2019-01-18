package main

/**
https://leetcode.com/problems/custom-sort-string/
*/

func customSortString(S string, T string) string {
	count := make(map[byte]int)
	for i := 0; i < len(T); i++ {
		count[T[i]]++
	}
	var result string
	for i := 0; i < len(S); i++ {
		if v, ok := count[S[i]]; ok {
			for v > 0 {
				result += string(S[i])
				v--
			}
			delete(count, S[i])
		}
	}
	for k, v := range count {
		for v > 0 {
			result += string(k)
			v--
		}
	}
	return result
}

func main() {
	println(customSortString("cba", "abcdef"))
}
