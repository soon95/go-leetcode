package no383

func canConstruct(ransomNote string, magazine string) bool {

	chMap := make(map[int32]int)

	for _, ch := range magazine {
		chMap[ch]++
	}

	for _, ch := range ransomNote {
		cnt := chMap[ch]
		if cnt <= 0 {
			return false
		}
		chMap[ch]--
	}
	return true
}
