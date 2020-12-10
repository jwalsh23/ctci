package data_structures

type IntMinHeap struct {
	slc []int
}

func(h *IntMinHeap) Add(val int) {
	if len(h.slc) < 1 {
		h.slc = []int{0,val}
		return
	}
	h.slc = append(h.slc, val)
	if len(h.slc) < 4 {
		return
	}
	for i := len(h.slc) -1; i > 1; {
		parentIdx := (i+1) /2 -1
		if len(h.slc) % 2 != 0 {
			parentIdx++
		}
		if h.slc[i] > h.slc[parentIdx] {
			break
		}
		h.slc[i], h.slc[parentIdx] = h.slc[parentIdx], h.slc[i]
		i = parentIdx
	}
}

func(h *IntMinHeap) Remove() int {
	if len(h.slc) < 1 {
		return -1
	}
	min := h.slc[1]
	lastItem := len(h.slc) -1
	h.slc[1] = h.slc[lastItem]
	h.slc = h.slc[:lastItem]
	for i := 1; i < len(h.slc) -1; {
		childIdx1, childIdx2 := i*2, i*2 + 1
		swapIdx := childIdx1
		if childIdx1 >= len(h.slc) {
			break
		}
		if childIdx2 >= len(h.slc) {
			swapIdx = childIdx1
		} else {
			if h.slc[childIdx1] > h.slc[childIdx2] {
				swapIdx = childIdx2
			}
		}
		if h.slc[i] > h.slc[swapIdx] {
			h.slc[i], h.slc[swapIdx] = h.slc[swapIdx], h.slc[i]
			i = swapIdx
			continue
		}
		break
	}
	return min
}
