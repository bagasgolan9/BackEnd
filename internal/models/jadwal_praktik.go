package models

import (
	"github.com/sev-2/raiden/pkg/db"
)

type JadwalPraktik struct {
	db.ModelBase
	IdJadwal   int32   `json:"id_jadwal,omitempty" column:"name:id_jadwal;type:integer;primaryKey;nullable:false"`
	IdDokter   *int32  `json:"id_dokter,omitempty" column:"name:id_dokter;type:integer;nullable"`
	Tanggal    *string `json:"tanggal,omitempty" column:"name:tanggal;type:varchar;nullable"`
	Hari       *string `json:"hari,omitempty" column:"name:hari;type:varchar;nullable"`
	JamMulai   *string `json:"jam_mulai,omitempty" column:"name:jam_mulai;type:varchar;nullable"`
	JamSelesai *string `json:"jam_selesai,omitempty" column:"name:jam_selesai;type:varchar;nullable"`
	IdRuangan  *int32  `json:"id_ruangan,omitempty" column:"name:id_ruangan;type:integer;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"jadwal_praktik" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	DokterId                        *Dokter      `json:"dokter_id,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id_dokter;foreignKey:id_dokter"`
	PasiensThroughReservasiIdJadwal []*Pasien    `json:"pasiens_through_reservasi_id_jadwal,omitempty" join:"joinType:manyToMany;through:reservasi;sourcePrimaryKey:id_jadwal;sourceForeignKey:id_jadwal;targetPrimaryKey:id_pasien;targetForeign:id_jadwal"`
	ReservasiIds                    []*Reservasi `json:"reservasi_ids,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id_jadwal;foreignKey:id_jadwal"`
	RuanganId                       *Ruangan     `json:"ruangan_id,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id_ruangan;foreignKey:id_ruangan"`
}
