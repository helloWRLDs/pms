package entity

import (
	"pms.pkg/consts"
	"pms.pkg/type/timestamp"
)

type Role struct {
	Name         string               `db:"name"`
	Persmissions consts.PermissionSet `db:"permissions"`
	CreatedAt    timestamp.Timestamp  `db:"created_at"`
	UpdatedAt    timestamp.Timestamp  `db:"updated_at"`
	CompanyID    *string              `db:"company_id"`
}
