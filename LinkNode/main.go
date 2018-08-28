package main

import "robin/LinkNode/LinkNode"

func main()  {
	LinkNode.CreateHead(1)
	LinkNode.AddNode(2)
	LinkNode.AddNode(3)
	LinkNode.AddNode(4)

	LinkNode.ShowNode()
}
