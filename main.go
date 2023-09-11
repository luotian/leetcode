package main

import (
	"fmt"
)

func main() {
	//[[1,4,5],[1,3,4],[2,6]]
	//l1 := NewListNodeS(1, 2, 3, 4, 5)
	//val := []*ListNode{NewListNodeS(1, 4, 5), NewListNodeS(1, 3, 4), NewListNodeS(2, 6)}
	//val := []int{3, 2, 2, 3}
	v := findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"})
	//v := reverseList(l1)
	fmt.Println(v)
}

//30. 串联所有单词的子串
//给定一个字符串 s 和一个字符串数组 words。 words 中所有字符串 长度相同。
//s 中的 串联子串 是指一个包含  words 中所有字符串以任意顺序排列连接起来的子串。
//例如，如果 words = ["ab","cd","ef"]， 那么 "abcdef"， "abefcd"，"cdabef"， "cdefab"，"efabcd"， 和 "efcdab" 都是串联子串。 "acdbef" 不是串联子串，因为他不是任何 words 排列的连接。
//返回所有串联子串在 s 中的开始索引。你可以以 任意顺序 返回答案。
//
//示例 1：
//输入：s = "barfoothefoobarman", words = ["foo","bar"]
//输出：[0,9]
//解释：因为 words.length == 2 同时 words[i].length == 3，连接的子字符串的长度必须为 6。
//子串 "barfoo" 开始位置是 0。它是 words 中以 ["bar","foo"] 顺序排列的连接。
//子串 "foobar" 开始位置是 9。它是 words 中以 ["foo","bar"] 顺序排列的连接。
//输出顺序无关紧要。返回 [9,0] 也是可以的。
//
//示例 2：
//
//输入：s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
//输出：[]
//解释：因为 words.length == 4 并且 words[i].length == 4，所以串联子串的长度必须为 16。
//s 中没有子串长度为 16 并且等于 words 的任何顺序排列的连接。
//所以我们返回一个空数组。
//示例 3：
//
//输入：s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]
//输出：[6,9,12]
//解释：因为 words.length == 3 并且 words[i].length == 3，所以串联子串的长度必须为 9。
//子串 "foobarthe" 开始位置是 6。它是 words 中以 ["foo","bar","the"] 顺序排列的连接。
//子串 "barthefoo" 开始位置是 9。它是 words 中以 ["bar","the","foo"] 顺序排列的连接。
//子串 "thefoobar" 开始位置是 12。它是 words 中以 ["the","foo","bar"] 顺序排列的连接。
//
//
//提示：
//
//1 <= s.length <= 104
//1 <= words.length <= 5000
//1 <= words[i].length <= 30
//words[i] 和 s 由小写英文字母组成
func findSubstring(s string, words []string) []int {
	var ans []int
	if len(words) == 0 {
		return ans
	}

	ls := len(s)
	lw := len(words)
	lsw := len(words[0])

	for i := 0; i < lsw && i+lw*lsw <= ls; i++ {
		var diffWords = make(map[string]int)
		for j := 0; j < lw; j++ {
			diffWords[s[i+j*lsw:i+(j+1)*lsw]]++
		}

		for _, v := range words {
			diffWords[v]--
			if diffWords[v] == 0 {
				delete(diffWords, v)
			}
		}

		if len(diffWords) == 0 {
			ans = append(ans, i)
		}

		// 每个单词做一个跨度
		for start := i + lsw; start <= ls-lw*lsw; start += lsw {
			// 减去最左边
			diffWords[s[start-lsw:start]]--
			if diffWords[s[start-lsw:start]] == 0 {
				delete(diffWords, s[start-lsw:start])
			}
			// 增加最右边
			diffWords[s[start+(lw-1)*lsw:start+lw*lsw]]++
			if diffWords[s[start+(lw-1)*lsw:start+lw*lsw]] == 0 {
				delete(diffWords, s[start+(lw-1)*lsw:start+lw*lsw])
			}
			if len(diffWords) == 0 {
				ans = append(ans, start)
			}
		}
	}

	return ans
}
