package utils

// BuildTreeByRecursive 通过递归建树
func BuildTreeByRecursive(offices []ForkTreeNode) []ForkTreeNode {
	//循环遍历每一个节点，寻找每一个节点的子节点
	//x节点找到父节点后，就把x节点移动到父节点的children中，然后标记x节点为子节点
	//因为find children为递归，所以子节点无需重复进行root find children
	var result = make([]ForkTreeNode, 0)
	var flag = make([]bool, len(offices))
	for i := 0; i < len(offices); i++ {
		if !flag[i] {
			a := findChildren(&offices[i], offices, &flag)
			result = append(result, *a)
		}
	}
	return result
}

func findChildren(sequence *ForkTreeNode, sequences []ForkTreeNode, flag *[]bool) *ForkTreeNode {
	for i := 0; i < len(sequences); i++ {
		o := (sequences)[i]
		if (*sequence).GetID() == o.GetPID() {
			(*flag)[i] = true
			var s []ForkTreeNode
			s = append(s, *findChildren(&o, sequences, flag))
			(*sequence).Following(s)
		}
	}
	return sequence
}
