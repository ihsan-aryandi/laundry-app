package entity

import "database/sql"

type AuditRepo struct {
	CreatedAt sql.NullTime
	CreatedBy sql.NullInt64
	UpdatedAt sql.NullTime
	UpdatedBy sql.NullInt64
	IsDeleted sql.NullBool
}