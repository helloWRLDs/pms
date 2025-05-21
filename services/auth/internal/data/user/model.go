package userdata

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
	"pms.pkg/transport/grpc/dto"
	"pms.pkg/utils"
)

type User struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	AvatarIMG []byte    `db:"avatar_img"`
	Phone     *string   `db:"phone"`
	Bio       *string   `db:"bio"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (u *User) DTO() *dto.User {
	return &dto.User{
		Id:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		AvatarImg: u.AvatarIMG,
		Phone:     utils.Value(u.Phone),
		Bio:       utils.Value(u.Bio),
		CreatedAt: timestamppb.New(u.CreatedAt),
		UpdatedAt: timestamppb.New(u.UpdatedAt),
	}
}

func Entity(dto *dto.User) *User {
	return &User{
		ID:        dto.Id,
		Name:      dto.Name,
		Email:     dto.Email,
		AvatarIMG: dto.AvatarImg,
		Phone:     utils.Ptr(dto.Phone),
		Bio:       utils.Ptr(dto.Bio),
		CreatedAt: dto.CreatedAt.AsTime(),
		UpdatedAt: dto.UpdatedAt.AsTime(),
	}
}
