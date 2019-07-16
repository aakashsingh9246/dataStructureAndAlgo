package main

import(
	"fmt"
)

type node struct{
	data int
	left *node
	right *node
}


func rightRotate(n *node) *node{
	node:= n.left
	n.left = node.right
	node.right = n
	return node
}

func leftRotate(n *node) *node{
	node := n.right
	n.right = node.left
	node.left=n
	return node
}

func leftRightRotate(n *node) *node{
	n.left = leftRotate(n.left)
	n = rightRotate(n)
	return n
}

func rightLeftRotate(n *node) *node{
	n.right = rightRotate(n.right)
	n = leftRotate(n)
	return n
}


func insertValue(v int, n *node) *node{
	if n==nil{
		n = &node{v,nil,nil}
	}
	if v<n.data{
		n.left = insertValue(v,n.left)
		dif:=getHeight(n.left)-getHeight(n.right)
		//fmt.Println(dif)
		if dif==2{
			if v<n.left.data {
				n=rightRotate(n)
			}else{
				n= leftRightRotate(n)
				inOrderTraversal(n)
			}
		}
	}
	if v>n.data{
		n.right = insertValue(v, n.right)
		dif:=getHeight(n.right)-getHeight(n.left)
		///fmt.Println(dif)
		if dif==2{
			if v>n.right.data {
				n=leftRotate(n)
			}else{
				n= rightLeftRotate(n)
			}
		}
	}
	fmt.Println(n.data)
	return n
}


func getHeight(n *node)int{
	if n==nil{
		return -1
	}
	lefth:=getHeight(n.left)
	righth:= getHeight(n.right)
	if lefth>righth{
		return lefth+1
	}else{
		return righth+1
	}
}

func inOrderTraversal(root *node){
	if root==nil{
		return
	}
	inOrderTraversal(root.left)
	fmt.Println(root.data)
	inOrderTraversal(root.right)
}

func deleteNode(v int, n *node) *node{
	if v==n.data{
		if n.left==nil && n.right==nil{
			n=nil
		}else if n.left==nil{
			n=n.right
		}else if n.right==nil{
			n=n.left
		}else{
			s:=nextInt(n.right)
			n.data = s
			deleteNode(s,n.right)
		}
	}else if v<n.data{
		n.left=deleteNode(v,n.left)
		dif:=getHeight(n.left)-getHeight(n.right)
		if dif==2{
			if v<n.left.data {
				n=rightRotate(n)
			}else{
				n= leftRightRotate(n)
			}
		}
	}else if v>n.data{
		n.right=deleteNode(v,n.right)
		dif:=getHeight(n.right)-getHeight(n.left)
		if dif==2{
			if v>n.right.data {
				n=leftRotate(n)
			}else{
				n= rightLeftRotate(n)
			}
		}
	}
	return n
}

func nextInt(n *node) int{
	if n.left !=nil{
		return nextInt(n.left)
	}
	return n.data
}

func main(){
	root:= &node{60,nil,nil}
	var arr = []int{15,20}
	for _,k:=range arr{
		insertValue(k,root)
	}
	inOrderTraversal(root)
}
