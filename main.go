package main

import (
	"fmt"
)

func main() {
	//[[1,4,5],[1,3,4],[2,6]]
	//l1 := NewListNodeS(1, 2, 3, 4, 5)
	//val := []*ListNode{NewListNodeS(1, 4, 5), NewListNodeS(1, 3, 4), NewListNodeS(2, 6)}
	//val := []int{3, 2, 2, 3}
	v := findSubstring("wordgoodgoodgoodbestwordddddddddddzzdafsdfadfafsd", []string{"word", "good", "best", "good"})
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
func findSubstring(s string, words []string) (ans []int) {
	// 总长度ls
	// m个单词
	// 每个单词的长度为n
	strLen, wordsLen, wordLen := len(s), len(words), len(words[0])

	var differWords = make(map[string]int)

	// 找到合适的划分区间
	// 因为单词的长度为n, 所以n次为一轮
	// 所有的单词都用上的最小长度为m*wordLen, 所以需要满足 i+wordsLen*wordLen <= strLen
	var c1 int
	var d1, d2, d3 int
	for i := 0; i < wordLen && i+wordsLen*wordLen <= strLen; i++ {
		c1++
		differWords = make(map[string]int)
		// 统计这一段初始字母中, 所有单词的计数
		differ := differWords
		for j := 0; j < wordsLen; j++ {
			d1++
			differ[s[i+j*wordLen:i+(j+1)*wordLen]]++
		}
		// 去掉目标单词对应的计数
		for _, word := range words {
			d2++
			differ[word]--
			if differ[word] == 0 {
				delete(differ, word)
			}
		}
		if len(differ) == 0 {
			ans = append(ans, i)
		}
		for start := i + wordLen; start < strLen-wordsLen*wordLen+1; start += wordLen {
			d3++
			// 右端点加入一个单词
			word := s[start+(wordsLen-1)*wordLen : start+wordsLen*wordLen]
			differ[word]++
			if differ[word] == 0 {
				delete(differ, word)
			}
			// 左端点删除一个单词
			word = s[start-wordLen : start]
			differ[word]--
			if differ[word] == 0 {
				delete(differ, word)
			}
			// 如果differ中的单词数量为空, 就说明这段字母由给定的单词组成
			if len(differ) == 0 {
				ans = append(ans, start)
			}
		}
	}

	fmt.Println(c1 * (d1 + d2 + d3))
	return
}

func findSubstringWaste(s string, words []string) []int {
	var c1, c2 int
	var d21, d22 int
	var ret []int
	sLen := len(s)
	wLen := len(words)
	if wLen == 0 {
		return nil
	}
	wordLen := len(words[0])
	strLen := wLen * wordLen
	if sLen < strLen {
		return nil
	}
	var srcWords = make(map[string]int, wLen)
	for _, v := range words {
		c1++
		srcWords[v] = srcWords[v] + 1
	}

wrapper:
	for i := 0; i <= sLen-strLen; i++ {
		c2++
		findWorlds := make(map[string]int)
		for j := 0; j < wLen; j++ {
			d21++
			str := s[i+j*wordLen : i+j*wordLen+wordLen]

			if _, exist := srcWords[str]; !exist {
				continue wrapper
			}

			if findWorlds[str]+1 > srcWords[str] {
				continue wrapper
			}
			findWorlds[str] = findWorlds[str] + 1
		}

		ret = append(ret, i)
	}

	fmt.Println(c1 + c2*(d21+d22))
	return ret
}
