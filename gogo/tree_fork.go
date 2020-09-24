package gogo

// ForkTreeNode 分叉树节点
type ForkTreeNode interface {
	GetID() uint
	GetPID() uint
	Following([]ForkTreeNode)
}
