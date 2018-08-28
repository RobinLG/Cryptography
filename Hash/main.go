package main

import "robin/Hash/HashCode"

func main() {

	HashCode.CreateBuckets()
	HashCode.AddKeyValue("aabbcc", "hello")
	HashCode.GetValueByKey("aabbcc")

	HashCode.AddKeyValue("a1", "hi")
	HashCode.GetValueByKey("a1")
}
