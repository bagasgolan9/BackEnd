package models

import (
	"github.com/sev-2/raiden/pkg/db"
	"time"
)

type Pasien struct {
	db.ModelBase
	IdPasien     int32      `json:"id_pasien,omitempty" column:"name:id_pasien;type:integer;primaryKey;nullable:false"`
	NamaPasien   *string    `json:"nama_pasien,omitempty" column:"name:nama_pasien;type:varchar;nullable"`
	TanggalLahir *time.Time `json:"tanggal_lahir,omitempty" column:"name:tanggal_lahir;type:date;nullable"`
	Alamat       *string    `json:"alamat,omitempty" column:"name:alamat;type:text;nullable"`
	NomorTelepon *string    `json:"nomor_telepon,omitempty" column:"name:nomor_telepon;type:varchar;nullable"`
	Email        *string    `json:"email,omitempty" column:"name:email;type:varchar;nullable"`
	JenisKelamin string     `json:"jenis_kelamin,omitempty" column:"name:jenis_kelamin;type:varchar;nullable:false"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"pasien" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	JadwalPraktiksThroughReservasiIdPasien []*JadwalPraktik `json:"jadwal_praktiks_through_reservasi_id_pasien,omitempty" join:"joinType:manyToMany;through:reservasi;sourcePrimaryKey:id_pasien;sourceForeignKey:id_pasien;targetPrimaryKey:id_jadwal;targetForeign:id_pasien"`
	ReservasiIds                           []*Reservasi     `json:"reservasi_ids,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id_pasien;foreignKey:id_pasien"`
}
