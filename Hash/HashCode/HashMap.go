package HashCode

import "robin/Hash/LinkNode"

var buckets[16] *LinkNode.Node

// 创建buckets中的每个元素
func CreateBuckets() {
	for i:=0; i<16; i++ {
		// 为每个bucket数组元素，添加一个头节点对象
		buckets[i] = LinkNode.CreateNode()
	}
}


// 自行编写简单的Hash散列算法
func HashCode(key string) int {
	// 此hash散列算法，将不同长度的key散列成[0,15]的整数
	sum := 0
	for i:=0; i<len(key); i++ {
		sum += int(key[i])
	}
	return sum%16
	/*
	var index int = 0
	index = int(key[0])
	for k := 0; k < len(key); k++ {
		index *= (1103515245 + int(key[k]))
	}
	index >>= 27
	index &= 16 - 1
	return index*/
}


// 添加键值对
func AddKeyValue(key string, value string) {
	// 通过Hash散列，将key换算成0-15的整数
	pos := HashCode(key)

	// 获取对应下标下的buckets中的头结点
	bucketsHeadNode := buckets[pos]

	// 每次都遍历到尾节点，确保每次都是在尾节点处添加新节点
	bucketsHeadNode = LinkNode.GetTailNode(bucketsHeadNode)

	// 向此头节点上添加新节点
	var kv = LinkNode.KV{key, value}
	LinkNode.AddNode(kv, bucketsHeadNode)
}

// 按键取值
func GetValueByKey(key string) {
	var pos = HashCode(key)

	var bucketsHeadNode = buckets[pos]

	// 链表的遍历
	LinkNode.ShowNode(key, bucketsHeadNode)
}