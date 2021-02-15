package dataaccess

import (
	"database/sql"

	"github.com/nvasilev98/calculator/pkg/database/postgres"
)

//go:generate mockgen --source=dataaccess.go --destination mocks/mock_dataaccess.go --package mocks

type CalculationDAO interface {
	InsertCalculation(db *sql.DB, model postgres.Model) error
	SelectByID(db *sql.DB, id int64) (string, error)
	SelectLastID(db *sql.DB) (int64, error)
	SelectResultByID(db *sql.DB, id int64) (postgres.Response, bool, error)
	Update(db *sql.DB, model postgres.Model) error
}

func NewDataAccess(db *sql.DB, calculationDAO CalculationDAO) *DataAccess {

	return &DataAccess{
		db:             db,
		calculationDAO: calculationDAO,
	}
}

type DataAccess struct {
	db             *sql.DB
	calculationDAO CalculationDAO
}

func (d *DataAccess) InsertCalculation(model postgres.Model) error {
	return d.calculationDAO.InsertCalculation(d.db, model)
}
func (d *DataAccess) SelectByID(id int64) (string, error) {
	return d.calculationDAO.SelectByID(d.db, id)
}
func (d *DataAccess) SelectLastID() (int64, error) {
	return d.calculationDAO.SelectLastID(d.db)
}
func (d *DataAccess) SelectResultByID(id int64) (postgres.Response, bool, error) {
	return d.calculationDAO.SelectResultByID(d.db, id)
}
func (d *DataAccess) Update(model postgres.Model) error {
	return d.calculationDAO.Update(d.db, model)
}
