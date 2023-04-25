package dto

type ProvinceResp struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type CityResp struct {
	Id         string `json:"id"`
	ProvinceId string `json:"province_id"`
	Name       string `json:"name"`
}

type ProvinceFilter struct {
	Page   int    `query:"page"`
	Limit  int    `query:"limit"`
	Search string `query:"search"`
}
