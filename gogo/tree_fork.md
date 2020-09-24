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
    ID uint
 	ParentID uint
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

case：
- 菜单栏
- 文件树
- 人力资源