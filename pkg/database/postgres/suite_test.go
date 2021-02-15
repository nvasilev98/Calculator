package postgres_test

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"

	"github.com/nvasilev98/calculator/pkg/database/connection"
	. "github.com/nvasilev98/calculator/pkg/database/postgres"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const host = "localhost"
const port = "5431"
const dbname = "calculation"
const username = "user"
const password = "pass"

var calculation *CalculationDAO
var dbTest *sql.DB

func TestPostgres(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Postgres Suite")
}

var _ = BeforeSuite(func() {
	var err error
	dbTest, err = connection.ConnectDB(connection.ConfigDatabase{
		Host:     host,
		Port:     port,
		Name:     dbname,
		Username: username,
		Password: password,
	})
	Expect(err).ToNot(HaveOccurred())
	calculation, err = NewCalculation(dbTest)
	Expect(err).ToNot(HaveOccurred())

})
