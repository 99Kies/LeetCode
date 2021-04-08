package main

import "fmt"

func hanio(n int, x string, y string, z string) {
	if n < 1 {
		fmt.Println("汉诺塔的层数不能小于1")
	} else if n == 1 {
		fmt.Println("移动：" + x + "——>" + z)
	} else {
		hanio(n-1, x, z, y)
		fmt.Println("移动：" + x + "——>" + z)
		hanio(n-1, y, x, z)
	}
}

func main() {
	var x = "X"
	var y = "Y"
	var z = "Z"
	hanio(3, x, y, z)
}
