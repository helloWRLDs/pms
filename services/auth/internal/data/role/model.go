package roledata

import (
	"time"

	"github.com/lib/pq"
	"google.golang.org/protobuf/types/known/timestamppb"
	"pms.pkg/transport/grpc/dto"
	"pms.pkg/utils"
)

type Role struct {
	Name        string         `db:"name"`
	Permissions pq.StringArray `db:"permissions"`
	CreatedAt   time.Time      `db:"created_at"`
	UpdatedAt   *time.Time     `db:"updated_at"`
	CompanyID   *string        `db:"company_id"`
}

func (r *Role) DTO() *dto.Role {
	return &dto.Role{
		Name:        r.Name,
		Permissions: r.Permissions,
		CreatedAt:   timestamppb.New(r.CreatedAt),
		UpdatedAt:   timestamppb.New(utils.Value(r.UpdatedAt)),
		CompanyId:   utils.Value(r.CompanyID),
	}
}

// type RoleFilter struct {
// 	list.Pagination
// 	list.Date
// 	list.Order
// 	CompanyID   string `json:"company_id"`
// 	CompanyName string `json:"company_name"`
// 	Name        string `json:"name"`
// }
