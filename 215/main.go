// https://leetcode.com/problems/kth-largest-element-in-an-array/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func pivot(nums []int, k, left, right int) int {
	pivotIndex := left + (right-left)/2
	pivotNum := nums[pivotIndex]
	nums[pivotIndex], nums[right] = nums[right], nums[pivotIndex]
	i, j, l := left, right, len(nums)
	for j >= i {
		for i < l && nums[i] > pivotNum {
			i++
		}
		for j >= 0 && nums[j] <= pivotNum {
			j--
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	nums[i], nums[right] = nums[right], nums[i]
	kv := i + 1
	if k == kv {
		return pivotNum
	} else if k < kv {
		return pivot(nums, k, left, i-1)
	} else {
		return pivot(nums, k, i+1, right)
	}

}

func findKthLargest(nums []int, k int) int {
	return pivot(nums, k, 0, len(nums)-1)
}

func main() {
	inputFile, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer inputFile.Close()
	content, err := ioutil.ReadAll(inputFile)
	if err != nil {
		log.Fatalln(err)
	}
	strs := strings.Split(string(content), "\n")
	line1 := strings.Split(strs[0], ",")
	nums := make([]int, len(line1))
	for i := 0; i < len(line1); i++ {
		nums[i], err = strconv.Atoi(line1[i])
		if err != nil {
			log.Fatalln(err)
		}

	}
	k, _ := strconv.Atoi(strs[1])
	fmt.Println(findKthLargest(nums, k) == 8221)

	input := []int{3, 2, 3, 1, 2, 4, 5, 5, 6, 7, 7, 8, 2, 3, 1, 1, 1, 10, 11, 5, 6, 2, 4, 7, 8, 5, 6}
	fmt.Println(findKthLargest(input, 20) == 2)

	input = []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	fmt.Println(findKthLargest(input, 9) == 1)
	// //
	input = []int{3, 2, 4, 5, 6, 4}
	fmt.Println(findKthLargest(input, 2) == 5)

}
