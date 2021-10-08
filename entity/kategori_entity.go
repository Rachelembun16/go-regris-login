package entity

import (
	"github.com/gofrs/uuid"
)

const (
	KategoriTableName = "kategori"
)

type Kategori struct {
	ID            uuid.UUID `gorm:"type:uuid;primary_key" json:"id_kategori"`
	Nama_kategori string    `gorm:"type:varchar(200);not_null" json:"nama_kategori"`
}

func NewKategori(id_admin uuid.UUID, nama_kategori string) *Kategori {
	return &Kategori{
		ID:            id_admin,
		Nama_kategori: nama_kategori,
	}
}

func (model *Kategori) TableName() string {
	return KategoriTableName
}
