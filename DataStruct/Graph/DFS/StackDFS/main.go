package main

import "fmt"

type Node struct {
	Next *Node
	Val  int
}

type Stack struct {
	top    *Node
	length int
}

func CreateStackNode(value int) *Node {
	return &Node{
		nil,
		value,
	}
}

func CreateStack() *Stack {
	// 下标0 是空Node，所以Top的时候是 Stack.top.Next.Val
	return &Stack{
		&Node{
			nil,
			0,
		},
		0,
	}
}

func (stack *Stack) IsStackEmpty() bool {
	return stack.top.Next == nil
}

func (stack *Stack) Push(value int) {
	p := CreateStackNode(value)
	p.Next = stack.top.Next
	stack.top.Next = p
	stack.length++
}

func (stack *Stack) Pop() int {
	if stack.top.Next == nil {
		return -999
	}
	p := stack.top.Next
	value := p.Val
	stack.top.Next = p.Next
	p = nil
	stack.length--
	return value
}

func (stack *Stack) Top() int {
	if stack.top.Next == nil {
		return -999
	}
	return stack.top.Next.Val
}

func (stack *Stack) UpsideDown() {
	// 倒着输出
	if stack.top.Next == nil {
		fmt.Println("This is Empty Stack")
		return
	}
	p := stack.top.Next
	for p != nil {
		fmt.Printf("%d ", p.Val)
		p = p.Next
	}
}

type GraphNode struct {
	value    int //节点为int型
	searched bool
}

type Graph struct {
	nodes []*GraphNode                // 所有的节点
	edges map[*GraphNode][]*GraphNode //邻接表示的无向图
}

//增加节点
//可以理解为Graph的成员函数
func (g *Graph) AddNode(n *GraphNode) {
	g.nodes = append(g.nodes, n)
}

//增加边
func (g *Graph) AddEdge(u, v *GraphNode) {
	g.edges[u] = append(g.edges[u], v) //u->v边
	g.edges[v] = append(g.edges[v], u) //u->v边
}

//打印图
func (g *Graph) Print() {
	//range遍历 g.nodes，返回索引和值
	for _, iNode := range g.nodes {
		fmt.Printf("%v:", iNode.value)
		for _, next := range g.edges[iNode] {
			fmt.Printf("%v->", next.value)
		}
		fmt.Printf("\n")
	}
}

func initGraph() Graph {
	g := Graph{}
	for i := 1; i <= 9; i++ {
		g.AddNode(&GraphNode{i, false})
	}

	//生成边
	A := [...]int{1, 1, 2, 2, 2, 3, 4, 5, 5, 6, 1}
	B := [...]int{2, 5, 3, 4, 5, 4, 5, 6, 7, 8, 9}
	g.edges = make(map[*GraphNode][]*GraphNode) //初始化边
	for i := 0; i < 11; i++ {
		g.AddEdge(g.nodes[A[i]-1], g.nodes[B[i]-1])
	}
	return g
}

func (g *Graph) DFS() {
	for _, iNode := range g.nodes {
		if !iNode.searched {
			iNode.searched = true
			fmt.Printf("%v->", iNode.value)
			g.visitNode(iNode)
			fmt.Printf("\n")
			g.DFS()
		}
	}
}

func (g *Graph) visitNode(n *GraphNode) {
	for _, iNode := range g.edges[n] {
		if !iNode.searched {
			iNode.searched = true
			fmt.Printf("%v->", iNode.value)
			g.visitNode(iNode)
			return
		}
	}
}

func (g *Graph) huifenvisitNode(n *GraphNode) {
	for _, iNode := range g.edges[n] {
		if !iNode.searched {
			iNode.searched = true
			fmt.Printf("%v->", iNode.value)
			g.huifenvisitNode(iNode)
			return
		}
	}
}

func (g *Graph) huifenDFS() {
	for _, iNode := range g.nodes {
		if !iNode.searched {
			iNode.searched = true
			fmt.Printf("%v->", iNode.value)
			g.huifenvisitNode(iNode)
			fmt.Printf("\n")
			g.huifenDFS()
		}
	}
}

func (g *Graph) StackNodeDFS(start int) {
	aStack := CreateStack()
	aStack.Push(start)

}

func StackDFS(start int, maze [][]int, visited []int, N int) {
	aStack := CreateStack()
	aStack.Push(start)
	visited[start] = 1
	is_push := false
	for !aStack.IsStackEmpty() {
		is_push = false
		v := aStack.Top()
		for i := 0; i < N; i++ {
			if maze[v][i] == 1 && visited[i+1] == 0 {
				visited[i+1] = 1
				aStack.Push(i)
				is_push = true
				break
			}
		}
		if !is_push {
			fmt.Println(v)
			aStack.Pop()
		}
	}
}

func main() {
	g := initGraph()
	g.Print()

	N := 5
	maze := [][]int{
		{0, 1, 1, 0, 0},
		{0, 0, 1, 0, 1},
		{0, 0, 1, 0, 0},
		{1, 1, 0, 0, 1},
		{0, 0, 1, 0, 0},
	}
	visited := make([]int, N+1)
	fmt.Println(maze)
	fmt.Println(visited)
	for i := 0; i < N; i++ {
		if visited[i] == 1 {
			continue
		}
		StackDFS(i, maze, visited, N)
	}
}
