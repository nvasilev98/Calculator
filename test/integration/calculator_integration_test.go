package integration_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Calculator Integration Tests", func() {

	It("successfully calculates with precede operator", func() {
		res := calculatorPO.CalculateSuccessfully("2+4*3/2-1-14")
		Expect(res).To(Equal(float64(-7)))
	})
	It("successfully calculates with brackets", func() {
		res := calculatorPO.CalculateSuccessfully("(32/2 +20)/3-(2+3)*1.2")
		Expect(res).To(Equal(float64(6)))
	})
	It("fails to calculate with illegal symbol", func() {
		calculatorPO.CalculateFails("(a+12)-3/2")
	})
	It("fails to calculate with mismatched brackets", func() {
		calculatorPO.CalculateFails(")1+1)/2")
	})
	It("fails to calculate with more operators", func() {
		calculatorPO.CalculateFails("2*/3+3-2")

	})
})
