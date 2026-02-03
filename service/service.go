package service

import (
	"context"
	"log"
	"net/http"

	"github.com/aashisDevv/login-api/dto"
	"github.com/aashisDevv/login-api/models"
	"github.com/aashisDevv/login-api/security"
	"github.com/aashisDevv/login-api/utils"
)

type UserRepository interface {
	GetUserByEmail(context.Context, string) (*models.User, error)
}

type userService struct {
	userRepo UserRepository
	tokenService *security.TokenService
}

func NewUserService(userRepo UserRepository, tokenService *security.TokenService) *userService {
	return &userService{
		userRepo: userRepo,
		tokenService: tokenService,
	}
}

func (s *userService) Login(ctx context.Context, req dto.LoginRequest) (LoginResult, *utils.AppError) {
	user, err := s.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return LoginResult{}, utils.New(http.StatusUnauthorized, "Invalid credentials")
	}
	
	log.Printf("we are here: %v", user)
	
	ok, err := security.ComparePassword(req.Password, user.PasswordHash)
	if err != nil || !ok {
		return LoginResult{}, utils.New(http.StatusUnauthorized, "Invalid credentials")
	}
	
	payload := security.RequestClaims{
		UserID: user.ID,
	}
	
	token, err := s.tokenService.GenerateAccessToken(payload)
	if err != nil {
		return LoginResult{}, utils.New(http.StatusInternalServerError, "Internal server error")
	}
	
	return LoginResult{
		ID: user.ID,
		Token: token,
	}, nil	
}
