package service

import (
	"database/sql"
	"laundry/errtype"
	"laundry/repo"
	"laundry/warehouse"
	"time"
)

/*
	Services & Repo
*/
var (
	/*
		Services
	*/
	userProfileService = NewUserProfileSvc()

	/*
		Repositories
	*/
	genderRepo = repo.NewGenderRepo()
	authRepo = repo.NewUserRepo()
	roleRepo = repo.NewRoleRepo()
	userProfileRepo = repo.NewUserProfileRepo()
)

/*
	Transactional Helpers
*/
func createTransactional() (*sql.Tx, time.Time, *errtype.Error) {
	now := time.Now()
	tx, err := warehouse.Conn.DB.Begin()
	if err != nil {
		return nil, time.Time{}, errtype.InternalServerError(err)
	}

	return tx, now, nil
}

func rollback(tx *sql.Tx, errType *errtype.Error) *errtype.Error {
	if errType != nil {
		if err := tx.Rollback(); err != nil {
			return errtype.InternalServerError(err)
		}

		return errType
	}
	return nil
}

func commit(tx *sql.Tx, errType *errtype.Error) *errtype.Error {
	if errType != nil {
		return errType
	}
	if err := tx.Commit(); err != nil {
		return errtype.InternalServerError(err)
	}
	return nil
}