package models

import (
	"github.com/sev-2/raiden/pkg/db"
)

type Items struct {
	db.ModelBase
	Id          int32   `json:"id,omitempty" column:"name:id;type:integer;primaryKey;nullable:false;default:nextval('items_id_seq')"`
	Name        *string `json:"name,omitempty" column:"name:name;type:text;nullable"`
	Description *string `json:"description,omitempty" column:"name:description;type:text;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"items" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`
}
