package postgres

import (
	"database/sql"

	"github.com/charmingruby/impr/service/poll/internal/poll/core/model"
	"github.com/charmingruby/impr/service/poll/internal/shared/custom_err"
	"github.com/jmoiron/sqlx"
)

const (
	findPollByID              = "find poll by id"
	findPollByIDAndOwnerID    = "find poll by id and owner id"
	findPollByTitleAndOwnerID = "find poll by title and owner id"
	createPoll                = "create poll"
	updatePoll                = "update poll"
)

func pollQueries() map[string]string {
	return map[string]string{
		findPollByID: `SELECT * FROM polls 
		WHERE id = $1`,
		findPollByIDAndOwnerID: `SELECT * FROM polls
		WHERE id = $1 AND owner_id = $2`,
		findPollByTitleAndOwnerID: `SELECT * FROM polls
		WHERE title = $1 AND owner_id = $2`,
		createPoll: `INSERT INTO polls
		(id, title, question, status, expiration_time, owner_id)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING *`,
		updatePoll: `UPDATE polls
		SET status = $1 AND updated_at = $2
		WHERE id = $3`,
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
		model.Title,
		model.Question,
		model.Status,
		model.ExpirationTime,
		model.OwnerID,
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

	var model model.Poll
	if err := stmt.Get(&model, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, custom_err.NewPersistenceErr(err, "poll find_by_id", "postgres")
	}

	return &model, nil
}

func (r *PollRepository) FindByIDAndOwnerID(id, ownerID string) (*model.Poll, error) {
	stmt, err := r.statement(findPollByIDAndOwnerID)
	if err != nil {
		return nil, err
	}

	var model model.Poll
	if err := stmt.Get(&model, id, ownerID); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, custom_err.NewPersistenceErr(err, "poll find_by_id_and_owner_id", "postgres")
	}

	return &model, nil
}

func (r *PollRepository) FindByTitleAndOwnerID(title, ownerID string) (*model.Poll, error) {
	stmt, err := r.statement(findPollByTitleAndOwnerID)
	if err != nil {
		return nil, err
	}

	var model model.Poll
	if err := stmt.Get(&model, title, ownerID); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, custom_err.NewPersistenceErr(err, "poll find_by_title_and_owner_id", "postgres")
	}

	return &model, nil
}

func (r *PollRepository) Save(model *model.Poll) error {
	stmt, err := r.statement(updatePoll)
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(model.Status, model.UpdatedAt, model.ID); err != nil {
		return custom_err.NewPersistenceErr(err, "poll update", "postgres")
	}

	return nil
}
