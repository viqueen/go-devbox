package scan_tasks

import (
	"github.com/fatih/color"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

func ScanGoFile(file string, enabledChecks []Check) error {
	fileSet := token.NewFileSet()
	astFile, err := parser.ParseFile(fileSet, file, nil, parser.AllErrors)
	if err != nil {
		return err
	}
	checkFound := make(map[Check]bool)
	ast.Inspect(astFile, func(node ast.Node) bool {
		for _, check := range enabledChecks {
			if !checkFound[check] && checksFn[check](node) {
				checksColor[check]("found %s in %s\n", check, file)
				checkFound[check] = true
			}
		}
		if checkForConstant(node, "/etc/passwd") {
			color.Green("found password in %s\n", file)
		}
		if checkForConstant(node, "/etc/hosts") {
			color.Green("found hosts in %s\n", file)
		}
		return true
	})
	return nil
}

func checkForConstant(node ast.Node, target string) bool {
	switch n := node.(type) {
	case *ast.BasicLit:
		if n.Kind == token.STRING && strings.Contains(n.Value, target) {
			return true
		}
	}
	return false
}
