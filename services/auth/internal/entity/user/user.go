package userentity

import (
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"pms.pkg/protobuf/dto"
	"pms.pkg/type/timestamp"
)

type User struct {
	ID        uuid.UUID           `db:"id"`
	Email     string              `db:"email"`
	Password  string              `db:"password"`
	Name      string              `db:"name"`
	CreatedAt timestamp.Timestamp `db:"created_at"`
	UpdatedAt timestamp.Timestamp `db:"updated_at"`
}

func (u *User) DTO() *dto.User {
	return &dto.User{
		Id:        u.ID.String(),
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: timestamppb.New(u.CreatedAt.Time),
		UpdatedAt: timestamppb.New(u.UpdatedAt.Time),
	}
}

func FromDTO(user *dto.User) User {
	return User{
		ID:        uuid.MustParse(user.Id),
		Email:     user.Email,
		Name:      user.Name,
		Password:  user.Password,
		CreatedAt: timestamp.NewTimestamp(user.CreatedAt.AsTime()),
		UpdatedAt: timestamp.NewTimestamp(user.UpdatedAt.AsTime()),
	}
}
