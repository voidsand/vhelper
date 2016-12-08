package vhelper

import (
	"math/rand"
	"time"
)

func GenerateRandomNumbers(start int, end int, count int) []int {
	// 范围检查
	if end < start || (end-start+1) < count {
		return nil
	}
	// 存放结果的slice
	nums := make([]int, 0)
	// 随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		// 生成随机数
		num := r.Intn(end-start+1) + start
		// 重复检查
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}
		if !exist {
			nums = append(nums, num)
		}
	}
	return nums
}
