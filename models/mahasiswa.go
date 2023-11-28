package models

import "gorm.io/gorm"

type Mahasiswa struct {
	gorm.Model
	Nama      string `gorm:"varchar(30), not null" json:"nama"`
	Npm       string `gorm:"varchar(30), not null" json:"npm"`
	KelasKode string `gorm:"type:varchar(5); not null; index" json:"kelas_kode"`
	Kelas     Kelas  `gorm:"foreignKey:KelasKode;references:Kode;onUpdate:CASCADE,onDelete:CASCADE" json:"kelas"`

	// KelasID int    `gorm:"int(3), not null" json:"kelas_id"`
}

type MahasiswaResponse struct {
	ID        uint   `gorm:"primarykey"`
	Nama      string `gorm:"varchar(30), not null" json:"nama"`
	Npm       string `gorm:"varchar(30), not null" json:"npm"`
	KelasKode string `gorm:"varchar(5), not null" json:"kelas_kode"`
	Kelas     Kelas  `gorm:"foreignKey:KelasKode" json:"kelas"`
}
