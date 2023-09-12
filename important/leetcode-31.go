package important

//31. 下一个排列
//整数数组的一个 排列  就是将其所有成员以序列或线性顺序排列。
//
//例如，arr = [1,2,3] ，以下这些都可以视作 arr 的排列：[1,2,3]、[1,3,2]、[3,1,2]、[2,3,1] 。
//整数数组的 下一个排列 是指其整数的下一个字典序更大的排列。更正式地，如果数组的所有排列根据其字典顺序从小到大排列在一个容器中，那么数组的 下一个排列 就是在这个有序容器中排在它后面的那个排列。如果不存在下一个更大的排列，那么这个数组必须重排为字典序最小的排列（即，其元素按升序排列）。
//
//例如，arr = [1,2,3] 的下一个排列是 [1,3,2] 。
//类似地，arr = [2,3,1] 的下一个排列是 [3,1,2] 。
//而 arr = [3,2,1] 的下一个排列是 [1,2,3] ，因为 [3,2,1] 不存在一个字典序更大的排列。
//给你一个整数数组 nums ，找出 nums 的下一个排列。
//
//必须 原地 修改，只允许使用额外常数空间。

//题解：从右向左找最小的大数与左边的数进行交换，再把右边进行逆序
func nextPermutation(nums []int) {
	ln := len(nums)
	if len(nums) == 1 {
		return
	}

	if len(nums) == 2 {
		nums[0], nums[1] = nums[1], nums[0]
		return
	}

	var idx int = -1
	// 降序[idx, end]
	for i := ln - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			idx = i
			break
		}
	}

	if idx == -1 {
		reverse(nums)
		return
	}

	value := nums[idx-1]
	var firstGreaterIdx int
	for j := ln - 1; j >= idx; j-- {
		if nums[j] > value {
			firstGreaterIdx = j
			break
		}
	}

	nums[idx-1], nums[firstGreaterIdx] = nums[firstGreaterIdx], nums[idx-1]
	reverse(nums[idx:])
}

func reverse(nums []int) {
	l := len(nums)
	mid := len(nums) / 2
	for i := 0; i < mid; i++ {
		nums[i], nums[l-1-i] = nums[l-1-i], nums[i]
	}
}
