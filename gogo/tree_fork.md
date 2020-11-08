# tree_fork
> 建树函数

example:
```
var nodes []gogo.ForkTreeNode
result := gogo.BuildTreeByRecursive(nodes)
```
 	
> ForkTreeNode接口
- GetID() int
- GetPID() int
- Following([]ForkTreeNode)

> 元素类型实现ForkTreeNode接口，就能使用建树函数

example:
 ```
type PositionSerializer struct {
    ID int
 	ParentID int
 	Children []gogo.ForkTreeNode
 }
 func (receiver *PositionSerializer) GetID() int {
 	return receiver.ID
 }
 func (receiver *PositionSerializer) GetPID() int {
 	return receiver.ParentID
 }
 func (receiver *PositionSerializer) Following(v []gogo.ForkTreeNode) {
 	receiver.Children = append(receiver.Children, v...)
 }
 ```

case：
- 菜单栏
- 文件树
- 组织关系