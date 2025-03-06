package postgres

import (
	"database/sql"

	"github.com/charmingruby/impr/service/poll/internal/poll/core/model"
	"github.com/charmingruby/impr/service/poll/internal/shared/custom_err"
	"github.com/jmoiron/sqlx"
)

const (
	createPoll   = "create poll"
	findPollByID = "find poll by id"
	deletePoll   = "delete poll"
)

func pollQueries() map[string]string {
	return map[string]string{
		createPoll: `INSERT INTO polls
		(id, name)
		VALUES ($1, $2)
		RETURNING *`,
		findPollByID: `SELECT * FROM polls 
		WHERE id = $1`,
		deletePoll: `UPDATE polls
		SET deleted_at = $1, updated_at = $2
		WHERE id = $3 AND deleted_at IS NULL`,
	}
}

func NewPollRepository(db *sqlx.DB) (*PollRepository, error) {
	stmts := make(map[string]*sqlx.Stmt)

	for queryName, statement := range pollQueries() {
		stmt, err := db.Preparex(statement)
		if err != nil {
			return nil,
				custom_err.NewPreparationErr(queryName, "poll", err)
		}

		stmts[queryName] = stmt
	}

	return &PollRepository{
		db:    db,
		stmts: stmts,
	}, nil
}

type PollRepository struct {
	db    *sqlx.DB
	stmts map[string]*sqlx.Stmt
}

func (r *PollRepository) statement(queryName string) (*sqlx.Stmt, error) {
	stmt, ok := r.stmts[queryName]

	if !ok {
		return nil,
			custom_err.NewStatementNotPreparedErr(queryName, "poll")
	}

	return stmt, nil
}

func (r *PollRepository) Store(model *model.Poll) error {
	stmt, err := r.statement(createPoll)
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(
		model.ID,
		model.Name,
	); err != nil {
		return custom_err.NewPersistenceErr(err, "poll store", "postgres")
	}

	return nil
}

func (r *PollRepository) FindByID(id string) (*model.Poll, error) {
	stmt, err := r.statement(findPollByID)
	if err != nil {
		return nil, err
	}

	var poll model.Poll
	if err := stmt.Get(&poll, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, custom_err.NewPersistenceErr(err, "poll find_by_id", "postgres")
	}

	return &poll, nil
}

func (r *PollRepository) Delete(model *model.Poll) error {
	stmt, err := r.statement(deletePoll)
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(model.UpdatedAt, model.ID); err != nil {
		return custom_err.NewPersistenceErr(err, "poll delete", "postgres")
	}

	return nil
}
