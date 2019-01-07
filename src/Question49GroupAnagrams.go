package main

/**
https://leetcode.com/problems/group-anagrams/
*/
func groupAnagrams(strs []string) [][]string {
	var result [][]string
	value := make(map[string][]string)
	for i := 0; i < len(strs); i++ {
		anas := make([]int, 26)
		for j := 0; j < len(strs[i]); j++ {
			anas[strs[i][j]-'a']++
		}
		var str string
		for j := 0; j < len(anas); j++ {
			str += string(anas[j])
		}
		if v, ok := value[str]; ok {
			v = append(v, strs[i])
			value[str] = v
		} else {
			value[str] = []string{strs[i]}
		}
	}
	for _, v := range value {
		result = append(result, v)
	}
	return result
}

func main() {
	result := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[i]); j++ {
			print(result[i][j], " ")
		}
		println()
	}
}
