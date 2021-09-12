package entity

import (
	"database/sql"
	"time"
)

type User struct {
	Id            int64     `json:"id"`
	Username      string    `json:"username"`
	Password      string    `json:"password"`
	UserProfileId int64     `json:"-"`
	CreatedAt     time.Time `json:"created_at"`
	CreatedBy     int64     `json:"created_by"`
	UpdatedAt     time.Time `json:"updated_at"`
	UpdatedBy     int64     `json:"updated_by"`
}

type UserBody struct {
	Name          string `json:"name" rh_label:"Name" rh_valid:"required,alpha,max=256"`
	Username      string `json:"username" rh_label:"Username" rh_valid:"required,min=5,max=50"`
	Password      string `json:"password" rh_label:"Password" rh_valid:"required,min=7,max=256"`
	GenderId      int64  `json:"gender_id" rh_label:"Gender" rh_valid:"required"`
	Address       string `json:"address" rh_label:"Gender" rh_valid:"required"`
	UserProfileId int64  `json:"-"`
}

type UserRepo struct {
	Id            sql.NullInt64
	Username      sql.NullString
	Password      sql.NullString
	UserProfileId sql.NullInt64
	CreatedAt     sql.NullTime
	CreatedBy     sql.NullInt64
	UpdatedAt     sql.NullTime
	UpdatedBy     sql.NullInt64
	IsDeleted     sql.NullBool
	Role          sql.NullString
}
