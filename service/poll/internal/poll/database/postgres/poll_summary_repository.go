package postgres

import (
	"time"

	"github.com/charmingruby/impr/service/poll/internal/poll/core/model"
	"github.com/charmingruby/impr/service/poll/internal/shared/custom_err"
	"github.com/jmoiron/sqlx"
)

const findSummaryByPollID = "find summary by poll id"

func summaryQueries() map[string]string {
	return map[string]string{
		findSummaryByPollID: `
SELECT 
    p.id AS poll_id,
    p.expires_at,
    po.id AS poll_option_id,
    po.content,
    COUNT(v.id) AS votes
FROM 
    polls p
LEFT JOIN 
    poll_options po ON p.id = po.poll_id
LEFT JOIN 
    votes v ON po.id = v.poll_option_id
WHERE 
    p.id = $1
GROUP BY 
    p.id, po.id
`,
	}
}

type SummaryRepository struct {
	db    *sqlx.DB
	stmts map[string]*sqlx.Stmt
}

func NewSummaryRepository(db *sqlx.DB) (*SummaryRepository, error) {
	stmts := make(map[string]*sqlx.Stmt)

	for queryName, statement := range summaryQueries() {
		stmt, err := db.Preparex(statement)
		if err != nil {
			return nil,
				custom_err.NewPreparationErr(queryName, "summary", err)
		}

		stmts[queryName] = stmt
	}

	return &SummaryRepository{
		db:    db,
		stmts: stmts,
	}, nil
}

func (r *SummaryRepository) statement(queryName string) (*sqlx.Stmt, error) {
	stmt, ok := r.stmts[queryName]

	if !ok {
		return nil, custom_err.NewStatementNotPreparedErr(queryName, "summary")
	}

	return stmt, nil
}

func (r *SummaryRepository) FindByPollID(pollID string) (*model.PollSummary, error) {
	stmt, err := r.statement(findSummaryByPollID)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Queryx(pollID)
	if err != nil {
		return nil, custom_err.NewPersistenceErr(err, "summary find_by_poll_id", "postgres")
	}
	defer rows.Close()

	var pollSummary model.PollSummary
	var options []model.PollSummaryOption

	for rows.Next() {
		var option model.PollSummaryOption
		var pollID string
		var expiresAt *time.Time

		if err := rows.Scan(&pollID, &expiresAt, &option.PollOptionID, &option.Content, &option.Votes); err != nil {
			return nil, custom_err.NewPersistenceErr(err, "summary find_by_poll_id", "postgres")
		}

		pollSummary.PollID = pollID
		pollSummary.ExpiresAt = expiresAt
		options = append(options, option)
	}

	pollSummary.Options = options

	return &pollSummary, nil
}
