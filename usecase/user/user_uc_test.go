package users_test

import (
	"errors"
	"testing"

	"github.com/pradiptarana/book-online-store/mocks"
	"github.com/pradiptarana/book-online-store/model"

	"github.com/golang/mock/gomock"

	usersUC "github.com/pradiptarana/book-online-store/usecase/user"
)

func TestLoginFailed(t *testing.T) {
	req := &model.LoginRequest{
		Username: "user",
		Password: "pass",
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserRepo := mocks.NewMockUserRepository(mockCtrl)
	userUC := usersUC.NewUserUC(mockUserRepo)

	mockUserRepo.EXPECT().GetUser(req.Username).Return(nil, errors.New("user not found")).Times(1)

	_, err := userUC.Login(req)
	if err == nil {
		t.Fail()
	}
}

func TestLoginSuccess(t *testing.T) {
	req := &model.LoginRequest{
		Username: "user",
		Password: "pass",
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserRepo := mocks.NewMockUserRepository(mockCtrl)
	userUC := usersUC.NewUserUC(mockUserRepo)

	mockUserRepo.EXPECT().GetUser(req.Username).Return(&model.User{
		Id:       1,
		Username: req.Username,
		Password: "$2a$10$.lcutDtRp6R0r.pOJYUrpuoySCKXpjfclEkc7XTQXdCAjb5LUG.bm",
		Email:    "abcd@test.com",
	}, nil).Times(1)

	_, err := userUC.Login(req)
	if err != nil {
		t.Fail()
	}
}

func TestSignUpSuccess(t *testing.T) {
	req := &model.User{
		Username: "user",
		Password: "pass",
		Email:    "ranapradipta4@gmail.com",
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserRepo := mocks.NewMockUserRepository(mockCtrl)
	userUC := usersUC.NewUserUC(mockUserRepo)

	mockUserRepo.EXPECT().SignUp(req).Return(nil).Times(1)

	err := userUC.SignUp(req)
	if err != nil {
		t.Fail()
	}
}

func TestSignUpFailed(t *testing.T) {
	req := &model.User{
		Username: "user",
		Password: "pass",
		Email:    "ranapradipta4@gmail.com",
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserRepo := mocks.NewMockUserRepository(mockCtrl)
	userUC := usersUC.NewUserUC(mockUserRepo)

	mockUserRepo.EXPECT().SignUp(req).Return(errors.New("user already exist")).Times(1)

	err := userUC.SignUp(req)
	if err == nil {
		t.Fail()
	}
}
