package DFA

type SearchTree struct {
	isEnd    bool
	children map[rune]SearchTree
}

// 创建个子节点
func (s *SearchTree) newTreeChildNode(key rune) *SearchTree {
	node := SearchTree{
		isEnd:    false,
		children: make(map[rune]SearchTree, 3),
	}
	s.children[key] = node
	return &node
}

// BuildSearchTree 创建多叉搜索树
func BuildSearchTree(words []string) *SearchTree {
	head := &SearchTree{
		isEnd:    false,
		children: make(map[rune]SearchTree, len(words)),
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
