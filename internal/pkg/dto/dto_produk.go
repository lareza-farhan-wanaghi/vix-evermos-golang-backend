package dto

type AllProdukResp struct {
	Page  int           `json:"page"`
	Limit int           `json:"limit"`
	Data  []*ProdukResp `json:"data"`
}

type ProdukResp struct {
	Id            uint             `json:"id"`
	NamaProduk    string           `json:"nama_produk"`
	Slug          string           `json:"slug"`
	HargaReseller int              `json:"harga_reseler"`
	HargaKonsumen int              `json:"harga_konsumen"`
	Stok          int              `json:"stok"`
	Deskripsi     string           `json:"deskripsi"`
	Toko          TokoResp         `json:"toko"`
	Category      CategoryResp     `json:"category"`
	Photo         []FotoProdukResp `json:"photos"`
}

type ProdukFilter struct {
	NamaProduk string `query:"nama_produk"`
	Limit      int    `query:"limit"`
	Page       int    `query:"page"`
	CategoryId uint   `query:"category_id"`
	TokoId     uint   `query:"toko_id"`
	MaxHarga   int    `query:"max_harga"`
	MinHarga   int    `query:"min_harga"`
}

type ProdukCreateReq struct {
	NamaProduk    string `form:"nama_produk" validate:"required"`
	CategoryId    string `form:"category_id" validate:"required"`
	HargaReseller string `form:"harga_reseller" validate:"required"`
	HargaKonsumen string `form:"harga_konsumen" validate:"required"`
	Stok          string `form:"stok" validate:"required"`
	Deskripsi     string `form:"deskripsi" validate:"required"`
}

type ProdukUpdateReq struct {
	NamaProduk    string `form:"nama_produk,omitempty"`
	CategoryId    string `form:"category_id,omitempty"`
	HargaReseller string `form:"harga_reseller,omitempty"`
	HargaKonsumen string `form:"harga_konsumen,omitempty"`
	Stok          string `form:"stok,omitempty"`
	Deskripsi     string `form:"deskripsi,omitempty"`
}

type FotoProdukResp struct {
	ID       uint   `json:"id"`
	ProdukId uint   `json:"product_id"`
	Url      string `json:"url"`
}
