package dfa

// SearchTree 多叉树结构体
type SearchTree struct {
	isEnd    bool
	children map[rune]*SearchTree
}

// 创建个子节点
func (s *SearchTree) newTreeChildNode(key rune) *SearchTree {
	if node, exist := s.children[key]; exist { // 已经存在该节点的情况下
		return node
	}
	node := &SearchTree{ // 不存在该节点的情况
		isEnd:    false,
		children: make(map[rune]*SearchTree, 3),
	}
	s.children[key] = node
	return node
}

// BuildSearchTree 创建多叉搜索树
func BuildSearchTree(words []string) *SearchTree {
	head := &SearchTree{
		isEnd:    false,
		children: make(map[rune]*SearchTree, len(words)), // 手动设置切片容量，避免频繁扩容
	}
	for _, word := range words { // 遍历每一个词
		parentNode := head
		for _, char := range word {
			parentNode = parentNode.newTreeChildNode(char) // 逐字向父节点添加子节点
		}
		parentNode.isEnd = true // 词的最后一个字 isEnd 状态要设置成true
	}
	return head
}

// Detection 检测目标文本是否包含了指定的词，有的话返回词第一个字所在索引，没则返回-1
func Detection(target string, tree *SearchTree) int {
	if !(tree != nil && tree.isEnd != true) {
		return -1
	}
	charArray := []rune(target) // 转成rune数组
	for charIndex, char := range charArray {
		if node, exist := tree.children[char]; exist {
			if node.isEnd { // 检查这个词是不是就一个字
				return charIndex
			} else {
				for subsetIndex := charIndex + 1; subsetIndex < len(charArray); subsetIndex++ {
					if child, childExist := node.children[charArray[subsetIndex]]; childExist {
						if child.isEnd { // 匹配到了词最后一个字
							return charIndex
						}
						node = child
					} else {
						break
					}
				}
			}
		}
	}
	return -1
}
