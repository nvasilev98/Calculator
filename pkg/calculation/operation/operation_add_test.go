package operation_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/nvasilev98/calculator/pkg/calculation/operation"
)

var _ = Describe("Operation Add Unit Tests", func() {

	var (
		newAdd Add
	)
	Describe("Creating Operation Add", func() {
		It("adds numbers correctly", func() {
			res, err := newAdd.Evaluate(1, 2)
			Expect(err).To(BeNil())
			Expect(res).To(Equal(float64(3)))
		})
		It("adds bigger numbers correctly", func() {
			res, err := newAdd.Evaluate(13, 22)
			Expect(err).To(BeNil())
			Expect(res).To(Equal(float64(35)))
		})
	})
	Describe("Getting Weight of Operation", func() {
		It("weights operation correctly", func() {
			Expect(newAdd.Weight()).To(Equal(1))
		})
	})
})
