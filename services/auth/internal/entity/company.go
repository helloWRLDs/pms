package entity

import (
	"github.com/google/uuid"
	"pms.pkg/type/timestamp"
)

type Company struct {
	ID        uuid.UUID           `db:"id"`
	Name      string              `db:"name"`
	Codename  string              `db:"codename"`
	CreatedAt timestamp.Timestamp `db:"created_at"`
	UpdatedAt timestamp.Timestamp `db:"updated_at"`
}
