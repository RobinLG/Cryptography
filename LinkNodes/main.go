package main

import (
	"fmt"
	"robin/LinkNodes/LinkNodes"
)

func main() {
	fmt.Println("HelloWorld")

	// 调用CreateHeadNode创建头节点
	head := LinkNodes.CreateHeadNode(1)
	// 添加新节点
	node := LinkNodes.AddNode(2, head)
	node = LinkNodes.AddNode(3, node)
	node = LinkNodes.AddNode(4, node)

	//head = LinkNodes.InsertNodeWithIndex(100, 0, head)

	//LinkNodes.InsertNodeWithIndex(100, 2, head)

	//LinkNodes.UpdateNodeByIndex(100, 0, head)

	head = LinkNodes.DeleteNodeByIndex(0, head)

	LinkNodes.ShowNodes(head)
	// fmt.Println(head.Data)
	// fmt.Println("链表总长度为：", LinkNodes.NLen(head))
}

