package env_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/nvasilev98/calculator/cmd/webapp/env"
)

var _ = Describe("EnvReader Unit Tests", func() {
	const envName = "ENV"
	const data = "val"
	var env string
	BeforeEach(func() {
		Expect(os.Setenv(envName, data)).To(Succeed())
	})
	AfterEach(func() {
		Expect(os.Unsetenv(envName)).To(Succeed())
	})
	Describe("Read String as Env Variable", func() {
		It("reads environment variable when it is set", func() {
			err := ReadEnvVariable(envName, &env)
			Expect(err).ToNot(HaveOccurred())
			Expect(env).To(Equal(data))
		})
		It("returns an error when environment variable is not set", func() {
			Expect(os.Unsetenv(envName)).To(Succeed())
			err := ReadEnvVariable(envName, &env)
			Expect(err).To(HaveOccurred())
		})
	})
})
