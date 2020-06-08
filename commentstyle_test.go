package commentstyle_test

import (
	"testing"

	"github.com/yk2220s/commentstyle"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, commentstyle.Analyzer, "a")
}
