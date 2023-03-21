package main

import (
	"fmt"
)

/*
10. 正则表达式匹配
https://leetcode.cn/problems/regular-expression-matching/
*/
func isMatch(s string, p string) bool {

	dic := make(map[string]map[string]bool)

	return doIsMatch(s, p, dic)
}

func doIsMatch(s string, p string, dic map[string]map[string]bool) bool {

	sLength := len(s)
	pLength := len(p)

	if sLength == 0 && pLength == 0 {
		// 表示匹配到最后了
		return true
	} else if sLength == 0 {
		if pLength > 1 && p[1] == '*' {
			return isMatchFromDic(s, p[2:], dic)
		} else {
			return false
		}
	} else if pLength == 0 {
		return false
	}

	if pLength == 1 {
		if s[0] == p[0] || p[0] == '.' {
			return isMatchFromDic(s[1:], p[1:], dic)
		} else {
			return false
		}
	}

	if p[1] == '*' {

		if s[0] == p[0] || p[0] == '.' {
			// 匹配上了
			return isMatchFromDic(s, p[2:], dic) || isMatchFromDic(s[1:], p[2:], dic) || isMatchFromDic(s[1:], p, dic)

		} else {
			// 没匹配上 则需要将*看做0
			return isMatchFromDic(s, p[2:], dic)

		}

	} else {
		if p[0] == '.' {
			// 匹配 .
			return isMatchFromDic(s[1:], p[1:], dic)
		} else {
			// 匹配 a
			return s[0] == p[0] && isMatchFromDic(s[1:], p[1:], dic)
		}

	}
}

func isMatchFromDic(s string, p string, dic map[string]map[string]bool) bool {
	m, exist := getMatchFromDic(dic, s, p)
	if exist {
		return m
	} else {
		match := doIsMatch(s, p, dic)
		fillBackDic(dic, s, p, match)
		return match
	}
}

func fillBackDic(dic map[string]map[string]bool, s string, p string, match bool) {
	pMatch, pExist := dic[p]

	if !pExist {
		pMatch = make(map[string]bool)
		dic[p] = pMatch
	}

	pMatch[s] = match
}

func getMatchFromDic(dic map[string]map[string]bool, s string, p string) (bool, bool) {

	pMatch, pExist := dic[p]

	if !pExist {
		return false, false
	}

	sMatch, sExist := pMatch[s]

	if !sExist {
		return false, false
	}

	return sMatch, true
}

func main() {

	s := "aaaaaaaaaaaaaaaaaaab"
	p := "a*a*a*a*a*a*a*a*a*a*"

	fmt.Println(isMatch(s, p))
}
