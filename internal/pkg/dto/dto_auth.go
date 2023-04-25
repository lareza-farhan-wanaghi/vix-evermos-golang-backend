package dto

type AuthReqRegister struct {
	Nama         string `json:"nama" validate:"required"`
	KataSandi    string `json:"kata_sandi" validate:"required"`
	Notelp       string `json:"no_telp" validate:"required"`
	TanggalLahir string `json:"tanggal_lahir"`
	JenisKelamin string `json:"jenis_kelamin"`
	Tentang      string `json:"tentang"`
	Pekerjaan    string `json:"pekerjaan"`
	Email        string `json:"email" validate:"required"`
	IdProvinsi   string `json:"id_provinsi" validate:"required"`
	IdKota       string `json:"id_kota" validate:"required"`
}

type AuthReqLogin struct {
	Notelp    string `json:"no_telp" validate:"required"`
	KataSandi string `json:"kata_sandi" validate:"required"`
}

type AuthReqUpdate struct {
	Nama         string `json:"title,omitempty"`
	KataSandi    string `json:"kata_sandi,omitempty"`
	Notelp       string `json:"no_telp,omitempty"`
	TanggalLahir string `json:"tanggal_lahir,omitempty"`
	JenisKelamin string `json:"jenis_kelamin,omitempty"`
	Tentang      string `json:"tentang,omitempty"`
	Pekerjaan    string `json:"pekerjaan,omitempty"`
	Email        string `json:"email,omitempty"`
	IdProvinsi   string `json:"id_provinsi,omitempty"`
	IdKota       string `json:"id_kota,omitempty"`
}

type LoginResp struct {
	ID           uint          `json:"id"`
	Nama         string        `json:"nama"`
	Notelp       string        `json:"notelp"`
	TanggalLahir string        `json:"tanggal_lahir"`
	Tentang      string        `json:"tentang"`
	Pekerjaan    string        `json:"pekerjaan"`
	Email        string        `json:"email"`
	IdProvinsi   *ProvinceResp `json:"id_provinsi"`
	IdKota       *CityResp     `json:"id_kota"`
	Token        string        `json:"token"`
}
