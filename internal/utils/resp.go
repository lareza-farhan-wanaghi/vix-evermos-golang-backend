package utils

import (
	"fmt"
	"strconv"
	"tugas_akhir_example/internal/daos"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"
)

// ProdukToProdukResp parses the produk database data into produk respond data
func ProdukToProdukResp(data *daos.Produk) (res *dto.ProdukResp, err error) {
	hargaKonsumen, err := strconv.Atoi(data.HargaKonsumen)
	if err != nil {
		helper.Logger(GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return nil, err
	}

	hargaReseller, err := strconv.Atoi(data.HargaReseller)
	if err != nil {
		helper.Logger(GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", err.Error()))
		return nil, err
	}

	photos := []dto.FotoProdukResp{}
	for _, fotoProduk := range data.FotoProduks {
		photos = append(photos, dto.FotoProdukResp{
			ID:       fotoProduk.ID,
			ProdukId: fotoProduk.IdProduk,
			Url:      fotoProduk.Url,
		})
	}

	res = &dto.ProdukResp{
		Id:            data.ID,
		NamaProduk:    data.NamaProduk,
		Slug:          data.Slug,
		HargaKonsumen: hargaKonsumen,
		HargaReseller: hargaReseller,
		Stok:          data.Stok,
		Deskripsi:     data.Deskripsi,
		Toko: dto.TokoResp{
			ID:       data.Toko.ID,
			NamaToko: data.Toko.NamaToko,
			UrlFoto:  data.Toko.UrlFoto,
		},
		Category: dto.CategoryResp{
			ID:           data.Category.ID,
			NamaCategory: data.Category.NamaCategory,
		},
		Photo: photos,
	}

	return res, nil
}

func ProdukArrayToAllProdukResp(data []*daos.Produk) (res *dto.AllProdukResp, err error) {
	res = &dto.AllProdukResp{
		Data: []*dto.ProdukResp{},
	}
	for _, v := range data {
		produkResp, err := ProdukToProdukResp(v)
		if err != nil {
			return nil, err
		}

		res.Data = append(res.Data, produkResp)
	}

	return res, nil
}

// TrxToTrxResp parses the trx database data into trx respond data
func TrxToTrxResp(data *daos.Trx) (res *dto.TrxResp, err error) {
	detailTrxResps := []*dto.DetailTrxResp{}
	for _, v := range data.DetailTrxs {
		logProdukResp, err := LogProdukToLogProdukResp(v.LogProduk)
		if err != nil {
			return nil, err
		}

		detailTrxResps = append(detailTrxResps, &dto.DetailTrxResp{
			LogProduk: logProdukResp,
			Toko: &dto.TokoResp{
				ID:       v.LogProduk.Toko.ID,
				NamaToko: v.LogProduk.Toko.NamaToko,
				UrlFoto:  v.LogProduk.Toko.UrlFoto,
			},
			Kuantitas:  v.Kuantitas,
			HargaTotal: v.HargaTotal,
		})
	}

	res = &dto.TrxResp{
		Id:          data.ID,
		HargaTotal:  data.HargaTotal,
		KodeInvoice: data.KodeInvoice,
		MethodBayar: data.MethodBayar,
		AlamatKirim: &dto.AlamatResp{
			Id:           data.Alamat.ID,
			JudulAlamat:  data.Alamat.JudulAlamat,
			NamaPenerima: data.Alamat.NamaPenerima,
			Notelp:       data.Alamat.Notelp,
			DetailAlamat: data.Alamat.DetailAlamat,
		},
		DetailTrxes: detailTrxResps,
	}
	return res, nil
}

// TrxArrayToAllTrxResp parses the trx database data into alltrx respond data
func TrxArrayToAllTrxResp(data []*daos.Trx) (res *dto.AllTrxResp, err error) {
	res = &dto.AllTrxResp{
		Data: []*dto.TrxResp{},
	}

	for _, v := range data {
		trxResp, err := TrxToTrxResp(v)
		if err != nil {
			return nil, err
		}
		res.Data = append(res.Data, trxResp)
	}

	return res, nil
}

// LogProdukToLogProdukResp parses the logproduk database data into logproduk respond data
func LogProdukToLogProdukResp(data *daos.LogProduk) (res *dto.LogProdukResp, err error) {
	hargaReseller, err := strconv.Atoi(data.HargaReseller)
	if err != nil {
		return nil, err
	}

	hargaKonsumen, err := strconv.Atoi(data.HargaKonsumen)
	if err != nil {
		return nil, err
	}

	photos := []*dto.FotoProdukResp{}
	for _, v := range data.Produk.FotoProduks {
		photos = append(photos, &dto.FotoProdukResp{
			ID:       v.ID,
			ProdukId: v.IdProduk,
			Url:      v.Url,
		})
	}

	res = &dto.LogProdukResp{
		Id:            data.ID,
		NamaProduk:    data.NamaProduk,
		Slug:          data.Slug,
		HargaReseller: hargaReseller,
		HargaKonsumen: hargaKonsumen,
		Deskripsi:     data.Deskripsi,
		Toko: &dto.TokoResp{
			ID:       data.Toko.ID,
			NamaToko: data.Toko.NamaToko,
			UrlFoto:  data.Toko.UrlFoto,
		},
		Category: &dto.CategoryResp{
			ID:           data.Category.ID,
			NamaCategory: data.Category.NamaCategory,
		},
		Photos: photos,
	}
	return res, nil
}

// TokoToTokoResp parses the toko database data into toko respond data
func TokoToTokoResp(data *daos.Toko) (res *dto.TokoResp) {
	res = &dto.TokoResp{
		ID:       data.ID,
		NamaToko: data.NamaToko,
		UrlFoto:  data.UrlFoto,
		UserId:   data.IdUser,
	}

	return res
}

// TokoArrayToAllTokoResp parses the toko database data into alltoko respond data
func TokoArrayToAllTokoResp(data []*daos.Toko) (res *dto.AllTokoResp) {
	res = &dto.AllTokoResp{
		Data: []*dto.TokoResp{},
	}
	for _, v := range data {
		res.Data = append(res.Data, TokoToTokoResp(v))
	}

	return res
}

// UserToLoginResp parses the user database data into login respond data
func UserToLoginResp(data *daos.User) (res *dto.LoginResp) {
	res = &dto.LoginResp{
		ID:           data.ID,
		Nama:         data.Nama,
		Notelp:       data.Notelp,
		TanggalLahir: DateToString(data.TanggalLahir),
		Tentang:      data.Tentang,
		Pekerjaan:    data.Pekerjaan,
		Email:        data.Email,
	}

	return res
}

// CatergoryToCategoryResp parses the category database data into category respond data
func CatergoryToCategoryResp(data *daos.Category) (res *dto.CategoryResp) {
	res = &dto.CategoryResp{
		ID:           data.ID,
		NamaCategory: data.NamaCategory,
	}

	return res
}

// UserToUserResp parses the user database data into user respond data
func UserToUserResp(data *daos.User) (res *dto.UserResp) {
	alamats := []*dto.AlamatResp{}
	for _, v := range data.Alamats {
		alamats = append(alamats, &dto.AlamatResp{
			Id:           v.ID,
			JudulAlamat:  v.JudulAlamat,
			NamaPenerima: v.NamaPenerima,
			Notelp:       v.Notelp,
			DetailAlamat: v.DetailAlamat,
		})
	}

	res = &dto.UserResp{
		Id:           data.ID,
		Nama:         data.Nama,
		Notelp:       data.Notelp,
		TanggalLahir: DateToString(data.TanggalLahir),
		JenisKelamin: data.JenisKelamin,
		Tentang:      data.Tentang,
		Pekerjaan:    data.Pekerjaan,
		Email:        data.Email,
		Alamats:      alamats,
	}

	return res
}

// AlamatToAlamatResp parses the alamat database data into alamat respond data
func AlamatToAlamatResp(data *daos.Alamat) (res *dto.AlamatResp) {
	res = &dto.AlamatResp{
		Id:           data.ID,
		JudulAlamat:  data.JudulAlamat,
		NamaPenerima: data.NamaPenerima,
		Notelp:       data.Notelp,
		DetailAlamat: data.DetailAlamat,
	}

	return res
}
