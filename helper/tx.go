package helper

import "database/sql"

func CommitOrRollback(err error, tx *sql.Tx) {
	if err != nil {
		rollbackError := tx.Rollback()
		PanicIfError(rollbackError)
	} else {
		commitError := tx.Commit()
		PanicIfError(commitError)
	}
}
