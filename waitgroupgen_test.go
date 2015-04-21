package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"testing"
)

func TestWaitGroupGen(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "example/foo.go", nil, 0)
	if err != nil {
		fmt.Println("bad")
		fmt.Println(err)
	} else {
		fmt.Println("Good!")
		fmt.Println(f)
		var buf bytes.Buffer
		printer.Fprint(&buf, fset, f)
		fmt.Println(buf.String())

		//funcsToParse := []*ast.FuncDecl{}

		for _, node := range f.Decls {
			fmt.Printf("you got type %T\n", node)
			switch n := node.(type) {
			default:
				fmt.Printf("unexpected type %T\n", n)
			case *ast.FuncDecl:
				fmt.Printf("you got the func decl %T\n", n)
				for _, stmt := range n.Body.List {
					switch s := stmt.(type) {
					case *ast.AssignStmt:
						fmt.Printf("length left %d\n", len(s.Lhs))
						fmt.Printf("right hand side type %T\n", s.Lhs[0])
						fmt.Printf("right hand side type %T\n", s.Rhs[0])

						switch ss := s.Rhs[0].(type) {
						case *ast.BasicLit:
							fmt.Printf("%T\n", ss)
							//if s.Lhs[0] == "_" && s.Rhs[0] {

							//}
						}
						var buf2 bytes.Buffer
						//printer.Fprint(&buf2, fset, s)
						printer.Fprint(&buf2, fset, s.Lhs[0])
						fmt.Println(buf2.String())
					}
				}
			}
		}
	}
}
