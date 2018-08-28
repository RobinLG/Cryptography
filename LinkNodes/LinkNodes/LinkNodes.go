package LinkNodes

import (
	"fmt"
)

// 通过代码完成链表的常用功能，尾插法

// 全局变量记录头节点
var hNode *Node

// 创建节点类型
type Node struct {
	// 数据域
	Data int
	// 地址域
	NextNode *Node
}

// 创建头节点
func CreateHeadNode(data int) *Node {
	//创建Node对象，返回该对象地址
	node := &Node{data, nil}
	hNode = node
	return node
}

// 通过尾插法添加新节点
func AddNode(data int, node *Node) *Node {
	NewNode := &Node{data, nil}
	node.NextNode = NewNode
	return NewNode
}

// 遍历节点
func ShowNodes(head *Node) {
	//接收参数中的头节点
	node := head
	for {
		fmt.Println(node.Data)
		if node.NextNode == nil {
			break
		} else {
			// 实现节点位移
			node = node.NextNode
		}
	}
}

// 按照下标插入新节点
func InsertNodeWithIndex(data int, index int,node *Node) *Node {

	// 如果在在下标为0的位置插入节点相当于创建一个新的头节点
	if index == 0 {
		//创建新节点
		insertedNode := &Node{data, nil}
		insertedNode.NextNode = node
		hNode = insertedNode
		//返回新的头节点
		return insertedNode
	} else if index >= NLen(hNode)-1 {
		// 当输入的节点下标大于等于节点长度-1时，就相当于在链表上插入一个新节点

		// 获得尾节点
		node := hNode
		for {
			if node.NextNode == nil {
				// node是尾节点
				// 在尾节点插入新节点
				insertedNode := &Node{data, nil}
				node.NextNode = insertedNode
				break
			} else {
				node = node.NextNode
			}
		}
	} else {
		// 在链表中间的位置插入新节点
		head := node
		cnt := 1
		for {
			if cnt == index {
				//创建新节点
				insertedNode := &Node{data, nil}
				// 修改链的方向
				insertedNode.NextNode = head.NextNode
				head.NextNode = insertedNode
				break

			} else {
				head = node.NextNode
				cnt++
			}
		}
	}
	return nil
}

// 计算链表总长度
func NLen(head *Node) int {
	// 传入头节点参数
	node := head
	// 用来计算链表长度的标识
	cnt := 1
	for {
		if node.NextNode == nil {
			break
		} else {
			node = node.NextNode
			cnt++
		}
	}
	return cnt
}

// 修改指定下标上的节点数据
func UpdateNodeByIndex(data int, index int, head *Node)  {
	if index == 0 {
		// 修改头节点信息
		hNode.Data = data
	} else {
		// 通过遍历找到指定节点下标的值
		node := head
		cnt := 1
		for {
			if node.NextNode == nil {
				break
			} else {
				if cnt == index {
					node.NextNode.Data = data
					break
				} else {
					node = node.NextNode
					cnt++
				}
			}
		}
	}
}

// 删除指定下标的节点
func DeleteNodeByIndex(index int, head *Node) *Node {

	if index == 0 {
		hNode = head.NextNode
		return hNode
	} else {
		node := head
		cnt := 1

		for {
			if node.NextNode == nil {
				break
			} else {
				if cnt == index {
					node.NextNode = node.NextNode.NextNode
					// Go的垃圾回收机制会回收掉不用的内存，所以直接指向下下个内存地址即可
					break
				} else {
					node = node.NextNode
					cnt++
				}
			}
		}
		return nil
	}

}