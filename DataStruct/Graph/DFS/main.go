// DFS 深度优先算法
package main

import "fmt"

type Node struct {
	value    int //节点为int型
	searched bool
}

type Graph struct {
	nodes []*Node
	edges map[*Node][]*Node //邻接表示的无向图
}

//增加节点
//可以理解为Graph的成员函数
func (g *Graph) AddNode(n *Node) {
	g.nodes = append(g.nodes, n)
}

//增加边
func (g *Graph) AddEdge(u, v *Node) {
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

func (g *Graph) visitNode(n *Node) {
	for _, iNode := range g.edges[n] {
		if !iNode.searched {
			iNode.searched = true
			fmt.Printf("%v->", iNode.value)
			g.visitNode(iNode)
			return
		}
	}
}

func initGraph() Graph {
	g := Graph{}
	for i := 1; i <= 9; i++ {
		g.AddNode(&Node{i, false})
	}

	//生成边
	A := [...]int{1, 1, 2, 2, 2, 3, 4, 5, 5, 6, 1}
	B := [...]int{2, 5, 3, 4, 5, 4, 5, 6, 7, 8, 9}
	g.edges = make(map[*Node][]*Node) //初始化边
	for i := 0; i < 11; i++ {
		g.AddEdge(g.nodes[A[i]-1], g.nodes[B[i]-1])
	}
	return g
}

//func DFS(graph *Graph, used []int, x int, y int) {
//	if (graph )
//}

func main() {
	g := initGraph()
	g.Print()
	fmt.Println()

	//g.nodes[0].searched = true
	//fmt.Printf("%v->",g.nodes[0].value)
	//g.visitNode(g.nodes[0])

	g.DFS()
}
