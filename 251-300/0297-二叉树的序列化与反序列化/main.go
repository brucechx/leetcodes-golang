package _297_二叉树的序列化与反序列化

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val int
	Left *TreeNode
 	Right *TreeNode
}

type Codec struct {

}

// 前序遍历
func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	return serialize(root, "")
}

func serialize(root *TreeNode, str string) string{
	if root == nil{
		str += "null,"
	}else {
		str += strconv.Itoa(root.Val)
		str = serialize(root.Left, str)
		str = serialize(root.Right, str)
	}
	return str
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	datas := strings.Split(data, ",")
	return deserialize(datas)
}

func deserialize(data []string) *TreeNode{
	if len(data) == 0{
		return nil
	}
	if data[0] == "null"{
		data = data[1:]
		return nil
	}
	val, _ := strconv.Atoi(data[0])
	root := &TreeNode{
		Val:  val,
		Left:  nil,
		Right: nil,
	}
	data = data[1:]
	root.Left = deserialize(data)
	root.Right = deserialize(data)
	return root
}


/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
