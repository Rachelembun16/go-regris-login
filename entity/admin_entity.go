package entity

import (
	"github.com/gofrs/uuid"
)

const (
	AdminTableName = "admin"
)

type Admin struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key" json:"id_admin"`
	Nama_admin  string    `gorm:"type:varchar(200);not_null" json:"nama_admin"`
	Email_admin string    `gorm:"type:varchar(200);null" json:"email_admin"`
	Username    string    `gorm:"type:varchar(200);not_null" json:"username"`
	Password    string    `gorm:"type:varchar(200);not_null" json:"password"`
}

func NewAdmin(id_admin uuid.UUID, nama_admin, email_admin, username, password string) *Admin {
	return &Admin{
		ID:          id_admin,
		Nama_admin:  nama_admin,
		Email_admin: email_admin,
		Username:    username,
		Password:    password,
	}
}

func (model *Admin) TableName() string {
	return AdminTableName
}
