package no12

import "strings"

func intToRoman(num int) string {

	scaleList := []int{1000, 100, 10, 1}
	secondScaleList := []int{5000, 500, 50, 5}

	ans := ""
	for i, scale := range scaleList {
		cnt := num / scale
		cur := cnt * scale
		num -= cur
		if symbol, exist := symbolMap[cur]; exist {
			ans += symbol
		} else {
			if cur > secondScaleList[i] {
				ans += symbolMap[secondScaleList[i]]
				cnt = (cur - secondScaleList[i]) / scale
			}
			ans += strings.Repeat(symbolMap[scale], cnt)
		}
	}

	return ans
}

var symbolMap = map[int]string{
	1:    "I",
	4:    "IV",
	5:    "V",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M",
}
