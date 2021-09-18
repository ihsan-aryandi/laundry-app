package entity

import "database/sql"

type RegistrationLinkBody struct {
	Email string `json:"email"`
}

type RegistrationLinkRepo struct {
	Id        sql.NullInt64
	Email     sql.NullString
	Code      sql.NullString
	ExpiredAt sql.NullTime
	Type      sql.NullString
	LaundryId sql.NullInt64
	AuditRepo
}
