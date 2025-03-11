package postgres

import (
	"database/sql"

	"github.com/charmingruby/impr/service/poll/internal/poll/core/model"
	"github.com/charmingruby/impr/service/poll/internal/shared/custom_err"
	"github.com/jmoiron/sqlx"
)

const (
	findVoteByPollIDAndUserID = "find vote by poll id and user id"
	createVote                = "create vote"
)

func voteQueries() map[string]string {
	return map[string]string{
		findVoteByPollIDAndUserID: `SELECT * FROM votes
		WHERE poll_id = $1 AND user_id = $2`,
		createVote: `INSERT INTO votes
		(id, poll_option_id, poll_id, user_id)
		VALUES ($1, $2, $3, $4)
		RETURNING *`,
	}
}

func NewVoteRepository(db *sqlx.DB) (*VoteRepository, error) {
	stmts := make(map[string]*sqlx.Stmt)

	for queryName, statement := range voteQueries() {
		stmt, err := db.Preparex(statement)
		if err != nil {
			return nil,
				custom_err.NewPreparationErr(queryName, "vote", err)
		}

		stmts[queryName] = stmt
	}

	return &VoteRepository{
		db:    db,
		stmts: stmts,
	}, nil
}

type VoteRepository struct {
	db    *sqlx.DB
	stmts map[string]*sqlx.Stmt
}

func (r *VoteRepository) statement(queryName string) (*sqlx.Stmt, error) {
	stmt, ok := r.stmts[queryName]

	if !ok {
		return nil,
			custom_err.NewStatementNotPreparedErr(queryName, "vote")
	}

	return stmt, nil
}

func (r *VoteRepository) Store(model *model.Vote) error {
	stmt, err := r.statement(createVote)
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(
		model.ID,
		model.PollOptionID,
		model.PollID,
		model.UserID,
	); err != nil {
		return custom_err.NewPersistenceErr(err, "vote store", "postgres")
	}

	return nil
}

func (r *VoteRepository) FindByPollIDAndUserID(pollID, userID string) (*model.Vote, error) {
	stmt, err := r.statement(findVoteByPollIDAndUserID)
	if err != nil {
		return nil, err
	}

	var model model.Vote
	if err := stmt.Get(&model, pollID, userID); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, custom_err.NewPersistenceErr(err, "vote find_by_poll_id_and_user_id", "postgres")
	}

	return &model, nil
}
