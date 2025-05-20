package projectdata

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
	"pms.pkg/consts"
	"pms.pkg/transport/grpc/dto"
	"pms.pkg/utils"
)

type Project struct {
	ID          string               `db:"id"`
	Title       string               `db:"title"`
	CodeName    string               `db:"codename"`
	CodePrefix  *string              `db:"code_prefix"`
	Description string               `db:"description"`
	Status      consts.ProjectStatus `db:"status"`
	CompanyID   string               `db:"company_id"`
	CreatedAt   time.Time            `db:"created_at"`
	UpdatedAt   *time.Time           `db:"updated_at"`
	Progress    *int                 `db:"progress"`
}

func (p *Project) DTO() *dto.Project {
	return &dto.Project{
		Id:          p.ID,
		Title:       p.Title,
		CodeName:    p.CodeName,
		Description: p.Description,
		Status:      string(p.Status),
		CodePrefix:  utils.Value(p.CodePrefix),
		CompanyId:   p.CompanyID,
		CreatedAt:   timestamppb.New(p.CreatedAt),
		UpdatedAt:   timestamppb.New(utils.Value(p.UpdatedAt)),
	}
}

func Entity(dto *dto.Project) *Project {
	return &Project{
		ID:          dto.Id,
		Title:       dto.Title,
		CodeName:    dto.CodeName,
		Description: dto.Description,
		Status:      consts.ProjectStatus(dto.Status),
		CompanyID:   dto.CompanyId,
		CodePrefix:  utils.Ptr(dto.CodePrefix),
		CreatedAt:   dto.CreatedAt.AsTime(),
		UpdatedAt:   utils.Ptr(dto.UpdatedAt.AsTime()),
	}
}
