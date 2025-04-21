package entity

import (
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"pms.pkg/transport/grpc/dto"
	"pms.pkg/type/timestamp"
	"pms.pkg/utils"
)

type User struct {
	ID        uuid.UUID           `db:"id"`
	Name      string              `db:"name"`
	Email     string              `db:"email"`
	Password  string              `db:"password"`
	AvatarIMG []byte              `db:"avatar_img"`
	Phone     *string             `db:"phone"`
	Bio       *string             `db:"bio"`
	CreatedAt timestamp.Timestamp `db:"created_at"`
	UpdatedAt timestamp.Timestamp `db:"updated_at"`
}

func (u *User) DTO() *dto.User {
	return &dto.User{
		Id:        u.ID.String(),
		Name:      u.Name,
		Email:     u.Email,
		Phone:     utils.Value(u.Phone),
		AvatarImg: u.AvatarIMG,
		Bio:       utils.Value(u.Bio),
		CreatedAt: timestamppb.New(u.CreatedAt.Time),
		UpdatedAt: timestamppb.New(u.UpdatedAt.Time),
	}
}

func UserFromDTO(user *dto.User) User {
	return User{
		ID:        uuid.MustParse(user.Id),
		Email:     user.Email,
		Name:      user.Name,
		Phone:     utils.Ptr(user.Phone),
		AvatarIMG: user.AvatarImg,
		Bio:       utils.Ptr(user.Bio),
		CreatedAt: timestamp.NewTimestamp(user.CreatedAt.AsTime()),
		UpdatedAt: timestamp.NewTimestamp(user.UpdatedAt.AsTime()),
	}
}
