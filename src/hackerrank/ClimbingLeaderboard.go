package main

/**
https://www.hackerrank.com/challenges/climbing-the-leaderboard/problem
todo
*/

func climbingLeaderboard(scores []int32, alice []int32) []int32 {
	record := make([]int32, 0)
	for i := 0; i < len(scores); i++ {
		if len(record) == 0 || record[len(record)-1] != scores[i] {
			record = append(record, scores[i])
		}
	}
	var ans []int32
	l := 0
	r := len(record) - 1
	m := r
	for i := 0; i < len(alice); i++ {
		l = 0
		r = m
		for l <= r {
			m = (l + r) / 2
			if record[m] == alice[i] {
				ans = append(ans, int32(m+1))
				break
			} else if record[m] > alice[i] {
				l = m + 1
			} else {
				r = m - 1
			}
		}
		if record[m] > alice[i] {
			ans = append(ans, int32(m+2))
			record = insert(record, alice[i], m+1)
		} else if record[m] < alice[i] {
			ans = append(ans, int32(m+1))
			record = insert(record, alice[i], m)
		}
	}
	return ans
}

func insert(record []int32, value int32, index int) []int32 {
	result := make([]int32, index)
	copy(result, record[:index])
	result = append(result, value)
	for i := index; i < len(record); i++ {
		result = append(result, record[i])
	}
	return result
}

func main() {
	climbingLeaderboard([]int32{100, 100, 50, 40, 40, 20, 10}, []int32{5, 25, 50, 120})
}
