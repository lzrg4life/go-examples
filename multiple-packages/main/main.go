package main

import (
	"fmt"
	"main/bar"
	"main/foo"
)

func main() {
	fmt.Println("Hello World!")

	fmt.Println(foo.Test())
	fmt.Println(bar.Test())
}
