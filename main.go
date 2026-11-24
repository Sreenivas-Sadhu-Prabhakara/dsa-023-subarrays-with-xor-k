package main

import "fmt"

// CountSubarraysXor returns how many contiguous subarrays of nums have
// XOR equal to k, via prefix-XOR frequencies.
func CountSubarraysXor(nums []int, k int) int {
	freq := map[int]int{0: 1} // prefix XOR -> how often seen (empty prefix included)
	x, count := 0, 0
	for _, v := range nums {
		x ^= v
		count += freq[x^k]
		freq[x]++
	}
	return count
}

func main() {
	fmt.Println(CountSubarraysXor([]int{4, 2, 2, 6, 4}, 6))
}
