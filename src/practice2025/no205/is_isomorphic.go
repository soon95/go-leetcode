package no205

func isIsomorphic(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	fromMap, toMap := make(map[byte]byte), make(map[byte]byte)

	for i := range s {
		sCh := s[i]
		tCh := t[i]

		if toCh, exist := fromMap[sCh]; exist && toCh != tCh {
			return false
		}
		if fromCh, exist := toMap[tCh]; exist && fromCh != sCh {
			return false
		}

		fromMap[sCh] = tCh
		toMap[tCh] = sCh
	}

	return true
}
