package env_test

import (
	"os"

	. "github.com/nvasilev98/calculator/cmd/webapp/env"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Env Unit Tests", func() {
	Describe("Load Configuration", func() {
		hostENV := "HOST"
		portENV := "PORT"
		userENV := "USERNAME"
		passENV := "PASSWORD"
		host := "0.0.0.0"
		port := "1234"
		user := "user"
		pass := "password"
		When("all environment variables are presented", func() {

			BeforeEach(func() {
				Expect(os.Setenv(hostENV, host)).To(Succeed())
				Expect(os.Setenv(portENV, port)).To(Succeed())
				Expect(os.Setenv(userENV, user)).To(Succeed())
				Expect(os.Setenv(passENV, pass)).To(Succeed())
			})
			AfterEach(func() {
				Expect(os.Unsetenv(hostENV)).To(Succeed())
				Expect(os.Unsetenv(portENV)).To(Succeed())
				Expect(os.Unsetenv(userENV)).To(Succeed())
				Expect(os.Unsetenv(passENV)).To(Succeed())
			})
			It("loads all variables successfully", func() {
				cfg, err := Load()
				Expect(err).ToNot(HaveOccurred())
				Expect(cfg.Host).To(Equal(host))
				Expect(cfg.Port).To(Equal(port))
				Expect(cfg.Username).To(Equal(user))
				Expect(cfg.Password).To(Equal(pass))
			})
		})
		When("host is missing", func() {

			BeforeEach(func() {
				Expect(os.Setenv(portENV, port)).To(Succeed())
				Expect(os.Setenv(userENV, user)).To(Succeed())
				Expect(os.Setenv(passENV, pass)).To(Succeed())
			})
			AfterEach(func() {
				Expect(os.Unsetenv(portENV)).To(Succeed())
				Expect(os.Unsetenv(userENV)).To(Succeed())
				Expect(os.Unsetenv(passENV)).To(Succeed())
			})
			It("fails to load configuration due to missing host", func() {
				_, err := Load()
				Expect(err).To(HaveOccurred())
			})
		})
		When("port is missing", func() {

			BeforeEach(func() {
				Expect(os.Setenv(hostENV, host)).To(Succeed())
				Expect(os.Setenv(userENV, user)).To(Succeed())
				Expect(os.Setenv(passENV, pass)).To(Succeed())
			})
			AfterEach(func() {
				Expect(os.Unsetenv(hostENV)).To(Succeed())
				Expect(os.Unsetenv(userENV)).To(Succeed())
				Expect(os.Unsetenv(passENV)).To(Succeed())
			})
			It("fails to load configuration due to missing port", func() {
				_, err := Load()
				Expect(err).To(HaveOccurred())
			})
		})
		When("username is missing", func() {

			BeforeEach(func() {
				Expect(os.Setenv(hostENV, host)).To(Succeed())
				Expect(os.Setenv(portENV, port)).To(Succeed())
				Expect(os.Setenv(passENV, pass)).To(Succeed())
			})
			AfterEach(func() {
				Expect(os.Unsetenv(hostENV)).To(Succeed())
				Expect(os.Unsetenv(portENV)).To(Succeed())
				Expect(os.Unsetenv(passENV)).To(Succeed())
			})
			It("fails to load configuration due to missing username", func() {
				_, err := Load()
				Expect(err).To(HaveOccurred())
			})
		})
		When("password is missing", func() {

			BeforeEach(func() {
				Expect(os.Setenv(hostENV, host)).To(Succeed())
				Expect(os.Setenv(portENV, port)).To(Succeed())
				Expect(os.Setenv(userENV, user)).To(Succeed())
			})
			AfterEach(func() {
				Expect(os.Unsetenv(hostENV)).To(Succeed())
				Expect(os.Unsetenv(portENV)).To(Succeed())
				Expect(os.Unsetenv(userENV)).To(Succeed())
			})
			It("fails to load configuration due to missing password", func() {
				_, err := Load()
				Expect(err).To(HaveOccurred())
			})
		})

	})
})
