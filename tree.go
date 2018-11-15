package main

import "fmt"



type Node struct {
	data int
	left *Node
	right *Node
}
type Tree struct {
	onePath []int
	allPath [][]int
}

func (t *Tree) GetPaths(node *Node)  {
	t.onePath = append(t.onePath,node.data)
	if node.left != nil {
		t.GetPaths(node.left)
	}
	if node.right != nil {
		t.GetPaths(node.right)
	}
	if node.left == nil && node.right == nil {
		onePath := make([]int , len(t.onePath))
		copy(onePath,t.onePath)
		t.allPath = append(t.allPath,onePath)
	}
	t.onePath = t.onePath[:len(t.onePath)-1]
}

/**            10
               / \
              5  12
             / \
            4  7
 */
func main()  {
	node10 := Node{data:10}
	node5 := Node{data:5}
	node4 := Node{data:4}
	node7 := Node{data:7}
	node12 := Node{data:12}
	node10.left = &node5
	node10.right = &node12
	node5.left = &node4
	node5.right = &node7
	tree := new(Tree)
	tree.GetPaths(&node10)
	fmt.Println(tree.allPath)

}
