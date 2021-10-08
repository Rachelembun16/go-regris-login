package entity

import (
	"html"
	"strings"

	"github.com/gofrs/uuid"
)

const (
	BukuTableName = "buku"
)

type Buku struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key" json:"id_buku"`
	Nama_buku string    `gorm:"type:varchar(200);not_null" json:"nama_buku"`
	ISBN      string    `gorm:"type:varchar(200);null" json:"isbn"`
	Harga     int       `gorm:"type:int;not_null" json:"harga"`
}

func NewBuku(id_buku uuid.UUID, nama_buku string, isbn string, harga int) *Buku {
	return &Buku{
		ID:        id_buku,
		Nama_buku: nama_buku,
		ISBN:      isbn,
		Harga:     harga,
	}
}

func (model *Buku) TableName() string {
	return BukuTableName
}

func (article *Buku) GenerateSlug() string {
	return html.EscapeString(strings.ToLower(strings.ReplaceAll(article.Nama_buku, " ", "-")))
}
