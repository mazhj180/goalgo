package list

import (
	"fmt"
	"math"
	"math/rand"
)

var (
	skipRate     = 0.25
	maximumLayer = 30
)

type skipListNode struct {
	score int
	elem  []any
	next  []*skipListNode
}

type SkipList struct {
	head        *skipListNode
	size, level int
}

func NewSkipList() *SkipList {
	return &SkipList{
		head: &skipListNode{
			math.MinInt64,
			nil,
			make([]*skipListNode, maximumLayer),
		},
		level: 1,
	}
}

func (s *SkipList) Insert(elem any, score int) {
	defer func() { s.size++ }()

	update := make([]*skipListNode, maximumLayer)

	// 找到每一层需要插入的位置
	cur := s.head
	for i := s.level - 1; i >= 0; i-- {
		for cur.next[i] != nil && cur.next[i].score < score {
			cur = cur.next[i]
		}
		update[i] = cur
	}

	// 查找是否已经存在该key
	if cur.next[0] != nil && cur.next[0].score == score {
		cur.elem = append(cur.elem, elem)
		return
	}

	// 随机生成新节点的层级
	level := s.randomLevel()
	if level > maximumLayer {
		level = maximumLayer
	}

	// 更新层级
	if level > s.level {
		for i := s.level; i < level; i++ {
			update[i] = s.head
		}
		s.level = level
	}

	// 插入新节点
	node := &skipListNode{
		score: score,
		elem:  []any{elem},
		next:  make([]*skipListNode, level),
	}
	for i := 0; i < level; i++ {
		node.next[i] = update[i].next[i]
		update[i].next[i] = node
	}
}

func (s *SkipList) Remove(score int) {
	defer func() { s.size-- }()

	update := make([]*skipListNode, maximumLayer)
	cur := s.head
	for i := s.level - 1; i >= 0; i-- {
		for cur.next[i] != nil && cur.next[i].score < score {
			cur = cur.next[i]
		}
		update[i] = cur
	}
	tar := cur.next[0]
	if tar == nil || tar.score != score {
		return
	}

	// 删除各层级的节点
	for i := 0; i < s.level; i++ {
		if update[i].next[i] != tar {
			break
		}
		update[i].next[i] = tar.next[i]
	}

	// 最高层级无节点，降级
	for s.level > 1 && s.head.next[s.level-1] == nil {
		s.level--
	}
}

func (s *SkipList) Search(score int) ([]any, bool) {
	cur := s.head
	for i := s.level - 1; i >= 0; i-- {
		for cur.next[i] != nil && cur.next[i].score < score {
			cur = cur.next[i]
		}
	}
	cur = cur.next[0]
	if cur != nil && cur.score == score {
		return cur.elem, true
	}
	return nil, false
}

func (s *SkipList) Size() int {
	return s.size
}

func (s *SkipList) Level() int {
	return s.level
}

func (s *SkipList) randomLevel() int {
	level := 1
	for rand.Int()&0xffff < int(skipRate*0xffff) {
		level++
	}
	return level
}

func Print(sl *SkipList) {
	for i := sl.level - 1; i >= 0; i-- {
		cur := sl.head
		fmt.Printf("Level %d: ", i+1)
		for cur.next[i] != nil {
			fmt.Printf("[%s::%d] ", cur.next[i].elem[0], cur.next[i].score)
			cur = cur.next[i]
		}
		fmt.Println()
	}
}
