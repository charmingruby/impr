package postgres

import (
	"database/sql"
	"time"

	"github.com/charmingruby/impr/service/poll/internal/poll/core/model"
	"github.com/charmingruby/impr/service/poll/internal/shared/custom_err"
	"github.com/jmoiron/sqlx"
)

const (
	findPollByID                       = "find poll by id"
	findPollByIDAndOwnerID             = "find poll by id and owner id"
	findPollByTitleAndOwnerID          = "find poll by title and owner id"
	findAllOPollsWithInferiorExpiresAt = "find all polls with inferior expires at"
	createPoll                         = "create poll"
	updatePoll                         = "update poll"
)

func pollQueries() map[string]string {
	return map[string]string{
		findPollByID: `SELECT * FROM polls 
		WHERE id = $1`,
		findPollByIDAndOwnerID: `SELECT * FROM polls
		WHERE id = $1 AND owner_id = $2`,
		findPollByTitleAndOwnerID: `SELECT * FROM polls
		WHERE title = $1 AND owner_id = $2`,
		findAllOPollsWithInferiorExpiresAt: `SELECT * FROM polls
		WHERE expires_at < $1`,
		createPoll: `INSERT INTO polls
		(id, title, question, status, expires_at, owner_id)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING *`,
		updatePoll: `UPDATE polls
		SET status = $1, updated_at = $2
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
		model.ExpiresAt,
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

func (r *PollRepository) FindAllWithInferiorExpiresAt(expiresAt time.Time) ([]model.Poll, error) {
	stmt, err := r.statement(findAllOPollsWithInferiorExpiresAt)
	if err != nil {
		return nil, err
	}

	var models []model.Poll
	if err := stmt.Select(&models, expiresAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, custom_err.NewPersistenceErr(err, "poll find_all_with_inferior_expires_at", "postgres")
	}

	return models, nil
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
