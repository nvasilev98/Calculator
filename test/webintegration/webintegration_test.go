package webintegration_test

import (
	"context"
	"net/http"

	"github.com/nvasilev98/calculator/pkg/api"
	. "github.com/nvasilev98/calculator/test/webintegration/api"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Web Integration Tests", func() {
	var calcClient *CalculatorClient
	When("credentials are incorrect", func() {
		It("returns unauthorized when credentials are not provided", func() {
			calcClient = NewCalculatorClient(http.DefaultClient, calcExecutable.URL(), Credentials{})
			_, err := calcClient.GetResult(context.Background(), "2*/3+3-2")
			Expect(err).To(HaveOccurred())
			Expect(err).To(BeAssignableToTypeOf(api.UnauthorizedError{}))
		})
		It("returns unauthorized when credentials are invalid", func() {
			calcClient = NewCalculatorClient(http.DefaultClient, calcExecutable.URL(), Credentials{
				Username: "invalid-name",
				Password: "invalid-pass",
			})
			_, err := calcClient.GetResult(context.Background(), "2*/3+3-2")
			Expect(err).To(HaveOccurred())
			Expect(err).To(BeAssignableToTypeOf(api.UnauthorizedError{}))
		})
	})

	When("credentials are correct", func() {
		BeforeEach(func() {
			calcClient = NewCalculatorClient(http.DefaultClient, calcExecutable.URL(), Credentials{
				Username: username,
				Password: password,
			})
		})
		It("successfully calculates with precede operator", func() {
			resp, err := calcClient.GetResult(context.Background(), "2+4*3/2-1-14")
			Expect(err).ToNot(HaveOccurred())
			Expect(resp).To(Equal(float64(-7)))
		})
		It("successfully calculates with brackets", func() {
			resp, err := calcClient.GetResult(context.Background(), "(32/2 +20)/3-(2+3)*1.2")
			Expect(err).ToNot(HaveOccurred())
			Expect(resp).To(Equal(float64(6)))
		})
		It("fails to calculate with illegal symbol", func() {
			_, err := calcClient.GetResult(context.Background(), "(a+12)-3/2")
			Expect(err).To(HaveOccurred())
			Expect(err).To(BeAssignableToTypeOf(api.APIError{}))
		})
		It("fails to calculate with mismatched brackets", func() {
			_, err := calcClient.GetResult(context.Background(), ")1+1)/2")
			Expect(err).To(HaveOccurred())
			Expect(err).To(BeAssignableToTypeOf(api.APIError{}))
		})
		It("fails to calculate with more operators", func() {
			_, err := calcClient.GetResult(context.Background(), "2*/3+3-2")
			Expect(err).To(HaveOccurred())
			Expect(err).To(BeAssignableToTypeOf(api.APIError{}))
		})
	})
})
