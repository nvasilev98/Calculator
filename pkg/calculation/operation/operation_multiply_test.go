package operation_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/nvasilev98/calculator/pkg/calculation/operation"
)

var _ = Describe("Operation Multiply Unit Tests", func() {
	var (
		newMultiply Multiply
	)
	Describe("Creating Operation Multiply", func() {
		It("multiplies two numbers correctly", func() {
			res, err := newMultiply.Evaluate(10, 2)
			Expect(err).To(BeNil())
			Expect(res).To(Equal(float64(20)))
		})
		It("multiplies two bigger numbers correctly", func() {
			res, err := newMultiply.Evaluate(13, 10)
			Expect(err).To(BeNil())
			Expect(res).To(Equal(float64(130)))
		})
	})
	Describe("Getting Weight of Operation", func() {
		It("weights operation correctly", func() {
			Expect(newMultiply.Weight()).To(Equal(2))
		})
	})
})
