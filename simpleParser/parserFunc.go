package main

import (
	"go/ast"
	"go/parser"
)

const src6 = `package foo
type xType struct {}
func (p *xType) Hello(arg1, arg2 int) (bool) { 
	return true
}
`

//只有参数和返回值才是函数签名部分
func ParserFunc() {
	file := FsetInit(parser.AllErrors, src6)
	for _, decl := range file.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok {
			ast.Print(nil, fn)
		}
	}
}
