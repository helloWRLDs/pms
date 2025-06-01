package userdata

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
	"pms.pkg/transport/grpc/dto"
	"pms.pkg/utils"
)

type User struct {
	ID        string    `db:"id"`
	FirstName string    `db:"first_name"`
	LastName  *string   `db:"last_name"`
	Email     string    `db:"email"`
	Password  *string   `db:"password"`
	AvatarIMG *[]byte   `db:"avatar_img"`
	AvatarURL *string   `db:"avatar_url"`
	Phone     *string   `db:"phone"`
	Bio       *string   `db:"bio"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (u *User) DTO() *dto.User {
	avatarIMG := make([]byte, 0)
	if u.AvatarIMG != nil {
		avatarIMG = []byte(*u.AvatarIMG)
	}
	return &dto.User{
		Id:        u.ID,
		FirstName: u.FirstName,
		LastName:  utils.Value(u.LastName),
		Email:     u.Email,
		AvatarImg: avatarIMG,
		AvatarUrl: utils.Value(u.AvatarURL),
		Phone:     utils.Value(u.Phone),
		Bio:       utils.Value(u.Bio),
		CreatedAt: timestamppb.New(u.CreatedAt),
		UpdatedAt: timestamppb.New(u.UpdatedAt),
	}
}

func Entity(dto *dto.User) *User {
	avatarIMG := new([]byte)
	if dto.AvatarImg != nil {
		avatarIMG = &dto.AvatarImg
	}
	return &User{
		ID:        dto.Id,
		FirstName: dto.FirstName,
		LastName:  utils.Ptr(dto.LastName),
		Email:     dto.Email,
		AvatarIMG: avatarIMG,
		AvatarURL: utils.Ptr(dto.AvatarUrl),
		Phone:     utils.Ptr(dto.Phone),
		Bio:       utils.Ptr(dto.Bio),
		CreatedAt: dto.CreatedAt.AsTime(),
		UpdatedAt: dto.UpdatedAt.AsTime(),
	}
}
