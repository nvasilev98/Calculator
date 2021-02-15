package operation_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/nvasilev98/calculator/pkg/calculation/operation"
)

var _ = Describe("Operation Extract Unit Tests", func() {
	var (
		newExtract Extract
	)
	Describe("Creating Operation Extract", func() {
		It("extracts two numbers correctly", func() {
			res, err := newExtract.Evaluate(3, 2)
			Expect(err).To(BeNil())
			Expect(res).To(Equal(float64(1)))
		})
		It("extracts two bigger numbers correctly", func() {
			res, err := newExtract.Evaluate(10, 3)
			Expect(err).To(BeNil())
			Expect(res).To(Equal(float64(7)))
		})
	})
	Describe("Getting Weight of Operation", func() {
		It("weights operation correctly", func() {
			Expect(newExtract.Weight()).To(Equal(1))
		})
	})

})
