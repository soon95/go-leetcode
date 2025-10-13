package no215

func findKthLargest(nums []int, k int) int {

	maxHeap := NewMaxHeap()

	for _, num := range nums {
		maxHeap.add(num)
	}

	for k > 1 {
		maxHeap.get()
		k--
	}

	return maxHeap.get()
}

type MaxHeap struct {
	Data []int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		Data: make([]int, 0),
	}
}

func (m *MaxHeap) add(num int) {

	m.Data = append(m.Data, num)

	index := len(m.Data) - 1

	for index > 0 {
		pIndex := (index - 1) / 2

		if m.Data[pIndex] < m.Data[index] {
			m.Data[pIndex], m.Data[index] = m.Data[index], m.Data[pIndex]
			index = pIndex
		} else {
			break
		}
	}
}

func (m *MaxHeap) get() int {

	top := m.Data[0]

	m.Data[0] = m.Data[len(m.Data)-1]

	m.Data = m.Data[:len(m.Data)-1]

	pIndex := 0

	for pIndex < len(m.Data) {

		lIndex := pIndex*2 + 1
		rIndex := pIndex*2 + 2

		if rIndex < len(m.Data) {

			if m.Data[pIndex] > m.Data[rIndex] && m.Data[pIndex] > m.Data[lIndex] {
				break
			} else if m.Data[lIndex] > m.Data[rIndex] {
				m.Data[pIndex], m.Data[lIndex] = m.Data[lIndex], m.Data[pIndex]
				pIndex = lIndex
			} else {
				m.Data[pIndex], m.Data[rIndex] = m.Data[rIndex], m.Data[pIndex]
				pIndex = rIndex
			}

		} else if lIndex < len(m.Data) {
			if m.Data[pIndex] < m.Data[lIndex] {
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
