package models

import (
	"github.com/google/uuid"
	"pms.pkg/consts"
	"pms.pkg/type/timestamp"
)

type Project struct {
	ID          uuid.UUID            `db:"id"`
	Title       string               `db:"title"`
	Description string               `db:"description"`
	Status      consts.ProjectStatus `db:"status"`
	CompanyID   string               `db:"company_id"`
	CreatedAt   timestamp.Timestamp  `db:"created_at"`
	UpdatedAt   timestamp.Timestamp  `db:"updated_at"`
	Progress    *int                 `db:"progress"`
}
