package calculation_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/nvasilev98/calculator/pkg/calculation"
)

var _ = Describe("Parser Unit Tests", func() {
	var (
		operationfactory = NewOperationFactory()
		parser           = NewParser(operationfactory)
	)
	Describe("Validate Input", func() {
		It("expects to fail on empty expression", func() {
			_, err := parser.Parse("")
			Expect(err).To(HaveOccurred())
			Expect(err).To(BeAssignableToTypeOf(ErrEmptyString{}))
		})
		It("expects to fail on empty spaces expression", func() {
			_, err := parser.Parse("   ")
			Expect(err).To(HaveOccurred())
			Expect(err).To(BeAssignableToTypeOf(ErrEmptyString{}))
		})
		It("validates digit expression correctly", func() {
			res, err := parser.Parse("5")
			Expect(err).To(BeNil())
			Expect(res).To(Equal([]string{"5"}))
		})
		It("validates number expression correctly", func() {
			res, err := parser.Parse("51")
			Expect(err).To(BeNil())
			Expect(res).To(Equal([]string{"51"}))
		})
		It("validates float expression correctly", func() {
			res, err := parser.Parse("3.6")
			Expect(err).To(BeNil())
			Expect(res).To(Equal([]string{"3.6"}))
		})
		It("expects to fail on wrong float expression", func() {
			_, err := parser.Parse("3..6")
			Expect(err).To(HaveOccurred())
			Expect(err).To(BeAssignableToTypeOf(ErrInvalidFloatingNumber{}))
		})
		It("validates bracket expression correctly", func() {
			res, err := parser.Parse("(5)")
			Expect(err).To(BeNil())
			Expect(res).To(Equal([]string{"(", "5", ")"}))
		})
		It("validates expression with operation correctly", func() {
			res, err := parser.Parse("5+3")
			Expect(err).To(BeNil())
			Expect(res).To(Equal([]string{"5", "+", "3"}))
		})
		It("expects to fail on extra operation expression", func() {
			_, err := parser.Parse("5++3")
			Expect(err).To(HaveOccurred())
			Expect(err).To(BeAssignableToTypeOf(ErrOperationOrder{}))
		})
		It("expects to fail on expression with last symbol equal to operation", func() {
			_, err := parser.Parse("5+3-")
			Expect(err).To(HaveOccurred())
			Expect(err).To(BeAssignableToTypeOf(ErrOperationOrder{}))
		})
		It("expects to fail on expression without operation", func() {
			_, err := parser.Parse("5 3")
			Expect(err).To(HaveOccurred())
			Expect(err).To(BeAssignableToTypeOf(ErrOperationOrder{}))
		})
		It("validates expression with spaces correctly", func() {
			res, err := parser.Parse("5 + 3")
			Expect(err).To(BeNil())
			Expect(res).To(Equal([]string{"5", "+", "3"}))
		})
		It("expects to fail on expression with invalid symbol", func() {
			_, err := parser.Parse("5+3a")
			Expect(err).To(HaveOccurred())
			Expect(err).To(BeAssignableToTypeOf(ErrInvalidSymbol{}))
		})
	})

})
