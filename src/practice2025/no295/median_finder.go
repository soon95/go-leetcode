package no295

type MedianFinder struct {
	MinHeap []int
	MaxHeap []int
}

func Constructor() MedianFinder {
	return MedianFinder{
		MinHeap: make([]int, 0),
		MaxHeap: make([]int, 0),
	}
}

func (m *MedianFinder) AddNum(num int) {

	if len(m.MinHeap) == 0 {
		m.MinHeap = append(m.MinHeap, num)
		return
	}

	if num > m.MinHeap[0] {

		if len(m.MinHeap) > len(m.MaxHeap) {
			tmp := m.getFromMinHeap()
			m.addToMinHeap(num)
			m.addToMaxHeap(tmp)
		} else {
			m.addToMinHeap(num)
		}

	} else if len(m.MaxHeap) == 0 || num < m.MaxHeap[0] {
		if len(m.MinHeap) < len(m.MaxHeap) {

			tmp := m.getFromMaxHeap()
			m.addToMaxHeap(num)
			m.addToMinHeap(tmp)

		} else {
			m.addToMaxHeap(num)
		}
	} else {
		if len(m.MinHeap) < len(m.MaxHeap) {
			m.addToMinHeap(num)
		} else {
			m.addToMaxHeap(num)
		}

	}
}

func (m *MedianFinder) FindMedian() float64 {

	if len(m.MinHeap) > len(m.MaxHeap) {
		return float64(m.MinHeap[0])
	} else if len(m.MinHeap) < len(m.MaxHeap) {
		return float64(m.MaxHeap[0])
	} else {
		return 0.5 * float64(m.MinHeap[0]+m.MaxHeap[0])
	}
}

func (m *MedianFinder) addToMinHeap(num int) {
	m.MinHeap = append(m.MinHeap, num)

	index := len(m.MinHeap) - 1

	for index > 0 {
		pIndex := (index - 1) / 2

		if m.MinHeap[pIndex] > m.MinHeap[index] {
			m.MinHeap[pIndex], m.MinHeap[index] = m.MinHeap[index], m.MinHeap[pIndex]
			index = pIndex
		} else {
			break
		}
	}
}

func (m *MedianFinder) addToMaxHeap(num int) {

	m.MaxHeap = append(m.MaxHeap, num)

	index := len(m.MaxHeap) - 1

	for index > 0 {
		pIndex := (index - 1) / 2

		if m.MaxHeap[pIndex] < m.MaxHeap[index] {
			m.MaxHeap[pIndex], m.MaxHeap[index] = m.MaxHeap[index], m.MaxHeap[pIndex]
			index = pIndex
		} else {
			break
		}
	}
}

func (m *MedianFinder) getFromMinHeap() int {

	top := m.MinHeap[0]

	m.MinHeap[0] = m.MinHeap[len(m.MinHeap)-1]

	m.MinHeap = m.MinHeap[:len(m.MinHeap)-1]

	pIndex := 0

	for pIndex < len(m.MinHeap) {

		lIndex := pIndex*2 + 1
		rIndex := pIndex*2 + 2

		if rIndex < len(m.MinHeap) {

			if m.MinHeap[pIndex] < m.MinHeap[rIndex] && m.MinHeap[pIndex] < m.MinHeap[lIndex] {
				break
			} else if m.MinHeap[lIndex] < m.MinHeap[rIndex] {
				m.MinHeap[pIndex], m.MinHeap[lIndex] = m.MinHeap[lIndex], m.MinHeap[pIndex]
				pIndex = lIndex
			} else {
				m.MinHeap[pIndex], m.MinHeap[rIndex] = m.MinHeap[rIndex], m.MinHeap[pIndex]
				pIndex = rIndex
			}

		} else if lIndex < len(m.MinHeap) {
			if m.MinHeap[pIndex] > m.MinHeap[lIndex] {
				m.MinHeap[pIndex], m.MinHeap[lIndex] = m.MinHeap[lIndex], m.MinHeap[pIndex]
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

func (m *MedianFinder) getFromMaxHeap() int {
	top := m.MaxHeap[0]

	m.MaxHeap[0] = m.MaxHeap[len(m.MaxHeap)-1]

	m.MaxHeap = m.MaxHeap[:len(m.MaxHeap)-1]

	pIndex := 0

	for pIndex < len(m.MaxHeap) {

		lIndex := pIndex*2 + 1
		rIndex := pIndex*2 + 2

		if rIndex < len(m.MaxHeap) {

			if m.MaxHeap[pIndex] > m.MaxHeap[rIndex] && m.MaxHeap[pIndex] > m.MaxHeap[lIndex] {
				break
			} else if m.MaxHeap[lIndex] > m.MaxHeap[rIndex] {
				m.MaxHeap[pIndex], m.MaxHeap[lIndex] = m.MaxHeap[lIndex], m.MaxHeap[pIndex]
				pIndex = lIndex
			} else {
				m.MaxHeap[pIndex], m.MaxHeap[rIndex] = m.MaxHeap[rIndex], m.MaxHeap[pIndex]
				pIndex = rIndex
			}

		} else if lIndex < len(m.MaxHeap) {
			if m.MaxHeap[pIndex] < m.MaxHeap[lIndex] {
				m.MaxHeap[pIndex], m.MaxHeap[lIndex] = m.MaxHeap[lIndex], m.MaxHeap[pIndex]
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
