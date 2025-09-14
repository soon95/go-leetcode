package no239

func maxSlidingWindow(nums []int, k int) []int {

	maxDump := NewMaxDump(nums)

	left, right := 0, k-1

	for i := 0; i < right; i++ {
		maxDump.add(i)
	}

	res := make([]int, 0)

	for right < len(nums) {

		maxDump.add(right)

		peekIndex := maxDump.peek()
		for peekIndex < left {
			maxDump.delete()
			peekIndex = maxDump.peek()
		}

		res = append(res, nums[maxDump.peek()])

		right++
		left++
	}

	return res
}

type MaxDump struct {
	nums      []int
	indexDump []int
}

func NewMaxDump(nums []int) *MaxDump {
	return &MaxDump{
		nums:      nums,
		indexDump: make([]int, 0),
	}
}

func (m *MaxDump) peek() int {
	if len(m.indexDump) == 0 {
		return 0
	}
	return m.indexDump[0]
}

func (m *MaxDump) add(index int) {

	m.indexDump = append(m.indexDump, index)

	m.float(len(m.indexDump) - 1)

}

func (m *MaxDump) delete() {

	m.indexDump[0] = m.indexDump[len(m.indexDump)-1]

	m.indexDump = m.indexDump[:len(m.indexDump)-1]

	m.drop(0)
}

func (m *MaxDump) float(dumpIndex int) {

	if dumpIndex == 0 {
		return
	}

	pIndex := (dumpIndex+1)/2 - 1

	if m.nums[m.indexDump[pIndex]] < m.nums[m.indexDump[dumpIndex]] {
		m.indexDump[pIndex], m.indexDump[dumpIndex] = m.indexDump[dumpIndex], m.indexDump[pIndex]
	} else {
		return
	}

	m.float(pIndex)
}

func (m *MaxDump) drop(dumpIndex int) {

	leftChildIndex, rightChildIndex := dumpIndex*2+1, dumpIndex*2+2

	if leftChildIndex >= len(m.indexDump) {
		// 叶子节点
		return
	}

	if rightChildIndex < len(m.indexDump) {
		// 有右子树

		if m.nums[m.indexDump[dumpIndex]] > m.nums[m.indexDump[leftChildIndex]] && m.nums[m.indexDump[dumpIndex]] > m.nums[m.indexDump[rightChildIndex]] {
			return
		} else if m.nums[m.indexDump[leftChildIndex]] > m.nums[m.indexDump[rightChildIndex]] {
			m.indexDump[dumpIndex], m.indexDump[leftChildIndex] = m.indexDump[leftChildIndex], m.indexDump[dumpIndex]
			m.drop(leftChildIndex)
		} else {
			m.indexDump[dumpIndex], m.indexDump[rightChildIndex] = m.indexDump[rightChildIndex], m.indexDump[dumpIndex]
			m.drop(rightChildIndex)
		}

	} else {
		// 无右子树
		if m.nums[m.indexDump[dumpIndex]] > m.nums[m.indexDump[leftChildIndex]] {
			return
		} else if m.nums[m.indexDump[dumpIndex]] <= m.nums[m.indexDump[leftChildIndex]] {
			m.indexDump[dumpIndex], m.indexDump[leftChildIndex] = m.indexDump[leftChildIndex], m.indexDump[dumpIndex]
			m.drop(leftChildIndex)
		}
	}
}
