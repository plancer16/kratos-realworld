package biz

import (
	"context"
	"fmt"

	v1 "realworld_02/api/realworld/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

type User struct {
	Username     string
	Email        string
	Bio          string
	Image        string
	Passwordhash string
}

type UserLogin struct {
	Username string
	Email    string
	Bio      string
	Image    string
	Token    string
}

func hashPassword(pwd string) string {
	b, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", b)
	return string(b)
}

func verifyPassword(hashed, input string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(input)); err != nil {
		return false
	}
	return true
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) Register(ctx context.Context, email, username, password string) (*UserLogin, error) {
	u := &User{
		Email:        email,
		Username:     username,
		Passwordhash: hashPassword(password),
	}
	err := uc.repo.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}
	return &UserLogin{
		Username: username,
		Email:    email,
		Token:    "token:" + password,
	}, nil
}

func (uc *UserUsecase) Login(ctx context.Context, email, password string) (*UserLogin, error) {
	u, err := uc.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if !verifyPassword(u.Passwordhash, password) {
		panic("wrong password,login failed")
	}
	return &UserLogin{
		Username: u.Username,
		Email:    u.Email,
		Token:    "token" + password,
	}, nil
}
