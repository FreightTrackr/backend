package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Transaksi struct {
	ID                primitive.ObjectID `json:"id" bson:"_id,omitempty"`                    // MongoDB ID
	No_Resi           string             `json:"no_resi" bson:"no_resi"`                     // Nomor Resi
	Layanan           string             `json:"layanan" bson:"layanan"`                     // Layanan pengiriman
	Isi_Kiriman       string             `json:"isi_kiriman" bson:"isi_kiriman"`             // Isi kiriman
	Nama_Pengirim     string             `json:"nama_pengirim" bson:"nama_pengirim"`         // Nama pengirim
	Alamat_Pengirim   string             `json:"alamat_pengirim" bson:"alamat_pengirim"`     // Alamat pengirim
	Kode_Pos_Pengirim int                `json:"kode_pos_pengirim" bson:"kode_pos_pengirim"` // Kode pos pengirim
	Kota_Asal         string             `json:"kota_asal" bson:"kota_asal"`                 // Kota asal pengirim
	Nama_Penerima     string             `json:"nama_penerima" bson:"nama_penerima"`         // Nama penerima
	Alamat_Penerima   string             `json:"alamat_penerima" bson:"alamat_penerima"`     // Alamat penerima
	Kode_Pos_Penerima int                `json:"kode_pos_penerima" bson:"kode_pos_penerima"` // Kode pos penerima
	Kota_Tujuan       string             `json:"kota_tujuan" bson:"kota_tujuan"`             // Kota tujuan pengiriman
	Berat_Kiriman     float64            `json:"berat_kiriman" bson:"berat_kiriman"`         // Berat kiriman
	Volumetrik        float64            `json:"volumetrik" bson:"volumetrik"`               // Volumetrik kiriman
	Nilai_Barang      int                `json:"nilai_barang" bson:"nilai_barang"`           // Nilai barang yang dikirim
	Biaya_Dasar       int                `json:"biaya_dasar" bson:"biaya_dasar"`             // Biaya dasar pengiriman
	Biaya_Pajak       int                `json:"biaya_pajak" bson:"biaya_pajak"`             // Biaya pajak
	Biaya_Asuransi    int                `json:"biaya_asuransi" bson:"biaya_asuransi"`       // Biaya asuransi
	Total_Biaya       int                `json:"total_biaya" bson:"total_biaya"`             // Total biaya pengiriman
	Tanggal_Kirim     primitive.DateTime `json:"tanggal_kirim" bson:"tanggal_kirim"`         // Tanggal kirim
	Tanggal_Terima    primitive.DateTime `json:"tanggal_terima" bson:"tanggal_terima"`       // Tanggal terima
	Tanggal_Tenggat   primitive.DateTime `json:"tanggal_tenggat" bson:"tanggal_tenggat"`     // Tanggal tenggat
	Status            string             `json:"status" bson:"status"`                       // Status pengiriman
	Tipe_Cod          string             `json:"tipe_cod" bson:"tipe_cod"`                   // Tipe COD
	Status_Cod        string             `json:"status_cod" bson:"status_cod"`               // Status COD
	Sla               int                `json:"sla" bson:"sla"`                             // SLA dalam hari
	Aktual_Sla        int                `json:"aktual_sla" bson:"aktual_sla"`               // Aktual SLA dalam hari
	Status_Sla        bool               `json:"status_sla" bson:"status_sla"`               // Status SLA tercapai atau tidak
	No_Pend_Kirim     string             `json:"no_pend_kirim" bson:"no_pend_kirim"`         // Nomor pend kirim
	No_Pend_Terima    string             `json:"no_pend_terima" bson:"no_pend_terima"`       // Nomor pend terima
	Kode_Pelanggan    string             `json:"kode_pelanggan" bson:"kode_pelanggan"`       // Kode pelanggan
	Created_By        struct {
		Username string `json:"username" bson:"username"` // Username pembuat transaksi
	} `json:"created_by" bson:"created_by"`
	ID_History string `json:"id_history" bson:"id_history"` // ID history (otomatis generate)
}
