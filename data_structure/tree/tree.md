### 二叉树
- 树的节点数量：n个叶子节点 + (n-1)个非叶节点
- 完美二叉树(Perfect Binary Tree)：Every node except the leaf nodes have two children and every level (last level too) is completely filled. 
- 完全二叉树(Complete Binary Tree): Every level except the last level is completely filled and all the nodes are left justified.
- 完满二叉树(Full Binary Tree): Every node except the leaf nodes have two children.
- 如果使用数组来表示树，根节点下标为0，任一节点下标为n，则左右儿子下标为(2*n)+1和(2*n)+2；任意下标为n的节点，其父节点下标为floor((n-1)/2)
- 根节点的层数为0；h层的完美二叉树，共有 2^h 个叶子节点