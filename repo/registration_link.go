package repo

import (
	"database/sql"
	"laundry/entity"
	"laundry/errtype"
)

type registrationLinkRepo struct {
	TableName string
}

func NewRegistrationLinkRepo() *registrationLinkRepo {
	return &registrationLinkRepo{
		TableName: "registration_link"}
}

func (rl *registrationLinkRepo) GetCountByEmail(db *sql.Tx, email string) (count int, errType *errtype.Error) {
	query := "SELECT " +
			"	COUNT(id) " +
			"FROM " + rl.TableName + " " +
			"WHERE " +
			"	email = $1 AND " +
			"	is_deleted = FALSE "

	scanFunc := func(row *sql.Row) error {
		return row.Scan(&count)
	}

	err := fetchRowTx(db, query, []interface{}{ email }, scanFunc)
	if err != nil {
		errType = errtype.InternalServerError(err)
	}
	return
}

func (rl *registrationLinkRepo) CreateRegistrationLink(repo entity.RegistrationLinkRepo) *errtype.Error {
	query := "INSERT INTO " + rl.TableName + " (" +
			"	email, expired_at, type, laundry_id," +
			"	created_at, created_by, updated_at, updated_by) " +
			"VALUES(" +
			"	$1, $2, $3, $4," +
			"	$5, $6, $7, $8)"

	params := []interface{}{
		repo.Email.String, repo.ExpiredAt.Time, repo.Type.String, repo.LaundryId.Int64,
		repo.CreatedAt.Time, repo.CreatedBy.Int64, repo.UpdatedAt.Time, repo.UpdatedBy.Int64}

	if err := preparedQuery(query, params); err != nil {
		return errtype.InternalServerError(err)
	}
	return nil
}