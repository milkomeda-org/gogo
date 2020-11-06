package gogo

// ForkTreeNode 分叉树节点
type ForkTreeNode interface {
	GetID() int
	GetPID() int
	Following([]ForkTreeNode)
}
