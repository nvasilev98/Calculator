package postgres_test

import (
	"github.com/nvasilev98/calculator/pkg/database/postgres"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CalculationDAO Unit Tests", func() {
	BeforeEach(func() {
		_, err := dbTest.Exec("truncate calculation restart identity")
		Expect(err).ToNot(HaveOccurred())
	})
	AfterEach(func() {
		_, err := dbTest.Exec("truncate calculation restart identity")
		Expect(err).ToNot(HaveOccurred())
	})
	Describe("Select last Id", func() {
		It("returns an error when table is empty", func() {
			_, err := calculation.SelectLastID(dbTest)
			Expect(err).To(HaveOccurred())
		})

		It("returns correct last id in single record table", func() {
			expression := postgres.Model{Expression: "1*3-2"}
			err := calculation.InsertCalculation(dbTest, expression)
			Expect(err).ToNot(HaveOccurred())
			res, err1 := calculation.SelectLastID(dbTest)
			Expect(err1).ToNot(HaveOccurred())
			Expect(res).To(Equal(int64(1)))
		})
		It("returns correct last id when more than one record is provided", func() {
			firstExpression := postgres.Model{Expression: "4213/3"}
			err := calculation.InsertCalculation(dbTest, firstExpression)
			Expect(err).ToNot(HaveOccurred())
			secondExpression := postgres.Model{Expression: "(1+2)*5"}
			err = calculation.InsertCalculation(dbTest, secondExpression)
			Expect(err).ToNot(HaveOccurred())
			res, err1 := calculation.SelectLastID(dbTest)
			Expect(err1).ToNot(HaveOccurred())
			Expect(res).To(Equal(int64(2)))
		})

	})

	Describe("Select By Id", func() {
		It("returns an error when table is empty", func() {
			_, err := calculation.SelectByID(dbTest, 1)
			Expect(err).To(HaveOccurred())
		})
		It("successfully selected row with valid id", func() {
			expression := postgres.Model{Expression: "5*4/2"}
			err := calculation.InsertCalculation(dbTest, expression)
			Expect(err).ToNot(HaveOccurred())
			res, err := calculation.SelectByID(dbTest, 1)
			Expect(err).ToNot(HaveOccurred())
			Expect(res).To(Equal("5*4/2"))

		})
		It("returns error when selected row with invalid id", func() {
			expression := postgres.Model{Expression: "3-4+(3-2)"}
			err := calculation.InsertCalculation(dbTest, expression)
			Expect(err).ToNot(HaveOccurred())
			_, err = calculation.SelectByID(dbTest, 2)
			Expect(err).To(HaveOccurred())

		})
	})

	Describe("Select Result By Id", func() {
		It("fails when table is empty", func() {
			_, found, err := calculation.SelectResultByID(dbTest, 1)
			Expect(err).ToNot(HaveOccurred())
			Expect(found).To(BeFalse())
		})
		It("successfully select result with valid id", func() {
			expression := postgres.Model{Expression: "5*4/2"}
			err := calculation.InsertCalculation(dbTest, expression)
			Expect(err).ToNot(HaveOccurred())
			err = calculation.Update(dbTest, postgres.Model{1, expression.Expression, 10, "nil"})
			response, found, err := calculation.SelectResultByID(dbTest, 1)
			Expect(err).ToNot(HaveOccurred())
			Expect(found).To(BeTrue())
			Expect(response.Error).To(Equal("nil"))
			Expect(response.Result).To(Equal(float64(10)))
		})
		It("fails to select result with invalid id", func() {
			expression := postgres.Model{Expression: "5*4"}
			err := calculation.InsertCalculation(dbTest, expression)
			Expect(err).ToNot(HaveOccurred())
			err = calculation.Update(dbTest, postgres.Model{1, expression.Expression, 20, "nil"})
			_, found, err := calculation.SelectResultByID(dbTest, 2)
			Expect(err).ToNot(HaveOccurred())
			Expect(found).To(BeFalse())

		})

	})

	Describe("Insert into Table", func() {
		It("successfully inserts a row in table", func() {
			expression := postgres.Model{Expression: "7-4/2"}
			err := calculation.InsertCalculation(dbTest, expression)
			Expect(err).ToNot(HaveOccurred())
			res, err := calculation.SelectByID(dbTest, 1)
			Expect(err).ToNot(HaveOccurred())
			Expect(res).To(Equal("7-4/2"))
		})
	})

	Describe("Update by Id", func() {
		It("successfully updates a row in table", func() {
			expression := postgres.Model{Expression: "3+5-2"}
			err := calculation.InsertCalculation(dbTest, expression)
			Expect(err).ToNot(HaveOccurred())
			res, err := calculation.SelectByID(dbTest, 1)
			Expect(err).ToNot(HaveOccurred())
			Expect(res).To(Equal("3+5-2"))
			updatedModel := postgres.Model{1, expression.Expression, 6, "nil"}
			err = calculation.Update(dbTest, updatedModel)
			Expect(err).ToNot(HaveOccurred())
			response, found, err := calculation.SelectResultByID(dbTest, 1)
			Expect(err).ToNot(HaveOccurred())
			Expect(found).To(BeTrue())
			Expect(response.Error).To(Equal("nil"))
			Expect(response.Result).To(Equal(float64(6)))

		})
		It("fails to update invalid row", func() {
			updatedModel := postgres.Model{1, "3+3", 6, "nil"}
			err := calculation.Update(dbTest, updatedModel)
			Expect(err).To(HaveOccurred())
		})
	})
})
