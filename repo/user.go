package repo

import (
	"database/sql"
	"laundry/entity"
	"laundry/errtype"
	"laundry/warehouse"
)

type userRepo struct {
	TableName string
}

func NewUserRepo() *userRepo {
	return &userRepo{TableName: "users"}
}

func (u *userRepo) Register(db *sql.Tx, user entity.UserRepo) (errType *errtype.Error) {
	query := `INSERT INTO ` + u.TableName +
			`	(username, password, user_profile_id, 
				 created_at, created_by, updated_at, 
				 updated_by)
			  VALUES
				($1, $2, $3,
				 $4, $5, $6,
				 $7)`

	stmt, err := db.Prepare(query)
	if err != nil {
		return errtype.InternalServerError(err)
	}

	params := []interface{} {
		user.Username.String, user.Password.String, user.UserProfileId.Int64,
		user.CreatedAt.Time, user.CreatedBy.Int64, user.UpdatedAt.Time,
		user.UpdatedBy.Int64}

	_, err = stmt.Exec(params...)
	if err != nil {
		return errtype.InternalServerError(err)
	}

	return
}

func (u *userRepo) GetUserCountByUsername(username string) (count int, errType *errtype.Error) {
	db := warehouse.Conn.DB
	query := `SELECT 
				COUNT(id) 
			  FROM ` + u.TableName + ` 
			  WHERE 
				is_deleted = FALSE AND 
				username = $1 
			  LIMIT 1`

	row := db.QueryRow(query, username)
	err := row.Scan(&count)
	if err != nil {
		errType = errtype.InternalServerError(err)
	}
	return
}

func (u *userRepo) FindUserByUsername(username string) (result entity.UserRepo, errType *errtype.Error) {
	db := warehouse.Conn.DB

	userProfileTable := NewUserProfileRepo().TableName
	roleTable := NewRoleRepo().TableName
	query := `SELECT 
				u.id, u.username, u.password, u.user_profile_id, 
				r.role 
			  FROM ` + u.TableName + ` AS u
			  LEFT JOIN ` + userProfileTable + ` AS up
				ON u.user_profile_id = up.id
			  LEFT JOIN ` + roleTable + ` AS r 
				ON up.role_id = r.id
			  WHERE 
				u.is_deleted = FALSE AND
				u.username = $1 
			  LIMIT 1`

	row := db.QueryRow(query, username)
	err := row.Scan(
		&result.Id, &result.Username, &result.Password, &result.UserProfileId,
		&result.Role)
	if err != nil {
		errType = errtype.InternalServerError(err)
	}
	return
}