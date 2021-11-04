package models

type Products struct {
	ID          int     `json:"id"`
	Nama_produk string  `json:"nm_produk"`
	Modal       float64 `json:"modal"`
	Harga       float64 `json:"harga"`
	Kategori_id int
	Kategoris   Kategoris `gorm:"foreignKey:Kategori_id"`
	Branch_id   int
	Stores      Stores `gorm:"foreignKey:Branch_id"`
	
}
