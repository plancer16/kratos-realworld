package data

import (
	"context"

	"realworld_02/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Username     string `gorm:"size:500"`
	Bio          string `gorm:"size:1000"`
	Email        string `grom:"size:1000"`
	Image        string `grom:"szie:1000"`
	Passwordhash string `grom:"size:500"`
}

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) CreateUser(ctx context.Context, user *biz.User) error {
	u := &User{
		Username:     user.Username,
		Bio:          user.Bio,
		Email:        user.Email,
		Image:        user.Image,
		Passwordhash: user.Passwordhash,
	}
	err := r.data.db.Create(u).Error
	return err
}

func (r *userRepo) GetUserByEmail(ctx context.Context, email string) (*biz.User, error) {
	u := &User{}
	err := r.data.db.Where("email = ?", email).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &biz.User{
		Username:     u.Username,
		Email:        u.Email,
		Bio:          u.Bio,
		Passwordhash: u.Passwordhash,
		Image:        u.Image,
	}, nil
}