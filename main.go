package main

import (
	"fmt"

	"github.com/rselbach/testmod"
	testmodML "github.com/rselbach/testmod/v2"
)

func main() {
	// 引用同一个外部包的不同版本
	fmt.Println(testmod.Hi("Roberto"))
	g, err := testmodML.Hi("Roberto", "pt")
	if err != nil {
		panic(err)
	}
	fmt.Println(g)
}
