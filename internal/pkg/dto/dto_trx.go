package dto

type AllTrxResp struct {
	Page  int        `json:"page"`
	Limit int        `json:"limit"`
	Data  []*TrxResp `json:"data"`
}
type TrxResp struct {
	Id          uint             `json:"id"`
	HargaTotal  int              `json:"harga_total"`
	KodeInvoice string           `json:"kode_invoice"`
	MethodBayar string           `json:"method_bayar"`
	AlamatKirim *AlamatResp      `json:"alamat_kirim"`
	DetailTrxes []*DetailTrxResp `json:"detail_trx"`
}

type DetailTrxResp struct {
	LogProduk  *LogProdukResp `json:"product"`
	Toko       *TokoResp      `json:"toko"`
	Kuantitas  int            `json:"kuantitas"`
	HargaTotal int            `json:"harga_total"`
}

type TrxFilter struct {
	Search string `query:"search"`
	Limit  int    `query:"limit"`
	Page   int    `query:"page"`
}

type LogProdukResp struct {
	Id            uint              `json:"id"`
	NamaProduk    string            `json:"nama_produk"`
	Slug          string            `json:"slug"`
	HargaReseller int               `json:"harga_reseler"`
	HargaKonsumen int               `json:"harga_konsumen"`
	Deskripsi     string            `json:"deskripsi"`
	Toko          *TokoResp         `json:"toko"`
	Category      *CategoryResp     `json:"category"`
	Photos        []*FotoProdukResp `json:"photos"`
}

type TrxCreateReq struct {
	MethodBayar string                `json:"method_bayar"`
	AlamatKirim uint                  `json:"alamat_kirim"`
	DetailTrxes []*DetailTrxCreateReq `json:"detail_trx"`
}

type DetailTrxCreateReq struct {
	ProductId uint `json:"product_id"`
	Kuantitas int  `json:"kuantitas"`
}
