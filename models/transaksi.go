package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Transaksi struct {
	ID                      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	No_Resi                 string             `json:"no_resi" bson:"no_resi"`
	Layanan                 string             `json:"layanan" bson:"layanan"`
	Isi_Kiriman             string             `json:"isi_kiriman" bson:"isi_kiriman"`
	Nama_Pengirim           string             `json:"nama_pengirim" bson:"nama_pengirim"`
	Alamat_Pengirim         string             `json:"alamat_pengirim" bson:"alamat_pengirim"`
	Kode_Pos_Pengirim       int                `json:"kode_pos_pengirim" bson:"kode_pos_pengirim"`
	Kota_Asal               string             `json:"kota_asal" bson:"kota_asal"`
	Nama_Penerima           string             `json:"nama_penerima" bson:"nama_penerima"`
	Alamat_Penerima         string             `json:"alamat_penerima" bson:"alamat_penerima"`
	Kode_Pos_Penerima       int                `json:"kode_pos_penerima" bson:"kode_pos_penerima"`
	Kota_Tujuan             string             `json:"kota_tujuan" bson:"kota_tujuan"`
	Berat_Kiriman           float64            `json:"berat_kiriman" bson:"berat_kiriman"`
	Volumetrik              float64            `json:"volumetrik" bson:"volumetrik"`
	Nilai_Barang            int                `json:"nilai_barang" bson:"nilai_barang"`
	Biaya_Dasar             int                `json:"biaya_dasar" bson:"biaya_dasar"`
	Biaya_Pajak             int                `json:"biaya_pajak" bson:"biaya_pajak"`
	Biaya_Asuransi          int                `json:"biaya_asuransi" bson:"biaya_asuransi"`
	Total_Biaya             int                `json:"total_biaya" bson:"total_biaya"`
	Tanggal_Kirim           primitive.DateTime `json:"tanggal_kirim" bson:"tanggal_kirim"`
	Tanggal_Antaran_Pertama primitive.DateTime `json:"tanggal_antaran_pertama,omitempty" bson:"tanggal_antaran_pertama,omitempty"`
	Tanggal_Terima          primitive.DateTime `json:"tanggal_terima,omitempty" bson:"tanggal_terima,omitempty"`
	Status                  string             `json:"status" bson:"status"`
	Tipe_Cod                string             `json:"tipe_cod" bson:"tipe_cod"`
	Status_Cod              string             `json:"status_cod,omitempty" bson:"status_cod,omitempty"`
	Sla                     int                `json:"sla" bson:"sla"`
	Aktual_Sla              int                `json:"aktual_sla,omitempty" bson:"aktual_sla,omitempty"`
	Status_Sla              *bool              `json:"status_sla,omitempty" bson:"status_sla,omitempty"`
	No_Pend_Kirim           string             `json:"no_pend_kirim" bson:"no_pend_kirim"`
	No_Pend_Terima          string             `json:"no_pend_terima" bson:"no_pend_terima"`
	Kode_Pelanggan          string             `json:"kode_pelanggan" bson:"kode_pelanggan"`
	Created_By              struct {
		Username string `json:"username" bson:"username"`
	} `json:"created_by" bson:"created_by"`
	ID_History string `json:"id_history" bson:"id_history"`
}
