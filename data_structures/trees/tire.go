package trees

type tireNode struct {
	token    rune
	children map[rune]*tireNode
	words    string
}

type Tire struct {
	root *tireNode
}

// AddWord 添加内容
func (t *Tire) AddWord(words string) {
	if t.root == nil {
		t.root = &tireNode{children: make(map[rune]*tireNode)}
	}
	cur := t.root
	for _, char := range words {
		if _, ok := cur.children[char]; !ok {
			cur.children[char] = &tireNode{children: make(map[rune]*tireNode), token: char}
		}
		cur = cur.children[char]
	}
	cur.words = words
}

// Retrieve 找出前缀为prefix 的所有句子
func (t *Tire) Retrieve(prefix string) []string {
	if t.root == nil || len(t.root.children) == 0 {
		return nil
	}
	var result []string
	cur := t.root
	for _, ch := range prefix {
		if _, ok := cur.children[ch]; !ok {
			return nil
		}
		cur = cur.children[ch]
	}

	// DFS
	root := cur
	stack := []*tireNode{root}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if top.words != "" {
			result = append(result, top.words)
		}

		for _, child := range top.children {
			stack = append(stack, child)
		}
	}

	return result
}
