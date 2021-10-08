package entity

import (
	"github.com/gofrs/uuid"
)

const (
	SiswaTableName = "Siswa"
)

type Siswa struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key" json:"id_Siswa"`
	Nama_Siswa  string    `gorm:"type:varchar(200);not_null" json:"nama_Siswa"`
	Email_Siswa string    `gorm:"type:varchar(200);null" json:"email_Siswa"`
	Username    string    `gorm:"type:varchar(200);not_null" json:"username"`
	Password    string    `gorm:"type:varchar(200);not_null" json:"password"`
}

func NewSiswa(id_Siswa uuid.UUID, nama_Siswa, email_Siswa, username, password string) *Siswa {
	return &Siswa{
		ID:          id_Siswa,
		Nama_Siswa:  nama_Siswa,
		Email_Siswa: email_Siswa,
		Username:    username,
		Password:    password,
	}
}

func (model *Siswa) TableName() string {
	return SiswaTableName
}
