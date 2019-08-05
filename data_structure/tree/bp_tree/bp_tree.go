package bp_tree

type CmpFunc func(a, b interface{}) int
type KeyType interface{}

type BpTree struct {
	root  *BpNode // tree root
	cmp   CmpFunc
	order int
}

type BpNode struct {
	keys     []KeyType
	pointers []interface{}
	parent   *BpNode
	keyNum   int
	isLeaf   bool
}

// locate target node for the key
func (t *BpTree) locate(cmp CmpFunc, key KeyType) *BpNode {
	cur := t.root
	for !cur.isLeaf {
		i := 0
		for i < cur.keyNum {
			if cmp(cur.keys[i], key) < 0 {
				break
			}
			i++
		}
		cur = cur.pointers[i].(*BpNode)
	}
	return cur
}

// insert the key
func (t *BpTree) insert(cmp CmpFunc, key KeyType, val interface{}) {
	// locate which node the key will insert to
	leaf := t.locate(cmp, key)

	var idx = len(leaf.keys)
	for i := 0; i < leaf.keyNum; i++ {
		if r := cmp(key, leaf.keys[i]); r > 0 {
			continue
		} else if r == 0 {
			// the key has existed, just update
			leaf.pointers[i] = val
			return
		} else {
			idx = i
			break
		}
	}

	// insert directly
	ln := len(leaf.keys)
	if ln < t.order-1 {
		leaf.keyNum++
		copy(leaf.keys[idx+1:ln+1], leaf.keys[idx:ln])
		copy(leaf.pointers[idx+1:ln+1], leaf.pointers[idx:ln])
		leaf.keys[idx] = key
		leaf.pointers[idx] = val
		return
	}

	// split the left
	mid := getSplitIndex(t.order)

	tempKeys := make([]KeyType, t.order)
	tempPtrs := make([]interface{}, t.order)

	copy(tempKeys, leaf.keys[:idx])
	copy(tempPtrs, leaf.pointers[:idx])

	tempKeys[idx] = key
	tempPtrs[idx] = val

	copy(tempKeys[idx+1:], leaf.keys[idx:ln])
	copy(tempPtrs[idx+1:], leaf.pointers[idx:ln])

	copy(leaf.keys, tempKeys[:mid])
	copy(leaf.pointers, tempPtrs[:mid])
	leaf.keyNum = mid

	sibling := newNode(t.order)
	sibling.isLeaf = true

	copy(sibling.keys, tempKeys[mid:])
	copy(sibling.pointers, tempPtrs[mid:])
	sibling.keyNum = len(tempKeys) - mid + 1

	// set the next link for leaf
	nextLinkIdx := t.order - 1
	sibling.pointers[nextLinkIdx] = leaf.pointers[nextLinkIdx]
	leaf.pointers[nextLinkIdx] = sibling

	t.insertIntoParent(leaf, sibling)
}

func (t *BpTree) insertIntoParent(left, right *BpNode) {
	// left is the parent
	if left.parent == nil {
		newRoot := newNode(t.order)

		newRoot.keyNum = 1
		newRoot.keys[0] = right.keys[0]
		newRoot.pointers[0] = left
		newRoot.pointers[1] = right

		left.parent = newRoot
		right.parent = newRoot

		t.root = newRoot
		return
	}
	parent := left.parent
	var leftAt int
	for i := 0; i < parent.keyNum; i++ {
		if left == parent.pointers[i] {
			leftAt = i
			break
		}
	}

	if parent.keyNum < t.order-1 {
		ln := parent.keyNum
		parent.keyNum++

		copy(parent.keys[leftAt+1:ln+1], parent.keys[leftAt:ln])
		copy(parent.pointers[leftAt+2:ln+1], parent.pointers[leftAt+1:ln])

		parent.keys[leftAt] = right.keys[0]
		parent.pointers[leftAt+1] = right
		return
	}

	// split the parent
	

}

func getSplitIndex(length int) int {
	if length%2 == 0 {
		return length / 2
	}

	return length/2 + 1
}

func newNode(order int) *BpNode {
	return &BpNode{
		keys:     make([]KeyType, order-1),
		pointers: make([]interface{}, order),
	}
}
