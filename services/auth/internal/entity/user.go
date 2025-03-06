package entity

import (
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"pms.pkg/protobuf/dto"
	"pms.pkg/type/timestamp"
)

type User struct {
	ID        uuid.UUID           `db:"id"`
	Name      string              `db:"name"`
	Email     string              `db:"email"`
	Password  string              `db:"password"`
	AvatarIMG []byte              `db:"avatar_img"`
	CreatedAt timestamp.Timestamp `db:"created_at"`
	UpdatedAt timestamp.Timestamp `db:"updated_at"`
}

func (u *User) DTO() *dto.User {
	return &dto.User{
		Id:        u.ID.String(),
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: timestamppb.New(u.CreatedAt.Time),
		UpdatedAt: timestamppb.New(u.UpdatedAt.Time),
	}
}

func UserFromDTO(user *dto.User) User {
	return User{
		ID:        uuid.MustParse(user.Id),
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: timestamp.NewTimestamp(user.CreatedAt.AsTime()),
		UpdatedAt: timestamp.NewTimestamp(user.UpdatedAt.AsTime()),
	}
}
