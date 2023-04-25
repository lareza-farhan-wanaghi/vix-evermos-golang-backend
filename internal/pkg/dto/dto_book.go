package dto

type BookFilter struct {
	Title string `query:"title"`
	Limit int    `query:"limit"`
	Page  int    `query:"page"`
}

type BookReqCreate struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Author      string `json:"author" validate:"required"`
}

type BookReqUpdate struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Author      string `json:"author,omitempty"`
}

type BookResp struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
}
