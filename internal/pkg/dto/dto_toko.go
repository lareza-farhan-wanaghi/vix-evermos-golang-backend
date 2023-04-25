package dto

type AllTokoResp struct {
	Page  int         `json:"page"`
	Limit int         `json:"limit"`
	Data  []*TokoResp `json:"data"`
}

type TokoResp struct {
	ID       uint   `json:"id"`
	NamaToko string `json:"nama_toko"`
	UrlFoto  string `json:"url_foto"`
	UserId   uint   `json:"user_id"`
}

type TokoUpdateReq struct {
	NamaToko string `form:"nama_toko"`
}

type TokoFilter struct {
	NamaToko string `query:"nama"`
	Limit    int    `query:"limit"`
	Page     int    `query:"page"`
}
