package main

func main() {
	print(reverse(1534236469))
}

const maxUint = ^uint32(0)
const maxInt = int(maxUint >> 1)
const minInt = -maxInt - 1

func reverse(x int) int {
	var result = 0
	for x != 0 {
		result *= 10
		value := x % 10
		result += value
		x /= 10
	}
	if result > maxInt || result < minInt {
		return 0
	} else {
		return result
	}
}
