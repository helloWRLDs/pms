package domain

import (
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"pms.auth/internal/domain/password"
	"pms.pkg/protobuf/dto"
	"pms.pkg/type/timestamp"
)

type User struct {
	ID        uuid.UUID           `db:"id"`
	FullName  string              `db:"full_name"`
	Email     string              `db:"email"`
	Password  password.Password   `db:"password"`
	CreatedAt timestamp.Timestamp `db:"created_at"`
	UpdatedAt timestamp.Timestamp `db:"updated_at"`
}

func (u User) DTO() dto.User {
	return dto.User{
		Id:        u.ID.String(),
		FullName:  u.FullName,
		Email:     u.Email,
		CreatedAt: timestamppb.New(u.CreatedAt.Time),
		UpdatedAt: timestamppb.New(u.CreatedAt.Time),
	}
}

func (u *User) FromDTO(dto *dto.User) {
	if dto == nil {
		return
	}
	if dto.Id != "" {
		id, err := uuid.Parse(dto.Id)
		if err == nil {
			u.ID = id
		}
	}
	if dto.CreatedAt != nil {
		u.CreatedAt = timestamp.WithFormat(dto.CreatedAt.AsTime(), timestamp.SQLITE_FORMAT)
	}
	if dto.UpdatedAt != nil {
		u.UpdatedAt = timestamp.WithFormat(dto.UpdatedAt.AsTime(), timestamp.SQLITE_FORMAT)
	}
	u.FullName = dto.FullName
	u.Email = dto.Email
	u.Password = password.Password(dto.Password)
}
