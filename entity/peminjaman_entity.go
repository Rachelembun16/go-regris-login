package entity

import (
	"time"

	"github.com/gofrs/uuid"
)

const (
	PeminjamanTableName = "Peminjaman"
)

type Peminjaman struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key" json:"id_pinjam"`
	Tgl_pinjam  time.Time `gorm:"type:date;null" json:"tgl_pinjam"`
	Tgl_kembali time.Time `gorm:"type:date;null" json:"tgl_kembali"`
	Denda       int       `gorm:"type:int;null" json:"denda"`
}

func NewPeminjaman(id_pinjam uuid.UUID, tgl_pinjam time.Time, tgl_kembali time.Time, denda int) *Peminjaman {
	return &Peminjaman{
		ID:          id_pinjam,
		Tgl_pinjam:  tgl_pinjam,
		Tgl_kembali: tgl_kembali,
		Denda:       denda,
	}
}

func (model *Peminjaman) TableName() string {
	return PeminjamanTableName
}
