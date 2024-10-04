package bst

func (root *Node) HardDelete(val int) error {
	return nil
}

func (root *Node) SoftDelete(val int) error {
	node, err := root.Find(val)
	if err != nil {
		return err
	}
	node.deleted = true
	return nil
}
