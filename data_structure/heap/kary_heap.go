package heap

import (
	"sync"
)

type Unlocker struct{}

func (Unlocker) Lock()   {}
func (Unlocker) Unlock() {}

var _ sync.Locker = Unlocker{}

type Option func(h *KaryHeap)

func Kary(k int) Option {
	return func(h *KaryHeap) {
		h.k = k
	}
}

func Locker(l sync.Locker) Option {
	return func(h *KaryHeap) {
		h.lk = l
	}
}

func NewKaryHeap(cmp func(i, j interface{}) bool, opts ...Option) *KaryHeap {
	h := &KaryHeap{
		k: 2,

		lk:  Unlocker{},
		cmp: cmp,
	}
	for _, opt := range opts {
		opt(h)
	}
	return h
}

type KaryHeap struct {
	array []interface{}
	k     int
	cmp   func(i, j interface{}) bool // i is root if true

	lk sync.Locker
}

func (h *KaryHeap) swap(i, j int) {
	h.array[i], h.array[j] = h.array[j], h.array[i]
}

func (h *KaryHeap) RestoreDown(idx int) {
	if idx < 0 {
		return
	}
	h.lk.Lock()
	if idx+1 >= len(h.array) {
		h.lk.Unlock()
		return
	}
	h.down(idx)
	h.lk.Unlock()
}

func (h *KaryHeap) down(idx int) {
	var (
		ln = len(h.array)
		k  = h.k
	)

	ci, pc := idx, idx
	for ci < ln {
		cl := ci*k + k
		for i := cl - k + 1; i <= cl && i < ln; i++ {
			if h.cmp(h.array[i], h.array[pc]) {
				pc = i
			}
		}
		if ci != pc {
			h.swap(ci, pc)
			ci = pc
		} else {
			break
		}
	}
}

func (h *KaryHeap) RestoreUp(idx int) {
	if idx <= 0 {
		return
	}
	h.lk.Lock()
	if idx > len(h.array)-1 {
		h.lk.Unlock()
		return
	}
	h.up(idx)
	h.lk.Unlock()
}

func (h *KaryHeap) up(idx int) {
	k := h.k
	ci := idx
	parent := (ci - 1) / k
	for ci > 0 {
		if h.cmp(h.array[ci], h.array[parent]) {
			h.swap(ci, parent)
			ci = parent
			parent = (ci - 1) / k
		} else {
			break
		}
	}
}

func (h *KaryHeap) Push(elem interface{}) {
	h.lk.Lock()
	h.push(elem)
	h.lk.Unlock()
}

func (h *KaryHeap) push(elem interface{}) {
	h.array = append(h.array, elem)
	h.up(len(h.array) - 1)
}

func (h *KaryHeap) Pop() interface{} {
	h.lk.Lock()
	defer h.lk.Lock()
	return h.pop()
}

func (h *KaryHeap) pop() interface{} {
	if len(h.array) == 0 {
		return nil
	}
	elem := h.array[0]
	h.swap(0, len(h.array)-1)
	h.array = h.array[:len(h.array)-1]
	h.down(0)
	return elem
}

func (h *KaryHeap) Peek(idx int) interface{} {
	h.lk.Lock()
	defer h.lk.Unlock()
	return h.peek(idx)
}

func (h *KaryHeap) peek(idx int) interface{} {
	if idx < 0 || idx >= len(h.array) {
		return nil
	}
	return h.array[idx]
}

func (h *KaryHeap) Walk(walk func(interface{}) bool) int {
	h.lk.Lock()
	defer h.lk.Unlock()
	for k, v := range h.array {
		if walk(v) {
			return k
		}
	}
	return -1
}

func (h *KaryHeap) Remove(idx int) interface{} {
	h.lk.Lock()
	defer h.lk.Unlock()
	return h.remove(idx)
}

func (h *KaryHeap) remove(idx int) interface{} {
	ret := h.array[idx]
	ln := len(h.array)
	if idx == ln-1 {
		h.array = h.array[:ln-1]
		return ret
	}

	h.array[idx] = h.array[ln-1]
	h.array = h.array[:ln-1]

	if h.cmp(h.array[idx], h.array[(idx-1)/h.k]) {
		h.up(idx)
	} else {
		h.down(idx)
	}

	return ret
}

func (h *KaryHeap) Len() int {
	h.lk.Lock()
	defer h.lk.Unlock()
	return len(h.array)
}

func BuildHeap(arr []interface{}, cmp func(i, j interface{}) bool, opts ...Option) *KaryHeap {
	h := &KaryHeap{
		cmp:   cmp,
		lk:    Unlocker{},
		k:     2,
		array: arr,
	}
	for _, o := range opts {
		o(h)
	}

	ln := len(h.array)
	for i := (ln - 1) / h.k; i >= 0; i-- {
		h.down(i)
	}
	return h
}
