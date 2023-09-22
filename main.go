package main

import (
	"fmt"
	"sort"
)

func main() {
	//[[1,4,5],[1,3,4],[2,6]]
	//l1 := NewListNodeS(1, 2, 3, 4, 5)
	//val := []*ListNode{NewListNodeS(1, 4, 5), NewListNodeS(1, 3, 4), NewListNodeS(2, 6)}
	val := []int{10, 1, 2, 7, 6, 1, 5}

	v := combinationSum2(val, 8)
	//v := reverseList(l1)
	fmt.Println(v)
}

//39. 组合综合
//给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合
//，并以列表形式返回。你可以按 任意顺序 返回这些组合。
//candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。
//对于给定的输入，保证和为 target 的不同组合数少于 150 个。

//提示：
//
//1 <= candidates.length <= 30
//2 <= candidates[i] <= 40
//candidates 的所有元素 互不相同
//1 <= target <= 40

var result [][]int

func combinationSum2(candidates []int, target int) [][]int {
	result = make([][]int, 0)
	sort.Ints(candidates)
	path := make([]int, 0)
	backtracking(path, candidates, target)
	return result
}

func backtracking(path []int, candidates []int, target int) {
	if target == 0 {
		tmp := make([]int, len(path))
		copy(tmp, path)
		result = append(result, tmp)
		return
	}

	for i := 0; i < len(candidates) && candidates[i] <= target; i++ {
		if i > 0 && candidates[i] == candidates[i-1] {
			continue
		}
		path = append(path, candidates[i])
		target -= candidates[i]
		backtracking(path, candidates[i+1:], target)
		path = path[:len(path)-1]
		target += candidates[i]
	}
}
