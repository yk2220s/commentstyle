package commentstyle_test

import (
	"testing"

	"github.com/yk2220s/commentstyle"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestCheckASCII(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), commentstyle.Analyzer, "ascii")
}

func TestCheckLineStyle(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), commentstyle.Analyzer, "linestyle")
}
