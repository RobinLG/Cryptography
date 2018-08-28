package LinkNode

import "fmt"

// 通过头插法	，完成链表基本功能

type Node struct {
	Data int
	NextNode *Node
}

// 通过全局变量，记录当前头节点
var hNode *Node

// 创建头节点
func CreateHead(data int) *Node {
	node := &Node{data, nil}
	hNode = node
	return hNode
}

// 添加新节点
func AddNode(data int) *Node {
	newNode := &Node{data, nil}
	newNode.NextNode = hNode
	hNode = newNode
	return newNode
}

// 链表的遍历
func ShowNode() {
	node := hNode
	for {
		fmt.Println("节点信息：", node.Data)
		if node.NextNode == nil {
			break
		} else {
			node = node.NextNode
		}
	}
}