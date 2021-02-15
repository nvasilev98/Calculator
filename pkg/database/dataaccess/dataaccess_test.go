package dataaccess_test

import (
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/nvasilev98/calculator/pkg/database/dataaccess"

	. "github.com/nvasilev98/calculator/pkg/database/dataaccess/mocks"
	"github.com/nvasilev98/calculator/pkg/database/postgres"
)

var _ = Describe("Dataaccess Unit Tests", func() {
	var (
		mockCtrl           *gomock.Controller
		mockCalculationDAO *MockCalculationDAO
		calculationDAO     *DataAccess
	)
	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockCalculationDAO = NewMockCalculationDAO(mockCtrl)
		calculationDAO = NewDataAccess(dbTest, mockCalculationDAO)

	})
	AfterEach(func() {
		mockCtrl.Finish()
	})

	It("inserts expression", func() {
		mockCalculationDAO.EXPECT().InsertCalculation(dbTest, postgres.Model{Expression: "2+3"}).Return(nil)
		err := calculationDAO.InsertCalculation(postgres.Model{Expression: "2+3"})
		Expect(err).ToNot(HaveOccurred())
	})
	It("propagates error when insert expression fail", func() {
		mockCalculationDAO.EXPECT().InsertCalculation(dbTest, postgres.Model{Expression: "4*2"}).Return(errors.New("Failed to insert"))
		err := calculationDAO.InsertCalculation(postgres.Model{Expression: "4*2"})
		Expect(err).To(HaveOccurred())
	})
	It("selects by id", func() {
		mockCalculationDAO.EXPECT().SelectByID(dbTest, int64(1)).Return("3+4", nil)
		expr, err := calculationDAO.SelectByID(1)
		Expect(err).ToNot(HaveOccurred())
		Expect(expr).To(Equal("3+4"))
	})
	It("propagates error when select by id fail", func() {
		mockCalculationDAO.EXPECT().SelectByID(dbTest, int64(1)).Return("", errors.New("Invalid id"))
		_, err := calculationDAO.SelectByID(1)
		Expect(err).To(HaveOccurred())
	})

	It("selects last id", func() {
		mockCalculationDAO.EXPECT().SelectLastID(dbTest).Return(int64(2), nil)
		id, err := calculationDAO.SelectLastID()
		Expect(err).ToNot(HaveOccurred())
		Expect(id).To(Equal(int64(2)))
	})
	It("propagates error when select last id fail", func() {
		mockCalculationDAO.EXPECT().SelectLastID(dbTest).Return(int64(0), errors.New("Empty table"))
		_, err := calculationDAO.SelectLastID()
		Expect(err).To(HaveOccurred())
	})
	It("selects result by last id", func() {
		mockCalculationDAO.EXPECT().SelectResultByID(dbTest, int64(2)).Return(postgres.Response{Result: 0, Error: "Parsing failed due to invalid symbol"}, true, nil)
		response, found, err := calculationDAO.SelectResultByID(2)
		Expect(err).ToNot(HaveOccurred())
		Expect(found).To(BeTrue())
		Expect(response.Error).To(Equal("Parsing failed due to invalid symbol"))
		Expect(response.Result).To(Equal(float64(0)))

	})
	It("propagates error when select result by last id fail", func() {
		mockCalculationDAO.EXPECT().SelectResultByID(dbTest, int64(3)).Return(postgres.Response{}, false, errors.New("Empty table"))
		_, _, err := calculationDAO.SelectResultByID(3)
		Expect(err).To(HaveOccurred())

	})
	It("updates result and error", func() {
		mockCalculationDAO.EXPECT().Update(dbTest, postgres.Model{int64(1), "2+3", float64(5), "nil"}).Return(nil)
		err := calculationDAO.Update(postgres.Model{int64(1), "2+3", float64(5), "nil"})
		Expect(err).ToNot(HaveOccurred())

	})
	It("propagates error when update fail", func() {
		mockCalculationDAO.EXPECT().Update(dbTest, postgres.Model{int64(1), "21/7", float64(5), "nil"}).Return(errors.New("Failed to update row"))
		err := calculationDAO.Update(postgres.Model{int64(1), "21/7", float64(5), "nil"})
		Expect(err).To(HaveOccurred())

	})
})
