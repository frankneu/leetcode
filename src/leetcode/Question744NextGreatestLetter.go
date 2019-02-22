package main

/**
https://leetcode.com/problems/find-smallest-letter-greater-than-target/
*/
func nextGreatestLetter(letters []byte, target byte) byte {
	i := 0
	j := len(letters) - 1
	m := 0
	for i <= j {
		m = (i + j) / 2
		if letters[m] == target && letters[(m+1)%len(letters)] != target {
			return letters[(m+1)%len(letters)]
		} else if letters[m] == target && letters[(m+1)%len(letters)] == target {
			i = m + 1
		} else if letters[m] < target {
			i = m + 1
		} else {
			j = m - 1
		}
	}
	if target > letters[m] {
		m += 1
	}
	return letters[m%len(letters)]
}

func main() {
	letters := []byte{'e', 'e', 'e', 'e', 'e', 'e', 'e', 'e', 'e', 'f', 'f'}
	println(string(nextGreatestLetter(letters, 'e')))
}
