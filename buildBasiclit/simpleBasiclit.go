package main

import (
	"go/ast"
	"go/token"
)
//构造面值
func main()  {
	var lit9527=&ast.BasicLit{
		Kind: token.INT,
		Value: "9527",
	}
	ast.Print(nil,lit9527)
}
