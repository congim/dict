/**
 * @Author:
 * @Description:
 * @File:  node
 * @Date: 2021/5/8 3:18 下午
 */

package dict

// Node Trie树上的一个节点.
type Node struct {
	isRootNode bool
	isPathEnd  bool
	Character  rune
	Children   map[rune]*Node
}

// NewNode 新建子节点
func NewNode(character rune) *Node {
	return &Node{
		Character: character,
		Children:  make(map[rune]*Node, 0),
	}
}

// NewRootNode 新建根节点
func NewRootNode(character rune) *Node {
	return &Node{
		isRootNode: true,
		Character:  character,
		Children:   make(map[rune]*Node, 0),
	}
}

// IsLeafNode 判断是否叶子节点
func (n *Node) IsLeafNode() bool {
	return len(n.Children) == 0
}

// IsRootNode 判断是否为根节点
func (n *Node) IsRootNode() bool {
	return n.isRootNode
}

// IsPathEnd 判断是否为某个路径的结束
func (n *Node) IsPathEnd() bool {
	return n.isPathEnd
}

// SoftDel 置软删除状态
func (n *Node) SoftDel() {
	n.isPathEnd = false
}
