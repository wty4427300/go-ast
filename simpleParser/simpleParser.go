package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strconv"
)

//类型解析
func main()  {
	//expr, _ := go-ast.ParseExpr(`9527`)
	//ast.Print(nil,expr)
	//newIdent()
	//parserIdent()
	//Binary()
	//PackageDome()
	//parserImport()
	//parserType()
	//parserConst()
	parserVar()
	//fmt.Println("vscode")
}

func newIdent() {
	ast.Print(nil, ast.NewIdent(`x`))
}

func parserIdent() {
	expr, _ := parser.ParseExpr(`1234`)
	ast.Print(nil, expr)
}

func fsetInit(mode parser.Mode,src interface{}) *ast.File {
	fset:=token.NewFileSet()
	file, err := parser.ParseFile(fset, "hello.go", src, mode)
	if err!=nil {
		log.Fatal(err)
	}
	return file
}

//二元表达式的解析
func Binary() {
	expr, err := parser.ParseExpr(`xyz`)
	if err!=nil {
		log.Fatal(err)
	}
	fmt.Println("计算的结果",Eval(expr, map[string]float64{
		"x":100,
		"xyz":1200,
	}))
}
//手工对表达式求值
//这个map用来存储临时变量
func Eval(exp ast.Expr,vars map[string]float64) float64{
	switch exp:=exp.(type) {
	case *ast.BinaryExpr:
		//如果是二元表达式则需要再分析
		return EvalBinaryExpr(exp,vars)
	case *ast.BasicLit:
		//如果是面值则返回对应的值
		float, err := strconv.ParseFloat(exp.Value, 64)
		if err!=nil {
			log.Fatal(err)
		}
		return float
	case *ast.Ident:
		return vars[exp.Name]
	}
	return 0
}

//处理add和mul的二元表达式
func EvalBinaryExpr(exp *ast.BinaryExpr,vars map[string]float64) float64{
	switch exp.Op {
	case token.ADD:
		return Eval(exp.X,vars)+Eval(exp.Y,vars)
	case token.MUL:
		return Eval(exp.X,vars)*Eval(exp.Y,vars)
	case token.ASSIGN:
		fmt.Print("傻逼本parser不支持这种表达式")
	}
	return 0
}

const src=`package pkgname

import ("a"; "b")
type SomeType int
const PI = 3.14
var Length = 1

func main() {}
`
func PackageDome() {
	fset:=token.NewFileSet()
	f,err:=parser.ParseFile(fset,"hello.go",src,parser.AllErrors)
	if err!=nil {
		fmt.Println(err)
		return
	}
	//输出包名
	fmt.Println(f.Name)
	//输出导入的三方包
	for _,s:=range f.Imports{
		fmt.Println("import:",s.Path.Value)
	}
	//当前文件全部的包级信息
	//decls包含import,type,var,const四种类型
	//所以遍历的时候需要做类型判断
	for _,v:=range f.Decls{
		decl,ok := v.(*ast.GenDecl);
		//判断decl的类型为import
		if ok&&decl.Tok==token.IMPORT{
			for _,v:=range decl.Specs{
				//输出decl
				fmt.Println("import:",v.(*ast.ImportSpec).Path.Value)
			}
		}
	}
	//自定义语法树输出
	ast.Walk(new(myNodeVisitor),f)
	ast.Inspect(f, func(node ast.Node) bool {
		x,ok:=node.(*ast.Ident)
		if ok {
			fmt.Println("ast.Inspect",x.Name)
		}
		return true
	})
}

type myNodeVisitor struct {}
//实现了ast.Visitor的Visit接口
func (p *myNodeVisitor) Visit(n ast.Node) (w ast.Visitor) {
	x,ok:=n.(*ast.Ident)
	if ok {
		fmt.Println("myNodeVisitor.Visitor:",x.Name)
	}
	return p
}
const src1 = `package foo
import "pkg-a"
import pkg_b_v2 "pkg-b"
import . "pkg-c"
import _ "pkg-d"
`
func parserImport()  {
	fset := token.NewFileSet()
	//最后一个参数是指定分析模式，该参数下只分析import相关
	file, err := parser.ParseFile(fset, "hello.go", src1, parser.ImportsOnly)
	if err!=nil{
		log.Fatal(err)
	}
	for _,s:=range file.Imports{
		fmt.Printf("import:name = %v,path=%#v\n\n", s.Name, s.Path)
	}
}
const src2 = `package foo
type MyInt1 int
type MyInt2 = int
`
func parserType(){
	fset:=token.NewFileSet()
	file, err := parser.ParseFile(fset, "hello.go", src2, parser.AllErrors)
	if err!=nil {
		log.Fatal(err)
	}
	for _,decl:=range file.Decls{
		genDecl,ok := decl.(*ast.GenDecl)
		if ok{
			for _,spec:=range genDecl.Specs{
				fmt.Println("type相关信息",spec)
			}
		}
	}
}

const src3 =`package foo
const Pi = 3.14
const E float64 = 2.71828`
func parserConst(){
	file := fsetInit(parser.AllErrors,src3)
	for _,decl:=range file.Decls{
		if genDecl,ok := decl.(*ast.GenDecl);ok{
			for _,spec:=range genDecl.Specs{
				fmt.Println("const相关信息",spec)
			}
		}
	}
}
const src4 = `package foo
var Pi = 3.14
var a int =1
`
const src5=`package foo
var (
	a int
	b bool
)
`
//变量的解析和常亮差不多,import,type,var,const都在ast.decls
func parserVar(){
	file := fsetInit(parser.AllErrors,src5)
	for _,decl:=range file.Decls{
		if genDecl,ok := decl.(*ast.GenDecl);ok{
			for _,spec:=range genDecl.Specs{
				fmt.Println("var相关信息",spec)
				ast.Print(nil, spec)
			}
		}
	}
}






