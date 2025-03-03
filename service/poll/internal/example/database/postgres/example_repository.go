package postgres

import (
	"database/sql"

	"github.com/charmingruby/bob/internal/shared/custom_err/database_err"
	"github.com/charmingruby/bob/internal/shared/custom_err/database_err/sql_err"
	"github.com/charmingruby/bob/internal/example/core/model"
	"github.com/jmoiron/sqlx"
)

const (
	createExample   = "create example"
	findExampleByID = "find example by id"
	deleteExample   = "delete example"
)

func exampleQueries() map[string]string {
	return map[string]string{
		createExample: `INSERT INTO examples
		(id, name)
		VALUES ($1, $2)
		RETURNING *`,
		findExampleByID: `SELECT * FROM examples 
		WHERE id = $1`,
		deleteExample: `UPDATE examples
		SET deleted_at = $1, updated_at = $2
		WHERE id = $3 AND deleted_at IS NULL`,
	}
}

func NewExampleRepository(db *sqlx.DB) (*ExampleRepository, error) {
	stmts := make(map[string]*sqlx.Stmt)

	for queryName, statement := range exampleQueries() {
		stmt, err := db.Preparex(statement)
		if err != nil {
			return nil,
				sql_err.NewPreparationErr(queryName, "example", err)
		}

		stmts[queryName] = stmt
	}

	return &ExampleRepository{
		db:    db,
		stmts: stmts,
	}, nil
}

type ExampleRepository struct {
	db    *sqlx.DB
	stmts map[string]*sqlx.Stmt
}

func (r *ExampleRepository) statement(queryName string) (*sqlx.Stmt, error) {
	stmt, ok := r.stmts[queryName]

	if !ok {
		return nil,
			sql_err.NewStatementNotPreparedErr(queryName, "example")
	}

	return stmt, nil
}

func (r *ExampleRepository) Store(model *model.Example) error {
	stmt, err := r.statement(createExample)
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(
		model.ID,
		model.Name,
	); err != nil {
		return database_err.NewPersistenceErr(err, "example store", "postgres")
	}

	return nil
}

func (r *ExampleRepository) FindByID(id string) (*model.Example, error) {
	stmt, err := r.statement(findExampleByID)
	if err != nil {
		return nil, err
	}

	var example model.Example
	if err := stmt.Get(&example, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, database_err.NewPersistenceErr(err, "example find_by_id", "postgres")
	}

	return &example, nil
}

func (r *ExampleRepository) Delete(model *model.Example) error {
	stmt, err := r.statement(deleteExample)
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(model.DeletedAt, model.UpdatedAt, model.ID); err != nil {
		return database_err.NewPersistenceErr(err, "example delete", "postgres")
	}

	return nil
}