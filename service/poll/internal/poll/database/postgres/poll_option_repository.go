package postgres

import (
	"database/sql"

	"github.com/charmingruby/impr/service/poll/internal/poll/core/model"
	"github.com/charmingruby/impr/service/poll/internal/shared/custom_err"
	"github.com/jmoiron/sqlx"
)

const (
	findPollOptionByID               = "find poll option by id"
	findAllPoolOptionByPollID        = "find poll option by poll id"
	findPollOptionByContentAndPollID = "find poll option by content and poll id"
	createPollOption                 = "create poll option"
)

func pollOptionQueries() map[string]string {
	return map[string]string{
		findPollOptionByID: `SELECT * FROM poll_options 
		WHERE id = $1`,
		findAllPoolOptionByPollID: `SELECT * FROM poll_options
		WHERE poll_id = $1`,
		findPollOptionByContentAndPollID: `SELECT * FROM poll_options
		WHERE content = $1 AND poll_id = $2`,
		createPollOption: `INSERT INTO poll_options
		(id, content, poll_id)
		VALUES ($1, $2, $3)
		RETURNING *`,
	}
}

func NewPollOptionRepository(db *sqlx.DB) (*PollOptionRepository, error) {
	stmts := make(map[string]*sqlx.Stmt)

	for queryName, statement := range pollOptionQueries() {
		stmt, err := db.Preparex(statement)
		if err != nil {
			return nil,
				custom_err.NewPreparationErr(queryName, "poll option", err)
		}

		stmts[queryName] = stmt
	}

	return &PollOptionRepository{
		db:    db,
		stmts: stmts,
	}, nil
}

type PollOptionRepository struct {
	db    *sqlx.DB
	stmts map[string]*sqlx.Stmt
}

func (r *PollOptionRepository) statement(queryName string) (*sqlx.Stmt, error) {
	stmt, ok := r.stmts[queryName]

	if !ok {
		return nil,
			custom_err.NewStatementNotPreparedErr(queryName, "poll option")
	}

	return stmt, nil
}

func (r *PollOptionRepository) FindByID(id string) (*model.PollOption, error) {
	stmt, err := r.statement(findPollOptionByID)
	if err != nil {
		return nil, err
	}

	var model model.PollOption
	if err := stmt.Get(&model, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, custom_err.NewPersistenceErr(err, "poll option find_by_id", "postgres")
	}

	return &model, nil
}

func (r *PollOptionRepository) FindAllByPollID(pollID string) ([]model.PollOption, error) {
	stmt, err := r.statement(findAllPoolOptionByPollID)
	if err != nil {
		return nil, err
	}

	var models []model.PollOption
	if err := stmt.Select(&models, pollID); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, custom_err.NewPersistenceErr(err, "poll option find_all_by_poll_id", "postgres")
	}

	return models, nil
}

func (r *PollOptionRepository) FindByContentAndPollID(content, pollID string) (*model.PollOption, error) {
	stmt, err := r.statement(findPollOptionByContentAndPollID)
	if err != nil {
		return nil, err
	}

	var model model.PollOption
	if err := stmt.Get(&model, content, pollID); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, custom_err.NewPersistenceErr(err, "poll option find_by_content_and_poll_id", "postgres")
	}

	return &model, nil
}

func (r *PollOptionRepository) Store(model *model.PollOption) error {
	stmt, err := r.statement(createPollOption)
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(
		model.ID,
		model.Content,
		model.PollID,
	); err != nil {
		return custom_err.NewPersistenceErr(err, "poll option store", "postgres")
	}

	return nil
}
