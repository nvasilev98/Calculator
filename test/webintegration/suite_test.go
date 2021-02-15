package webintegration_test

import (
	"flag"
	"testing"
	"time"

	. "github.com/nvasilev98/calculator/test/webintegration/app"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const calcPath = "github.com/nvasilev98/calculator/cmd/webapp"
const host = "127.0.0.1"
const port = "4043"
const username = "user"
const password = "password"

var calcApp *ExecutableApp
var calcExecutable *Executable

var _ = BeforeSuite(func() {
	calcApp = NewExecutableApp(calcPath)
	calcExecutable = waitUntilCalculatorStarts()
})

var _ = AfterSuite(func() {
	Expect(calcExecutable.Kill()).To(Succeed())
})

var integrationSuiteRun = flag.Bool("webintegration", false, "should webintegration tests run?")

func TestWebintegration(t *testing.T) {
	flag.Parse()
	RegisterFailHandler(Fail)
	if !*integrationSuiteRun {
		t.Skip("skip webintegration tests")
		return
	}
	RunSpecs(t, "Webintegration Suite")
}

func startCalculator() *Executable {
	calculatorExecutable, err := calcApp.Start(GinkgoWriter, GinkgoWriter, map[string]string{
		"HOST":     host,
		"PORT":     port,
		"USERNAME": username,
		"PASSWORD": password,
	})
	Expect(err).ToNot(HaveOccurred())
	return calculatorExecutable
}
func waitUntilCalculatorStarts() *Executable {
	brokerExecutable := startCalculator()
	Eventually(brokerExecutable.Netcat, 10*time.Second).Should(Succeed())
	return brokerExecutable
}
