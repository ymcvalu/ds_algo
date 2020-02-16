package algorithm

import (
	list2 "container/list"
	"encoding/json"
)

// 二叉树的序列化与反序列化，需要确保二叉树可以序列化为字符串，并且可以将此字符串反序列化为原始树结构。
// 你可以序列化如下的二叉树
//    8
//   / \
//  12  2
//     / \
//    6   4
// 为："[8, 12, 2, null, null, 6, 4, null, null, null, null]"

func treeSerialize(n *Node) string {
	if n == nil {
		return "[]"
	}
	elems := []*int{}
	list := list2.New()
	list.PushBack(n)
	for list.Len() > 0 {
		v := list.Front()
		list.Remove(v)
		n := v.Value.(*Node)
		if n != nil {
			elems = append(elems, &n.Value)
			list.PushBack(n.left)
			list.PushBack(n.right)
		} else {
			elems = append(elems, nil)
		}
	}

	byts, _ := json.Marshal(elems)
	return string(byts)
}

func treeDeserialize(str string) *Node {
	elems := make([]*int, 0)
	json.Unmarshal([]byte(str), &elems)
	if len(elems) == 0 || elems[0] == nil {
		return nil
	}

	list := list2.New()
	n := &Node{
		Value: *elems[0],
	}
	i := 1
	list.PushBack(n)
	for list.Len() > 0 || i < len(elems) {
		n := list.Front().Value.(*Node)
		list.Remove(list.Front())
		if elems[i] != nil {
			n.left = &Node{
				Value: *elems[i],
			}
			list.PushBack(n.left)
		}
		i++

		if i < len(elems) && elems[i] != nil {
			n.right = &Node{
				Value: *elems[i],
			}
			list.PushBack(n.right)
		}
		i++
	}

	return n
}
