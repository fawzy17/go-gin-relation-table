package models

type Kelas struct {
	Kode   string `gorm:"primarykey;type:varchar(5);not null" json:"kode"`
	Name   string `gorm:"varchar(30)" json:"name"`
	Jumlah int    `gorm:"int(5)" json:"jumlah"`
}
type KelasResponse struct {
	KodeKelas string `gorm:"type:varchar(5);not null" json:"kode_kelas"`
	Nama      string `gorm:"varchar(30)" json:"nama"`
	Jumlah    int    `gorm:"int(5)" json:"jumlah"`
}
