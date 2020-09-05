// https://leetcode.com/problems/shuffle-an-array/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Solution struct {
	seeds []int64
	nums  []int
}

func Constructor(nums []int) Solution {
	return Solution{seeds: make([]int64, 0), nums: nums}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	for i := len(this.seeds) - 1; i >= 0; i-- {
		this.reverseShuffle(this.seeds[i])
	}
	this.seeds = this.seeds[:0]
	return this.nums
}

func (this *Solution) reverseShuffle(seed int64) {
	if len(this.nums) < 2 {
		return
	}
	randGenerator := rand.New(rand.NewSource(seed))
	indices := make([]int, len(this.nums)-1)
	j, i := len(this.nums)-1, len(indices)-1
	for j > 0 {
		indices[i] = randGenerator.Intn(j + 1)
		j--
		i--
	}

	for i, j := 1, 0; i < len(this.nums) && j < len(indices); i, j = i+1, j+1 {
		this.nums[i], this.nums[indices[j]] = this.nums[indices[j]], this.nums[i]
	}
}

func (this *Solution) shuffle() []int {

	seed := time.Now().UnixNano()
	this.seeds = append(this.seeds, seed)
	randGenerator := rand.New(rand.NewSource(seed))
	j := len(this.nums) - 1
	for j > 0 {
		i := randGenerator.Intn(j + 1)
		this.nums[i], this.nums[j] = this.nums[j], this.nums[i]
		j--
	}
	return this.nums
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	return this.shuffle()
}

func main() {
	sol := Constructor([]int{1, 2})
	fmt.Println(sol.Shuffle())
	fmt.Println(sol.Shuffle())
	fmt.Println(sol.Shuffle())
	fmt.Println(sol.Reset())
	fmt.Println(sol.Reset())
	fmt.Println(sol.Reset())

}
