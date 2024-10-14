package scan_tasks

import (
	"github.com/fatih/color"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

func ScanGoFile(file string) error {
	fileSet := token.NewFileSet()
	astFile, err := parser.ParseFile(fileSet, file, nil, parser.AllErrors)
	if err != nil {
		return err
	}
	execCommandFound := false
	netDialFound := false
	httpMethodFound := false
	osMethodFound := false
	ast.Inspect(astFile, func(node ast.Node) bool {
		if !execCommandFound && checkForExecCommand(node) {
			color.Red("found exec.Command in %s\n", file)
			execCommandFound = true
		}
		if !netDialFound && checkForNetDial(node) {
			color.Cyan("found net.Dial in %s\n", file)
			netDialFound = true
		}
		if !httpMethodFound && checkForHttpMethods(node) {
			color.Yellow("found http method in %s\n", file)
			httpMethodFound = true
		}
		if !osMethodFound && checkForOsMethods(node) {
			color.Magenta("found os method in %s\n", file)
			osMethodFound = true
		}
		return true
	})
	return nil
}

func checkForExecCommand(node ast.Node) bool {
	switch n := node.(type) {
	case *ast.CallExpr:
		switch fun := n.Fun.(type) {
		case *ast.SelectorExpr:
			switch x := fun.X.(type) {
			case *ast.Ident:
				if x.Name == "exec" && fun.Sel.Name == "Command" {
					return true
				}
			}
		}
	}
	return false
}

func checkForNetDial(node ast.Node) bool {
	switch n := node.(type) {
	case *ast.CallExpr:
		switch fun := n.Fun.(type) {
		case *ast.SelectorExpr:
			switch x := fun.X.(type) {
			case *ast.Ident:
				if x.Name == "net" && strings.HasPrefix(fun.Sel.Name, "Dial") {
					return true
				}
			}
		}
	}
	return false
}

func checkForHttpMethods(node ast.Node) bool {
	switch n := node.(type) {
	case *ast.CallExpr:
		switch fun := n.Fun.(type) {
		case *ast.SelectorExpr:
			switch x := fun.X.(type) {
			case *ast.Ident:
				if x.Name == "http" && (strings.HasPrefix(fun.Sel.Name, "Get") ||
					strings.HasPrefix(fun.Sel.Name, "Post") ||
					strings.HasPrefix(fun.Sel.Name, "NewRequest")) {
					return true
				}
			}
		}
	}
	return false
}

func checkForOsMethods(node ast.Node) bool {
	switch n := node.(type) {
	case *ast.CallExpr:
		switch fun := n.Fun.(type) {
		case *ast.SelectorExpr:
			switch x := fun.X.(type) {
			case *ast.Ident:
				if x.Name == "os" && (strings.HasPrefix(fun.Sel.Name, "Open") ||
					strings.HasPrefix(fun.Sel.Name, "Create")) {
					return true
				}
			}
		}
	}
	return false
}
