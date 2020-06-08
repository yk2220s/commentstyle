package main

import (
	"github.com/yk2220s/commentstyle"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(commentstyle.Analyzer) }
