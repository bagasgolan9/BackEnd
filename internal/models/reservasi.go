package models

import (
	"github.com/sev-2/raiden/pkg/db"
	"time"
)

type Reservasi struct {
	db.ModelBase
	IdReservasi      int32      `json:"id_reservasi,omitempty" column:"name:id_reservasi;type:integer;primaryKey;nullable:false"`
	IdPasien         *int32     `json:"id_pasien,omitempty" column:"name:id_pasien;type:integer;nullable"`
	IdJadwal         *int32     `json:"id_jadwal,omitempty" column:"name:id_jadwal;type:integer;nullable"`
	TanggalReservasi *time.Time `json:"tanggal_reservasi,omitempty" column:"name:tanggal_reservasi;type:timestamp;nullable"`
	StatusReservasi  string     `json:"status_reservasi,omitempty" column:"name:status_reservasi;type:varchar;nullable:false"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"reservasi" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	JadwalPraktikId *JadwalPraktik `json:"jadwal_praktik_id,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id_jadwal;foreignKey:id_jadwal"`
	PasienId        *Pasien        `json:"pasien_id,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id_pasien;foreignKey:id_pasien"`
	PembayaranIds   []*Pembayaran  `json:"pembayaran_ids,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id_reservasi;foreignKey:id_reservasi"`
}
