package main

import "sync"

type NodeB struct {
	key int
	value Item
	left *NodeB
	right *NodeB
}

type ItembinarySearchTree struct {
	root *NodeB
	lock sync.RWMutex
}

func (bst *ItembinarySearchTree) Insert(key int, value Item){
	bst.lock.Lock()
	defer bst.lock.Unlock()
	n := &NodeB{key, value, nil, nil}
	if bst.root == nil{
		bst.root = n
	}else{
		insertNode(bst.root, n)
	}
}

func insertNode(node, newNode *NodeB){
	if newNode.key < node.key{
		if node.left == nil{
			node.left = newNode
		}else{
			insertNode(node.left, newNode)
		}
	}else{
		if node.right == nil{
			node.right = newNode
		}else{
			insertNode(node.right, newNode)
		}
	}
}

func preOrderTraverse(n *NodeB, f func(Item)){
	if n != nil{
		f(n.value)
		preOrderTraverse(n.left, f)
		preOrderTraverse(n.right, f)
	}
}