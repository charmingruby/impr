package custom_err

import (
	"fmt"
)

func NewPreparationErr(queryName string, repository string, err error) *PersistenceErr {
	preparationErr := fmt.Errorf("unable to prepare the query:`%s` on %s repository, original err: %s", queryName, repository, err.Error())
	return NewPersistenceErr(preparationErr, "prepare", "postgres")
}

func NewStatementNotPreparedErr(queryName string, repository string) *PersistenceErr {
	preparationErr := fmt.Errorf("query `%s` is not prepared on %s repository", queryName, repository)
	return NewPersistenceErr(preparationErr, "statement not prepared", "postgres")
}
