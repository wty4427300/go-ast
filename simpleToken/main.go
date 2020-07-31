package main

import (
	"fmt"
	"go/scanner"
	"go/token"
)

func main()  {
	var src =[]byte("println(\"你好，世界\")")
	var fset=token.NewFileSet()
	var file=fset.AddFile("hello.go",fset.Base(),len(src))

	var s scanner.Scanner
	s.Init(file,src,nil,scanner.ScanComments)

	for  {
		pos,tok,lit:=s.Scan()
		if tok==token.EOF{
			break
		}
		fmt.Printf("token的位置:%s\ttoken的值:%s\ttoken的源代码文本:%q\n", fset.Position(pos), tok, lit)
		//Position()将偏移量转换为带有
	}
	//初始化一个position的
	pos:=token.Position{Filename: "文件名",Column: 1,Line: 1,Offset: 2}
	fmt.Print(pos.Filename)
}
