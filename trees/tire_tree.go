package trees

type tireTreeNode struct {
	children map[rune]*tireTreeNode
	isEnding bool
}

type TireTree struct {
	root *tireTreeNode
}

// BuildTireTree 构建字典树
func BuildTireTree(words []string) (tt *TireTree) {
	node := &tireTreeNode{children: map[rune]*tireTreeNode{}, isEnding: false}
	for _, word := range words {
		runes := []rune(word)
		for _, r := range runes {
			temp, ok := node.children[r]
			if !ok {
				temp = &tireTreeNode{children: map[rune]*tireTreeNode{}}
				node.children[r] = temp
			}
			node = temp
		}
		node.isEnding = true
	}
	return
}

// InsertNode 向字典插入值
func (tt *TireTree) InsertNode(word string) {
	runes := []rune(word)
	node := tt.root
	for _, val := range runes {
		if node.children[val] == nil {
			node.children[val] = new(tireTreeNode)
		}
		node = node.children[val]
	}
	node.isEnding = true
}

// FindSensitiveWords 找出输入中包含的敏感词
func (tt *TireTree) FindSensitiveWords(word string) (res []string) {
	if word == "" {
		return nil
	}
	node := tt.root
	left, right := 0, 0
	runes := []rune(word)
	for len(runes) > left {
		w := runes[right]
		node = node.children[w]
		if node == nil {
			left++
			right = left
			node = tt.root
		} else if node.isEnding {
			temp := runes[left : right+1]
			res = append(res, string(temp))
			right++
			left = right
			node = tt.root
		} else {
			if right < len(runes)-1 {
				right++
			} else {
				left++
				right = left
				node = tt.root
			}
		}
	}
	return
}
