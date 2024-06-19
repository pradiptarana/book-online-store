package users

import (
	"fmt"

	"github.com/pradiptarana/book-online-store/internal/auth"
	"github.com/pradiptarana/book-online-store/model"
	"github.com/pradiptarana/book-online-store/repository"
	"golang.org/x/crypto/bcrypt"
)

type UsersUC struct {
	repository.UserRepository
}

func NewUserUC(repo repository.UserRepository) *UsersUC {
	return &UsersUC{repo}
}

func (uc *UsersUC) SignUp(req *model.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	fmt.Println(string(hashedPassword))
	if err != nil {
		return err
	}
	req.Password = string(hashedPassword)
	err = uc.UserRepository.SignUp(req)
	if err != nil {
		return err
	}
	return nil
}

func (uc *UsersUC) Login(req *model.LoginRequest) (string, error) {
	user, err := uc.UserRepository.GetUser(req.Username)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", err
	}

	token, err := auth.GenerateJWT(user.Id)
	if err != nil {
		return "", err
	}

	return token, nil
}
