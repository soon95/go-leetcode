package no17

func letterCombinations(digits string) []string {

	if len(digits) == 0 {
		return nil
	}

	res := []string{""}

	for i := range digits {

		ch := digits[i : i+1]

		tmp := make([]string, 0)

		for _, str := range dic[ch] {

			for j := range res {

				tmp = append(tmp, res[j]+string(str))
			}
		}

		res = tmp
	}

	return res
}

var dic = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}
