package no13

var symbolMap = map[string]int{
	"I":  1,
	"V":  5,
	"IV": 4,
	"IX": 9,
	"X":  10,
	"XL": 40,
	"L":  50,
	"XC": 90,
	"C":  100,
	"CD": 400,
	"D":  500,
	"CM": 900,
	"M":  1000,
}

func romanToInt(s string) int {

	num := 0

	for i := 0; i < len(s); i++ {

		symbol := ""

		if i < len(s)-1 {
			_, exist := symbolMap[s[i:i+2]]
			if exist {
				symbol = s[i : i+2]
				i++
			} else {
				symbol = s[i : i+1]
			}
		} else {
			symbol = s[i : i+1]
		}

		num += symbolMap[symbol]
	}

	return num
}
