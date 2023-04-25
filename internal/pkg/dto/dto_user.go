package dto

type AlamatResp struct {
	Id           uint   `json:"id"`
	JudulAlamat  string `json:"judul_alamat"`
	NamaPenerima string `json:"nama_penerima"`
	Notelp       string `json:"no_telp"`
	DetailAlamat string `json:"detail_alamat"`
}

type AlamatFilter struct {
	JudulAlamat string `query:"judul_alamat"`
}

type AlamatCreateReq struct {
	JudulAlamat  string `json:"judul_alamat" validate:"required"`
	NamaPenerima string `json:"nama_penerima" validate:"required"`
	Notelp       string `json:"no_telp" validate:"required"`
	DetailAlamat string `json:"detail_alamat" validate:"required"`
}

type AlamatUpdateReq struct {
	JudulAlamat  string `json:"judul_alamat,omitempty"`
	NamaPenerima string `json:"nama_penerima,omitempty"`
	Notelp       string `json:"no_telp,omitempty"`
	DetailAlamat string `json:"detail_alamat,omitempty"`
}

type UserResp struct {
	Id           uint          `json:"id"`
	Nama         string        `json:"nama"`
	Notelp       string        `json:"no_telp"`
	TanggalLahir string        `json:"tanggal_lahir"`
	JenisKelamin string        `json:"jenis_kelamin"`
	Tentang      string        `json:"tentang"`
	Pekerjaan    string        `json:"pekerjaan"`
	Email        string        `json:"email"`
	IdProvinsi   *ProvinceResp `json:"id_provinsi"`
	IdKota       *CityResp     `json:"id_kota"`
	Alamats      []*AlamatResp `json:"alamat"`
}

type UserUpdateReq struct {
	Nama         string `json:"nama,omitempty"`
	KataSandi    string `json:"kata_sandi,omitempty"`
	Notelp       string `json:"no_telp,omitempty"`
	TanggalLahir string `json:"tanggal_lahir,omitempty"`
	JenisKelamin string `json:"jenis_kelamin,omitempty"`
	Tentang      string `json:"tentang,omitempty"`
	Pekerjaan    string `json:"pekerjaan,omitempty"`
	Email        string `json:"email,omitempty"`
	IdProvinsi   string `json:"id_provinse,omitempty"`
	IdKota       string `json:"id_kota,omitempty"`
}
