package models

import (
	"github.com/sev-2/raiden/pkg/db"
)

type Ruangan struct {
	db.ModelBase
	IdRuangan   int32   `json:"id_ruangan,omitempty" column:"name:id_ruangan;type:integer;primaryKey;nullable:false"`
	NamaRuangan *string `json:"nama_ruangan,omitempty" column:"name:nama_ruangan;type:varchar;nullable"`
	Lokasi      *string `json:"lokasi,omitempty" column:"name:lokasi;type:varchar;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"ruangan" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	DoktersThroughJadwalPraktikIdRuangan []*Dokter        `json:"dokters_through_jadwal_praktik_id_ruangan,omitempty" join:"joinType:manyToMany;through:jadwal_praktik;sourcePrimaryKey:id_ruangan;sourceForeignKey:id_ruangan;targetPrimaryKey:id_dokter;targetForeign:id_ruangan"`
	JadwalPraktikIds                     []*JadwalPraktik `json:"jadwal_praktik_ids,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id_ruangan;foreignKey:id_ruangan"`
}
