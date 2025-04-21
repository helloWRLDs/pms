package entity

import (
	"time"

	"github.com/google/uuid"
)

// message Project {
//     string id = 1;
//     string title = 2;
//     string description = 3;
//     string status = 4;
//     string organization_id = 5;

//     google.protobuf.Timestamp created_at = 6;
//     google.protobuf.Timestamp updated_at = 7;
// }

type Project struct {
	ID          uuid.UUID `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	Status      string    `db:"status"`
	CompanyID   string    `db:"company_id"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
