package repo

import (
	"laundry/errtype"
	"laundry/warehouse"
)

type genderRepo struct {
	TableName string
}

func NewGenderRepo() *genderRepo {
	return &genderRepo{TableName: "genders"}
}

func (g *genderRepo) GetGenderCountById(id int64) (result int, errType *errtype.Error) {
	db := warehouse.Conn.DB
	query := `SELECT 
				COUNT(id) 
			  FROM ` + g.TableName + `
			  WHERE 
				is_deleted = FALSE AND 
				id = $1
			  LIMIT 1`

	row := db.QueryRow(query, id)
	err := row.Scan(&result)
	if err != nil {
		errType = errtype.InternalServerError(err)
		return
	}

	return
}