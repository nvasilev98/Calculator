package postgres

import (
	"database/sql"

	"github.com/pkg/errors"
)

type Response struct {
	Result float64
	Error  string
}

type Model struct {
	ID         int64
	Expression string
	Result     float64
	CalcError  string
}

func NewCalculation(db *sql.DB) (*CalculationDAO, error) {

	if _, err := db.Exec(createCalculationTableSQL); err != nil {
		return nil, errors.Wrap(err, "failed to create table")
	}
	insertCalculationStmt, err := db.Prepare(insertCalculationSQL)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create insert statement")
	}
	selectByIDStmt, err := db.Prepare(selectByIDSQL)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create select by id statement")
	}
	selectLastIDStmt, err := db.Prepare(selectLastIDSQL)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create select last id statement")
	}
	selectResultByIDStmt, err := db.Prepare(selectResultByIDSQL)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create select result by id statement")
	}
	updateStmt, err := db.Prepare(updateCalculationResultSQL)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create update statement")
	}
	return &CalculationDAO{
		insertCalculationStatement: insertCalculationStmt,
		selectByIDStatement:        selectByIDStmt,
		selectLastIDStatement:      selectLastIDStmt,
		updateStatement:            updateStmt,
		selectResultByIDStatement:  selectResultByIDStmt,
	}, nil
}

type CalculationDAO struct {
	insertCalculationStatement *sql.Stmt
	selectByIDStatement        *sql.Stmt
	selectLastIDStatement      *sql.Stmt
	updateStatement            *sql.Stmt
	selectResultByIDStatement  *sql.Stmt
}

func (c *CalculationDAO) InsertCalculation(db *sql.DB, model Model) error {

	tx, err := db.Begin()

	if err != nil {
		return errors.Wrap(err, "failed to start transaction")
	}

	defer tx.Rollback()
	_, err = tx.Stmt(c.insertCalculationStatement).Exec(model.Expression)
	if err != nil {
		return errors.Wrap(err, "failed to insert expression")
	}
	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, "failed to commit transaction")
	}
	return nil

}

func (c *CalculationDAO) SelectByID(db *sql.DB, id int64) (string, error) {
	tx, err := db.Begin()
	if err != nil {
		return "", errors.Wrap(err, "failed to start transaction")
	}
	defer tx.Rollback()
	sqlRow := tx.Stmt(c.selectByIDStatement).QueryRow(id)
	var expr sql.NullString
	err = sqlRow.Scan(&expr)
	if err != nil {
		if err == sql.ErrNoRows {
			// there were no rows with this id
			return "", errors.Wrap(err, "wanted id has not been found")
		}
		return "", errors.Wrap(err, "failed to return a row")
	}
	return expr.String, nil

}
func (c *CalculationDAO) SelectLastID(db *sql.DB) (int64, error) {
	tx, err := db.Begin()
	if err != nil {
		return 0, errors.Wrap(err, "failed to start transaction")
	}
	defer tx.Rollback()
	var result sql.NullInt64
	sqlRow := tx.Stmt(c.selectLastIDStatement).QueryRow()

	err = sqlRow.Scan(&result)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.Wrap(err, "empty table")
		}
		return 0, errors.Wrap(err, "failed to return a row")
	}
	return result.Int64, nil

}
func (c *CalculationDAO) SelectResultByID(db *sql.DB, id int64) (Response, bool, error) {
	tx, err := db.Begin()
	if err != nil {
		return Response{}, false, errors.Wrap(err, "failed to start transaction")
	}
	defer tx.Rollback()
	sqlRow := tx.Stmt(c.selectResultByIDStatement).QueryRow(id)
	var result sql.NullFloat64
	var calcError sql.NullString
	err = sqlRow.Scan(&result, &calcError)
	if err != nil {
		if err == sql.ErrNoRows {
			return Response{}, false, nil
		}
		return Response{}, false, errors.Wrap(err, "failed to return a row")
	}
	response := Response{result.Float64, calcError.String}
	return response, true, nil
}

func (c *CalculationDAO) Update(db *sql.DB, model Model) error {
	tx, err := db.Begin()
	if err != nil {
		return errors.Wrap(err, "failed to start transaction")
	}
	defer tx.Rollback()
	if _, err := c.SelectByID(db, model.ID); err != nil {
		return errors.Wrap(err, "failed to update due to missing id")
	}
	_, err = tx.Stmt(c.updateStatement).Exec(model.ID, model.Result, model.CalcError)
	if err != nil {
		return errors.Wrap(err, "failed to update result")
	}
	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, "failed to commit transaction")
	}
	return nil
}
