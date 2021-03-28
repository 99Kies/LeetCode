package main

import (
	"fmt"
	"os"
)

type Emp struct {
	Id   int
	Name string
	Next *Emp
}

func (this *Emp) ShowEmp() {
	fmt.Printf("链表%d id: %d name: %s\n", this.Id%7, this.Id, this.Name)
}

type EmpLink struct {
	Head *Emp
}

func (this *EmpLink) Insert(emp *Emp) {
	cur := this.Head
	var pre *Emp = nil
	if cur == nil {
		this.Head = emp
		return
	}

	for cur != nil {
		if cur.Id > emp.Id {
			break
		}
		pre = cur
		cur = cur.Next
	}
	pre.Next = emp
	emp.Next = cur

}

func (this *EmpLink) ShowLink(index int) {
	//fmt.Printf("%d ", this.Head)
	if this.Head == nil {
		fmt.Printf("链表 %d 为 nil\n", index)
	}
	cur := this.Head
	for cur != nil {
		fmt.Printf("链表%d 雇员=%d 名字=%s ->", index, cur.Id, cur.Name)
		cur = cur.Next
	}
	fmt.Println()
}

func (this *EmpLink) FindById(id int) *Emp {
	cur := this.Head
	for cur != nil {
		if cur.Id == id {
			return cur
		}
		cur = cur.Next
	}
	return nil
}

func (this *EmpLink) DeleteById(id int) {
	cur := this.Head
	if cur.Id == id {
		this.Head = cur.Next
	} else {
		for cur.Next != nil {
			if cur.Next.Id == id {
				cur.Next.Next = cur.Next
			}
			cur.Next = cur.Next.Next
		}
	}
}

type HashTable struct {
	LinkArr [7]EmpLink
}

func (this *HashTable) Insert(emp *Emp) {
	linkNum := this.HashFun(emp.Id)
	this.LinkArr[linkNum].Insert(emp)
}

func (this *HashTable) HashFun(id int) int {
	return id % 7
}

func (this *HashTable) ShowAll() {
	for i := 0; i < len(this.LinkArr); i++ {
		this.LinkArr[i].ShowLink(i)
	}
}

func (this *HashTable) FindById(id int) *Emp {
	linkindex := this.HashFun(id)
	return this.LinkArr[linkindex].FindById(id)
}

func (this *HashTable) FindEmpById(id int) *Emp {
	linkindex := this.HashFun(id)
	curLink := this.LinkArr[linkindex].Head
	for curLink != nil {
		if curLink.Id == id {
			return curLink
		}
		curLink = curLink.Next
	}
	return nil
}

func (this *HashTable) DeleteById(id int) {
	linkindex := this.HashFun(id)
	this.LinkArr[linkindex].DeleteById(id)
}

func (this *HashTable) ChangeEmpById(id int, name string) {
	emp := this.FindById(id)
	emp.Name = name
}

func main() {
	key := ""
	id := 0
	name := ""
	var hashtable HashTable
	for {
		fmt.Println("==============菜单===========")
		fmt.Println("insert 添加雇员")
		fmt.Println("show 显示雇员")
		fmt.Println("delete 解雇雇员")
		fmt.Println("find 查找雇员")
		fmt.Println("change 修改雇员信息")
		fmt.Println("exit 退出")
		fmt.Println("Please input your cmd: ")
		fmt.Scanln(&key)
		switch key {
		case "insert":
			fmt.Println("输入雇员id")
			fmt.Scanln(&id)
			fmt.Println("输入雇员name")
			fmt.Scanln(&name)
			emp := &Emp{
				Id:   id,
				Name: name,
				//Next: nil,
			}
			hashtable.Insert(emp)
		case "find":
			fmt.Println("请输入需要查找的id")
			fmt.Scanln(&id)
			emp := hashtable.FindById(id)
			if emp != nil {
				emp.ShowEmp()
			} else {
				fmt.Printf("id=%d 的雇员不存在.", id)
			}
		case "delete":
			fmt.Println("请输入需要删除的id")
			fmt.Scanln(&id)
			hashtable.DeleteById(id)
		case "change":
			fmt.Println("请输入需要修改的id位置")
			fmt.Scanln(&id)
			fmt.Println("讲这个id的名字修改为：")
			fmt.Scanln(&name)
			hashtable.ChangeEmpById(id, name)
		case "show":
			hashtable.ShowAll()
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("输入错误")
		}

	}
}
