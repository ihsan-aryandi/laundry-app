package entity

import "database/sql"

type UserProfileRepo struct {
	Id              sql.NullInt64
	Name            sql.NullString
	GenderId        sql.NullInt64
	Address         sql.NullString
	Photo           sql.NullString
	Bio             sql.NullString
	TotalReputation sql.NullInt64
	RoleId          sql.NullInt64
	CreatedAt       sql.NullTime
	CreatedBy       sql.NullInt64
	UpdatedAt       sql.NullTime
	UpdatedBy       sql.NullInt64
	IsDeleted       sql.NullInt64
	Role            sql.NullString
}
