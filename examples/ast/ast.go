package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	// src is the input for which we want to print the AST.
	//	src := `
	//package main
	//func main() {
	//	println("Hello, World!")
	//}
	//`
	// Create the AST by parsing src.
	set := token.NewFileSet() // positions are relative to set
	//f, err := parser.ParseDir(set, ".", nil, 0)
	f, err := parser.ParseFile(set, "../../complex/image.go", nil, 0)
	if err != nil {
		panic(err)
	}

	// Print the AST.
	err = ast.Print(set, f)
	if err != nil {
		return
	}
}
