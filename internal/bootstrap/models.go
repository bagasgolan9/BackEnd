// Code generated by raiden-cli; DO NOT EDIT.
package bootstrap

import (
	"github.com/sev-2/raiden/pkg/resource"
	"mobile/internal/models"
)

func RegisterModels() {
	resource.RegisterModels(
		&models.Dokter{},
		&models.Items{},
		&models.JadwalPraktik{},
		&models.Pasien{},
		&models.Pembayaran{},
		&models.Reservasi{},
		&models.Ruangan{},
	)
}
