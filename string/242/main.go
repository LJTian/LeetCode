package main

import (
	"sort"
)

// 1：hash 思路对了，写的代码不怎么好，太多。  isAnagram 是官方答案，写的很精简
// 2: 排序校验字符串是否相等(如果是用官方库的话，这个很简单),注意：需要先切换类型

func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	s1, t1 := []byte(s), []byte(t)
	sort.Slice(s1, func(i, j int) bool {
		return s1[i] > s1[j]
	})
	sort.Slice(t1, func(i, j int) bool {
		return t1[i] > t1[j]
	})
	return string(s1) == string(t1)
}

func isAnagram1s(s string, t string) bool {

	map1 := make(map[rune]int)
	for _, v := range s {
		if q, ok := map1[v]; ok {
			map1[v] = q + 1
		} else {
			map1[v] = 1
		}
	}

	map2 := make(map[rune]int)
	for _, v := range t {
		if q, ok := map2[v]; ok {
			map2[v] = q + 1
		} else {
			map2[v] = 1
		}
	}

	if len(map1) > len(map2) {
		for k, _ := range map1 {
			if map1[k] != map2[k] {
				return false
			}
		}
	} else {
		for k, _ := range map2 {
			if map1[k] != map2[k] {
				return false
			}
		}
	}
	return true
}

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	cnt := map[rune]int{}
	for _, ch := range s {
		cnt[ch]++
	}
	for _, ch := range t {
		cnt[ch]--
		if cnt[ch] < 0 {
			return false
		}
	}
	return true
}
