package repo

import (
	"database/sql"
	"laundry/entity"
	"laundry/errtype"
	"laundry/warehouse"
)

type userProfileRepo struct {
	TableName string
}

func NewUserProfileRepo() *userProfileRepo {
	return &userProfileRepo{TableName: "user_profiles"}
}

func (up *userProfileRepo) InsertUserProfile(tx *sql.Tx, data entity.UserProfileRepo) (id int64, errType *errtype.Error) {
    query := `INSERT INTO ` + up.TableName + `
				(name, gender_id, address, 
				 photo, total_reputation, role_id, 
				 created_at, created_by, updated_at, 
				 updated_by) 
			  VALUES
				($1, $2, $3,
				 $4, $5, $6,
				 $7, $8, $9,
				 $10)
			  RETURNING id`

    params := []interface{}{
    	data.Name.String, data.GenderId.Int64, data.Address.String,
    	data.Photo.String, data.TotalReputation.Int64, data.RoleId.Int64,
    	data.CreatedAt.Time, data.CreatedBy.Int64, data.UpdatedAt.Time,
    	data.UpdatedBy.Int64}

    err := tx.QueryRow(query, params...).Scan(&id)
    if err != nil {
    	errType = errtype.InternalServerError(err)
	}

	return
}

func (up *userProfileRepo) FindUserInfoById(id int64) (result entity.UserProfileRepo, errType *errtype.Error) {
	db := warehouse.Conn.DB

	roleTableName := NewRoleRepo().TableName
	query := "SELECT " +
			"	u.id, r.role " +
			"FROM " + up.TableName + " u " +
			"LEFT JOIN " + roleTableName + " r " +
			"	ON u.role_id = r.id " +
			"WHERE " +
			"	u.id = $1 AND " +
			"	u.is_deleted = FALSE"

	row := db.QueryRow(query, id)
	err := row.Scan(&result.Id, &result.Role)
	if err != nil {
		errType = errtype.InternalServerError(err)
	}

	return
}