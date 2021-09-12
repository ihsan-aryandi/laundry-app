package repo

import (
	"laundry/errtype"
	"laundry/warehouse"
)

type roleRepo struct {
	TableName string
}

func NewRoleRepo() *roleRepo {
	return &roleRepo{TableName: "roles"}
}

func (r *roleRepo) GetIdByRoleName(roleName string) (id int64, errType *errtype.Error) {
	db := warehouse.Conn.DB
	query := `SELECT 
				id 
			  FROM ` + r.TableName + ` 
			  WHERE 
				is_deleted = FALSE AND 
				role = $1`

	err := db.QueryRow(query, roleName).Scan(&id)
	if err != nil {
		errType = errtype.InternalServerError(err)
	}
	return
}