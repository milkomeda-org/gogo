package gogo

// ForkTreeNode 分叉树节点
type ForkTreeNode interface {
	// GetID this method should return the id of the sortable elements
	// GetID 此方法应该返回元素的唯一id
	GetID() int
	// GetPID this method should return the parent id of the sortable elements
	// GetPID 此方法应该返回元素的父元素id
	GetPID() int
	// Following this method be used to receive child elements
	// Following 此方法用于接受子元素
	Following([]ForkTreeNode)
}
