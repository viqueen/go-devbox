package scan_tasks

import (
	"github.com/fatih/color"
	"go/ast"
	"strings"
)

type Check string

const (
	CheckExec Check = "exec"
	CheckNet  Check = "net"
	CheckHttp Check = "http"
	CheckOs   Check = "os"
)

func ParseChecks(checks string) []Check {
	if checks == "" {
		return []Check{
			CheckExec,
			CheckOs,
			CheckHttp,
			CheckNet,
		}
	}
	var enabledChecks []Check
	for _, check := range strings.Split(checks, ",") {
		enabledChecks = append(enabledChecks, Check(check))
	}
	return enabledChecks
}

var checksFn = map[Check]func(node ast.Node) bool{
	CheckExec: checkForExecCommand,
	CheckNet:  checkForNetDial,
	CheckHttp: checkForHttpMethods,
	CheckOs:   checkForOsMethods,
}

var checksColor = map[Check]func(format string, args ...interface{}){
	CheckExec: color.Red,
	CheckNet:  color.HiMagenta,
	CheckHttp: color.Cyan,
	CheckOs:   color.HiRed,
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
