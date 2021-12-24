package main

type HeapNode struct {
	riskSoFar int
	x, y      int
	depth     int
}

type Heap struct {
	nodes []HeapNode
}

func (h *Heap) Insert(hn HeapNode) {
	h.nodes = append(h.nodes, hn)
	h.swim(len(h.nodes) - 1)
}

func (h *Heap) swim(idx int) {
	if idx == 0 {
		return
	}

	self := h.nodes[idx]
	pIdx := (idx - 1) / 2
	parent := h.nodes[pIdx]
	if self.riskSoFar < parent.riskSoFar {
		h.nodes[idx], h.nodes[pIdx] = parent, self
		h.swim(pIdx)
	} else if self.riskSoFar == parent.riskSoFar && self.depth > parent.depth {
		h.nodes[idx], h.nodes[pIdx] = parent, self
		h.swim(pIdx)
	}
}

func (h *Heap) Pop() HeapNode {
	head := h.nodes[0]

	if len(h.nodes) == 1 {
		h.nodes = nil
	} else {
		n := len(h.nodes) - 1
		last := h.nodes[n]
		h.nodes[0] = last
		h.nodes = h.nodes[:n]
		h.sink(0)
	}

	return head
}

func (h *Heap) sink(idx int) {
	self := h.nodes[idx]

	lIdx := 2*idx + 1
	if lIdx >= len(h.nodes) {
		return
	}
	left := h.nodes[lIdx]
	if self.riskSoFar > left.riskSoFar {
		h.nodes[idx], h.nodes[lIdx] = left, self
		h.sink(lIdx)
		self = left
	} else if self.riskSoFar == left.riskSoFar && self.depth < left.depth {
		h.nodes[idx], h.nodes[lIdx] = left, self
		h.sink(lIdx)
		self = left
	}

	rIdx := 2*idx + 2
	if rIdx >= len(h.nodes) {
		return
	}
	right := h.nodes[rIdx]
	if self.riskSoFar > right.riskSoFar {
		h.nodes[idx], h.nodes[rIdx] = right, self
		h.sink(rIdx)
	} else if self.riskSoFar == right.riskSoFar && self.depth < right.depth {
		h.nodes[idx], h.nodes[rIdx] = right, self
		h.sink(rIdx)
	}
}
