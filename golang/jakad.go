package main

import (
	"sort"
	"strings"
)

// Jakad returns Jakad value between two strings.
func Jakad(str1, str2 string, ignoreCase bool) float64 {
	if ignoreCase {
		str1 = strings.ToLower(str1)
		str2 = strings.ToLower(str2)
	}

	set1 := divideStringIntoSet(str1)
	set2 := divideStringIntoSet(str2)

	inter := getIntersection(set1, set2)
	union := getUnion(set1, set2)

	if len(inter) == len(union) {
		return 1
	}

	return float64(len(inter)) / float64(len(union))
}

func divideStringIntoSet(str string) (set []string) {
	for i := 0; i < len(str)-1; i++ {
		elem := string(str[i]) + string(str[i+1])
		if isAlpha(elem) {
			set = append(set, elem)
		}
	}
	return
}

// isAlpha modify this function to meet your condition.
// Say, you can take numbers into consideration if you want.
func isAlpha(str string) bool {
	for _, ch := range str {
		if !(ch >= 'a' && ch <= 'z') &&
			!(ch >= 'A' && ch <= 'Z') {
			return false
		}
	}
	return true
}

func getIntersection(set1 []string, set2 []string) (result []string) {
	for _, s1 := range set1 {
		for _, s2 := range set2 {
			if s1 == s2 {
				result = append(result, s1)
				break
			}
		}
	}
	return
}

func getUnion(set1 []string, set2 []string) (result []string) {
	find := func(str string, set []string) int {
		count := 0
		for _, s := range set {
			if str == s {
				count++
			}
		}
		return count
	}

	result = append(set1, set2...)
	sort.Strings(result)

	u := []string{}
	cur := ""
	for _, s := range result {
		if cur != s {
			u = append(u, s)
		}
		cur = s
	}
	result = nil

	for _, s := range u {
		n1 := find(s, set1)
		n2 := find(s, set2)

		if n1 >= n2 {
			for i := 0; i < n1; i++ {
				result = append(result, s)
			}
		} else {
			for i := 0; i < n2; i++ {
				result = append(result, s)
			}
		}
	}
	return
}
