package service

import (
	"database/sql"
	uuid "github.com/nu7hatch/gouuid"
	"laundry/constant"
	"laundry/entity"
	"laundry/errtype"
	"time"
)

type registrationLinkSvc struct {}

func NewRegistrationLinkSvc() *registrationLinkSvc {
	return &registrationLinkSvc{}
}

func (rl *registrationLinkSvc) CreateRegistrationLinkLaundry(body entity.RegistrationLinkBody) (link string, errType *errtype.Error) {
	/*
		Create new transactional instance
	*/
	tx, now, errType := createTransactional()
	if errType != nil {
		return
	}

	/*
		Generate code
	*/
	code, err := uuid.NewV4()
	if err != nil {
		errType = errtype.InternalServerError(err)
		return
	}

	/*
		Find existing email
	*/


	expiredAt := now.Add(time.Hour * 24)
	repoData := entity.RegistrationLinkRepo{
		Email:     sql.NullString{String: body.Email},
		ExpiredAt: sql.NullTime{Time: expiredAt},
		Code: 	   sql.NullString{String: code.String()},
		Type:      sql.NullString{String: constant.RegistrationLinkLaundryType},
		LaundryId: sql.NullInt64{Int64: 0},
		AuditRepo: entity.AuditRepo{
			CreatedAt: sql.NullTime{Time: now},
			UpdatedAt: sql.NullTime{Time: now},
		},
	}

	registrationLinkRepo.CreateRegistrationLink(repoData)
}