# tree_fork
> 建树函数

example:
```
var nodes []gogo.ForkTreeNode
result := gogo.BuildTreeByRecursive(nodes)
```
 	
> ForkTreeNode接口
- GetID() uint
- GetPID() uint
- Following([]ForkTreeNode)

example:
 ```
type PositionSerializer struct {
 	model.BaseModel
 	ParentID uint  // 上级ID
 	OfficeID uint  // 组织ID
 	Children []gogo.ForkTreeNode
 }
 func (receiver *PositionSerializer) GetID() uint {
 	return receiver.ID
 }
 func (receiver *PositionSerializer) GetPID() uint {
 	return receiver.ParentID
 }
 func (receiver *PositionSerializer) Following(v []gogo.ForkTreeNode) {
 	receiver.Children = append(receiver.Children, v...)
 }
 ```