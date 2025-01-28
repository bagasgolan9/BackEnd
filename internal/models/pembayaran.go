package models

import (
	"github.com/sev-2/raiden/pkg/db"
)

type Pembayaran struct {
	db.ModelBase
	IdPembayaran     int32    `json:"id_pembayaran,omitempty" column:"name:id_pembayaran;type:integer;primaryKey;nullable:false"`
	IdReservasi      *int32   `json:"id_reservasi,omitempty" column:"name:id_reservasi;type:integer;nullable"`
	JumlahBayar      *float64 `json:"jumlah_bayar,omitempty" column:"name:jumlah_bayar;type:numeric;nullable"`
	StatusPembayaran string   `json:"status_pembayaran,omitempty" column:"name:status_pembayaran;type:varchar;nullable:false"`
	MetodePembayaran string   `json:"metode_pembayaran,omitempty" column:"name:metode_pembayaran;type:varchar;nullable:false"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"pembayaran" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	ReservasiId *Reservasi `json:"reservasi_id,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id_reservasi;foreignKey:id_reservasi"`
}
