package main

func searchRange(nums []int, target int) []int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		m := (l + r) / 2
		if target < nums[m] {
			r = m - 1
		} else if target > nums[m] {
			l = m + 1
		} else {
			l1 := l
			r1 := m
			l2 := m
			r2 := r
			m1 := (l1 + r1) / 2
			m2 := (l2 + r2) / 2
			for l1 <= r1 {
				m1 = (l1 + r1) / 2
				if nums[m1] != target {
					l1 = m1 + 1
				} else {
					r1 = m1 - 1
				}
			}
			for l2 <= r2 {
				m2 = (l2 + r2) / 2
				if nums[m2] != target {
					r2 = m2 - 1
				} else {
					l2 = m2 + 1
				}
			}
			if nums[m1] != target {
				m1 = m1 + 1
			}
			if nums[m2] != target {
				m2 = m2 - 1
			}
			return []int{m1, m2}
		}
	}
	return []int{-1, -1}

}

func main() {

	nums := []int{8, 8}
	result := searchRange(nums, 8)
	println(result[0])
	println(result[1])
}
