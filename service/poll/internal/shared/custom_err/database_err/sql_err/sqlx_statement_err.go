package sql_err

import (
	"fmt"

	"github.com/charmingruby/bob/internal/shared/custom_err/database_err"
)

func NewPreparationErr(queryName string, repository string, err error) *database_err.PersistenceErr {
	preparationErr := fmt.Errorf("unable to prepare the query:`%s` on %s repository, original err: %s", queryName, repository, err.Error())
	return database_err.NewPersistenceErr(preparationErr, "prepare", "postgres")
}

func NewStatementNotPreparedErr(queryName string, repository string) *database_err.PersistenceErr {
	preparationErr := fmt.Errorf("query `%s` is not prepared on %s repository", queryName, repository)
	return database_err.NewPersistenceErr(preparationErr, "statement not prepared", "postgres")
}
