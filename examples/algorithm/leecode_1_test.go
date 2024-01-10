package algorithm

import (
	"fmt"
	"slices"
	"sort"
	"testing"
)

// 1. 两数之和
// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
// 你可以按任意顺序返回答案。
func Test_1(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	sort.Slice(nums, func(i, j int) bool {
		return i > j
	})
	fmt.Println(nums)
	//twoSum2(nums, 9)
}

type Num struct {
	val   int
	index int
}

func twoSum(nums []int, target int) []int {
	filterNums := make([]Num, len(nums))
	res := make([]int, 2)
	for i, n := range nums {
		if n < target {
			filterNums = append(filterNums, Num{n, i})
		}
	}
	if len(filterNums) == 0 {
		return res
	} else {
		for i := 0; i < len(filterNums)-1; i++ {
			for j := i + 1; j < len(filterNums); j++ {
				if filterNums[i].val+filterNums[j].val == target {
					res[0] = filterNums[i].index
					res[1] = filterNums[j].index
				}
			}
		}
	}
	return res
}

func twoSum2(nums []int, target int) []int {
	res := make([]int, 2)
	for i, v := range nums {
		minus := target - v
		fmt.Println("minus = ", minus)
		other := nums[(i + 1):]
		fmt.Println("other = ", other)
		index := slices.Index(other, minus)
		fmt.Println("index = ", index)
		if index != -1 {
			res[0] = i
			res[1] = index + i + 1
			fmt.Println("res = ", res)
			return res
		}
	}
	return res
}

// 使用map存储对应的值的index 循环的时候存储进去
// 计算每次需要的值，从map中取值，取到则返回数据，取不到继续遍历
func twoSum3(nums []int, target int) []int {
	mp := make(map[int]int)
	for k, v := range nums {
		rest := target - v
		if val, ok := mp[rest]; ok {
			return []int{k, val}
		} else {
			mp[v] = k
		}
	}
	return nil
}
