package main

import (
	"fmt"
	"strings"
)

func interpret(command string) string {
	command = strings.ReplaceAll(command, "(al)", "al")
	return strings.ReplaceAll(command, "()", "o")
}

func main() {
	fmt.Println(interpret("G()()()()()(al)(al)"))
}
