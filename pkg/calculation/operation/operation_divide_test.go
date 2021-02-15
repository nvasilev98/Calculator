package operation_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/nvasilev98/calculator/pkg/calculation/operation"
)

var _ = Describe("Operation Divide  Unit Tests", func() {
	var (
		newDivide Divide
	)
	Describe("Creating Operation Divide", func() {
		It("divides two numbers correctly", func() {
			res, err := newDivide.Evaluate(10, 2)
			Expect(err).To(BeNil())
			Expect(res).To(Equal(float64(5)))
		})
		It("returns error if division by zero", func() {
			_, err := newDivide.Evaluate(10, 0)
			Expect(err).To(HaveOccurred())
			Expect(err).To(BeAssignableToTypeOf(ErrZeroDivision{}))
		})
	})
	Describe("Getting Weight of Operation", func() {
		It("weights operation correctly", func() {
			Expect(newDivide.Weight()).To(Equal(2))
		})
	})
})
