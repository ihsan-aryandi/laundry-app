package service

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"laundry/constant"
	"laundry/entity"
	"laundry/errtype"
	"time"
)

type registerSvc struct {}

func NewRegisterSvc() *registerSvc {
	return &registerSvc{}
}

func (as *registerSvc) Register(user entity.UserBody) (errType *errtype.Error) {
	tx, now, errType := createTransactional()
	if errType != nil {
		return
	}

	user.UserProfileId, errType = as.createUserProfile(tx, user, now)
	if errType != nil {
		return rollback(tx, errType)
	}

	errType = as.registerUser(tx, user, now)
	if errType != nil {
		return rollback(tx, errType)
	}

	return commit(tx, errType)
}

func (as *registerSvc) registerUser(tx *sql.Tx, user entity.UserBody, now time.Time) (errType *errtype.Error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errtype.InternalServerError(err)
	}

	data := entity.UserRepo{
		UserProfileId: sql.NullInt64{Int64: user.UserProfileId},
		Email:         sql.NullString{String: user.Email},
		Password:      sql.NullString{String: string(hashedPassword)},
		CreatedAt:     sql.NullTime{Time: now},
		UpdatedAt:     sql.NullTime{Time: now},
	}

	return authRepo.Register(tx, data)
}

func (as *registerSvc) createUserProfile(tx *sql.Tx, userBody entity.UserBody, now time.Time) (userProfileId int64, errType *errtype.Error) {
	roleId, errType := roleRepo.GetIdByRoleName(constant.DefaultUserRole)
	if errType != nil {
		return
	}

	data := entity.UserProfileRepo{
		Name:            sql.NullString{String: userBody.Name},
		GenderId:        sql.NullInt64{Int64: userBody.GenderId},
		Address:         sql.NullString{String: userBody.Address},
		Photo:           sql.NullString{String: constant.DefaultUserPhoto},
		TotalReputation: sql.NullInt64{Int64: constant.DefaultTotalReputation},
		RoleId:          sql.NullInt64{Int64: roleId},
		CreatedAt:       sql.NullTime{Time: now},
		UpdatedAt:       sql.NullTime{Time: now},
	}

	userProfileId, errType = userProfileService.CreateUserProfile(tx, data)
	return
}
