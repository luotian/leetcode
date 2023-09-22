package important

import (
	"fmt"
	"sort"
)

//39. 组合综合
// 经典回溯递归
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

//func combinationSum(candidates []int, target int) [][]int {
//	result = make([][]int, 0)
//	sort.Ints(candidates)
//
//	for i := 0; i < len(candidates); i++ {
//		if candidates[i] > target {
//			break
//		} else if candidates[i] == target {
//			result = append(result, []int{target})
//			break
//		}
//
//		backtracking([]int{candidates[i]}, candidates[i:], target)
//	}
//
//	return result
//}

func backtracking(cur []int, candidates []int, target int) bool {
	//fmt.Printf("%+v\n", cur)
	v := calVal(cur)
	if v == target {
		result = append(result, cur)
		return true
	}

	if v > target {
		return true
	}

	for i := 0; i < len(candidates); i++ {
		cur = append(cur, candidates[i])
		b := backtracking(cur, candidates[i:], target)
		if b {
			return false
		}
		cur = append([]int{}, cur[:len(cur)-1]...)
	}

	return false
}

func calVal(v []int) int {
	var ret int
	for _, vv := range v {
		ret += vv
	}

	return ret
}

var (
	res  [][]int
	path []int
)

func combinationSum(candidates []int, target int) [][]int {
	res, path = make([][]int, 0), make([]int, 0, len(candidates))
	sort.Ints(candidates) // 排序，为剪枝做准备
	dfs(candidates, 0, target)
	return res
}

func dfs(candidates []int, start int, target int) {
	fmt.Printf("%+v\n", path)
	if target == 0 { // target 不断减小，如果为0说明达到了目标值
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		return
	}
	for i := start; i < len(candidates); i++ {
		if candidates[i] > target { // 剪枝，提前返回
			break
		}
		path = append(path, candidates[i])
		dfs(candidates, i, target-candidates[i])
		path = path[:len(path)-1]
	}
}
