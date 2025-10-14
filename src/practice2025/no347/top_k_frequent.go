package no347

func topKFrequent(nums []int, k int) []int {

	cntMap := make(map[int]int)

	for _, num := range nums {
		cntMap[num]++
	}

	maxHeap := NewMaxHeap()
	for key, value := range cntMap {
		maxHeap.add(&NumCnt{Num: key, Cnt: value})
	}

	ans := make([]int, 0)

	for k > 0 {
		ans = append(ans, maxHeap.get().Num)
		k--
	}

	return ans
}

type MaxHeap struct {
	Data []*NumCnt
}
type NumCnt struct {
	Num int
	Cnt int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		Data: make([]*NumCnt, 0),
	}
}

func (m *MaxHeap) add(numCnt *NumCnt) {

	m.Data = append(m.Data, numCnt)

	index := len(m.Data) - 1

	for index > 0 {
		pIndex := (index - 1) / 2

		if m.Data[pIndex].Cnt < m.Data[index].Cnt {
			m.Data[pIndex], m.Data[index] = m.Data[index], m.Data[pIndex]
			index = pIndex
		} else {
			break
		}
	}
}

func (m *MaxHeap) get() *NumCnt {

	top := m.Data[0]

	m.Data[0] = m.Data[len(m.Data)-1]

	m.Data = m.Data[:len(m.Data)-1]

	pIndex := 0

	for pIndex < len(m.Data) {

		lIndex := pIndex*2 + 1
		rIndex := pIndex*2 + 2

		if rIndex < len(m.Data) {

			if m.Data[pIndex].Cnt > m.Data[rIndex].Cnt && m.Data[pIndex].Cnt > m.Data[lIndex].Cnt {
				break
			} else if m.Data[lIndex].Cnt > m.Data[rIndex].Cnt {
				m.Data[pIndex], m.Data[lIndex] = m.Data[lIndex], m.Data[pIndex]
				pIndex = lIndex
			} else {
				m.Data[pIndex], m.Data[rIndex] = m.Data[rIndex], m.Data[pIndex]
				pIndex = rIndex
			}

		} else if lIndex < len(m.Data) {
			if m.Data[pIndex].Cnt < m.Data[lIndex].Cnt {
				m.Data[pIndex], m.Data[lIndex] = m.Data[lIndex], m.Data[pIndex]
				pIndex = lIndex
			} else {
				break
			}
		} else {
			break
		}
	}

	return top
}
