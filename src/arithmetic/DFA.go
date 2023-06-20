package arithmetic

type searchTree struct {
	isEnd    bool
	children map[rune]searchTree
}

// 创建个子节点
func (s *searchTree) newTreeChildNode(key rune) *searchTree {
	node := searchTree{
		isEnd:    false,
		children: make(map[rune]searchTree, 3),
	}
	s.children[key] = node
	return &node
}

// BuildSearchTree 创建多叉搜索树
func BuildSearchTree(words []string) *searchTree {
	head := &searchTree{
		isEnd:    false,
		children: make(map[rune]searchTree, len(words)),
	}
	for _, word := range words { // 遍历每一个词
		parentNode := head
		for _, char := range word { // 遍历每一个词的每一个字
			parentNode = parentNode.newTreeChildNode(char) // 逐字向父节点添加子节点
		}
		parentNode.isEnd = true // 最后一个字标记一下，设置成true
	}
	return head
}
