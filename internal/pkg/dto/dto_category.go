package dto

type CategoryResp struct {
	ID           uint   `json:"id"`
	NamaCategory string `json:"nama_category"`
}

type CategoryCreateReq struct {
	NamaCategory string `json:"nama_category" validate:"required"`
}

type CategoryUpdateReq struct {
	NamaCategory string `json:"nama_category,omitempty"`
}
