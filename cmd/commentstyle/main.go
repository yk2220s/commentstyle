package main

import (
	"github.com/yk2220s/commentstyle"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(commentstyle.Analyzer) }
