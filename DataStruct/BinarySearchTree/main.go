// 二叉排序树
package main

import "fmt"

type BST struct {
	Val   int
	Left  *BST
	Right *BST
}

func (this *BST) NewNode(value int) *BST {
	cur := &BST{Val: value}
	return cur
}

func (this *BST) Search(value int) bool {
	if this == nil {
		return false
	}
	compare := value - this.Val
	if compare < 0 {
		return this.Left.Search(value)
	} else if compare > 0 {
		return this.Right.Search(value)
	} else {
		return true
	}
}

func (this *BST) Insert(value int) *BST {
	// 左侧的节点值小于右侧的节点值
	if this == nil {
		newNode := &BST{Val: value}
		return newNode
	}
	if value < this.Val {
		// 插入的值比当前节点的小，那就插到左侧
		this.Left = this.Left.Insert(value)
	} else {
		// 插入的值比当前的值哒，那就插到右侧
		this.Right = this.Right.Insert(value)
	}
	return this
	// 递归出口 类似于后根遍历
}

func (this *BST) Delete(value int) *BST {
	if this == nil {
		return nil
	}
	if value < this.Val {
		// 说明value在Left
		this.Left = this.Left.Delete(value)
	} else if value > this.Val {
		// 说明value在Right
		this.Right = this.Right.Delete(value)
	} else {
		// 找到节点，删除节点。
		if this.Left != nil && this.Right != nil {
			// 找到的这个节点既有左子树又有右子树。
			this.Val = this.Right.getMin()
			this.Right = this.Delete(this.Val)
		} else if this.Left != nil {
			// 找到的节点只有左子树
			this = this.Left
		} else {
			// 找到的只有右子树
			this = this.Right
		}
	}
	return this
}

//按顺序获得树中元素
func (this *BST) InOrder(node *BST) {
	if node != nil {
		this.InOrder(node.Left)
		fmt.Printf("%d ", node.Val)
		this.InOrder(node.Right)
	}
}

func (this *BST) getAll() []int {
	values := []int{}
	return addValues(values, this)
}

//将一个节点加入切片中
func addValues(values []int, t *BST) []int {
	if t != nil {
		values = addValues(values, t.Left)
		values = append(values, t.Val)
		values = addValues(values, t.Right)
	}
	return values
}

func (this *BST) getMin() int {
	// 左 > 中 > 右
	if this == nil {
		return -1
	}
	if this.Left == nil {
		return this.Val
	} else {
		return this.Left.getMin()
	}
}

func (this *BST) getMax() int {
	if this == nil {
		return -1
	}
	if this.Right == nil {
		return this.Val
	} else {
		return this.Right.getMax()
	}
}

//查找最小结点
func (this *BST) getMinNode() *BST {
	if this == nil {
		return nil
	} else {
		for this.Left != nil {
			this = this.Left
		}
	}
	return this
}

//查找最大结点
func (this *BST) getMaxNode() *BST {
	if this == nil {
		return nil
	} else {
		for this.Right != nil {
			this = this.Right
		}
	}
	return this
}

func main() {
	bsTree := &BST{
		Val: 22,
	}
	bsTree.Insert(12)
	bsTree.Insert(33)
	bsTree.Insert(22)
	bsTree.Insert(12)
	bsTree.Insert(122)
	bsTree.Insert(62)
	bsTree.Insert(72)
	bsTree.Insert(25)
	bsTree.Insert(32)
	bsTree.Insert(42)
	fmt.Println(bsTree.getAll())
	fmt.Println(bsTree.getMin())
	fmt.Println(bsTree.getMax())

	bsTree.InOrder(bsTree)
}
