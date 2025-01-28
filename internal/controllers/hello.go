package controllers

import (
	"mobile/internal/models"

	"github.com/sev-2/raiden"
)

type ScheduleController struct {
	raiden.ControllerBase
	Http  string `path:"/dokter" type:"rest"`
	Model models.Dokter
}
