package models

import (
	"github.com/sev-2/raiden/pkg/db"
)

type Dokter struct {
	db.ModelBase
	IdDokter     int32   `json:"id_dokter,omitempty" column:"name:id_dokter;type:integer;primaryKey;nullable:false"`
	NamaDokter   *string `json:"nama_dokter,omitempty" column:"name:nama_dokter;type:varchar;nullable"`
	Spesialisasi *string `json:"spesialisasi,omitempty" column:"name:spesialisasi;type:varchar;nullable"`
	NomorTelepon *string `json:"nomor_telepon,omitempty" column:"name:nomor_telepon;type:varchar;nullable"`
	Email        *string `json:"email,omitempty" column:"name:email;type:varchar;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"dokter" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	DoktersThroughJadwalPraktikIdDokter  []*Dokter        `json:"dokters_through_jadwal_praktik_id_dokter,omitempty" join:"joinType:manyToMany;through:jadwal_praktik;sourcePrimaryKey:id_dokter;sourceForeignKey:id_dokter;targetPrimaryKey:id_dokter;targetForeign:id_dokter"`
	JadwalPraktikIds                     []*JadwalPraktik `json:"jadwal_praktik_ids,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id_dokter;foreignKey:id_dokter"`
	RuangansThroughJadwalPraktikIdDokter []*Ruangan       `json:"ruangans_through_jadwal_praktik_id_dokter,omitempty" join:"joinType:manyToMany;through:jadwal_praktik;sourcePrimaryKey:id_dokter;sourceForeignKey:id_dokter;targetPrimaryKey:id_ruangan;targetForeign:id_dokter"`
}
