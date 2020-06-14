// Package commentstyle provides an analyzer that helps to keep your comment style.
package commentstyle

import (
	"go/ast"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "commentstyle is linter to check comment style"

var (
	fPreferLineStyle bool = true
	fOnlyASCII       bool = true
)

func init() {
	Analyzer.Flags.BoolVar(&fPreferLineStyle, "prefer-line-style", true, "enable prefer-line-style style check")
	Analyzer.Flags.BoolVar(&fOnlyASCII, "only-ascii", true, "enable only-ascii style check")
}

var Analyzer = &analysis.Analyzer{
	Name: "commentstyle",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.File)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.File:
			for _, cg := range n.Comments {
				for _, cmt := range cg.List {
					if fOnlyASCII {
						checkASCII(pass, cmt)
					}

					if fPreferLineStyle {
						checkLineStyle(pass, cmt)
					}
				}
			}
		}
	})

	return nil, nil
}

func checkASCII(pass *analysis.Pass, cmt *ast.Comment) {
	if r, isASCII := isASCII(cmt.Text); !isASCII {
		pass.Reportf(cmt.Pos(), "[only-ascii] %#U is non-ASCII character", r)
	}
}

func isASCII(s string) (rune, bool) {
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		if r[i] > unicode.MaxASCII {
			return r[i], false
		}
	}

	return 0, true
}

func checkLineStyle(pass *analysis.Pass, cmt *ast.Comment) {
	if isBlock := isBlockStyle(cmt.Text); isBlock {
		pass.Reportf(cmt.Pos(), "[prefer-line-style] should use //-style")
	}
}

func isBlockStyle(s string) bool {
	return strings.HasPrefix(s, "/*")
}
