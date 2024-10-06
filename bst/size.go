package bst

var (
	nodeSize = map[*Node]int{}
)

func (root *Node) Size() int {
	return nodeSize[root]
}

func (root *Node) incrementSize() {
	nodeSize[root] += 1
}

func (root *Node) incrementSizeBy(i int) {
	nodeSize[root] += i
}
