package integration_test

import (
	"flag"
	"testing"

	. "github.com/nvasilev98/calculator/test/integration/fixture"
	. "github.com/nvasilev98/calculator/test/integration/util/app"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	mainPath = "github.com/nvasilev98/calculator/cmd/consoleapp"
)

var integrationSuiteRun = flag.Bool("integration", false, "should integration tests run?")

func TestIntegration(t *testing.T) {
	flag.Parse()
	RegisterFailHandler(Fail)
	if !*integrationSuiteRun {
		t.Skip("skip integration tests")
		return
	}
	RunSpecs(t, "Integration Suite")
}

var calculatorPO *CalculatorPageObject
var _ = BeforeSuite(func() {
	path, err := BuildApplication(mainPath)
	Expect(err).ToNot(HaveOccurred())
	calculatorPO = NewCalculatorPageObject(path)
})
