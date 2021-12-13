package models

type Products struct {
	ID               uint `json:"id" gorm:"size:36;not null;uniqueIndex;primary_key"`
	Owner_id          int     `json:"owner_Id"`
	Kategori_id int	`json:"kategori_id"`
	Nama_produk string  `json:"nama_produk"`
	Modal       float64 `json:"modal"`
	Harga       float64 `json:"harga"`
	Stok       float64 `json:"stok"`
	Image string  `json:"image"`
	Branch_id   int `json:"branch_id"`
	Stores      Stores `gorm:"foreignKey:Branch_id"`
	Kategoris   Kategoris `gorm:"foreignKey:Kategori_id"`
	
}
