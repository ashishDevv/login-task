package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/aashisDevv/login-api/dto"
	"github.com/aashisDevv/login-api/service"
	"github.com/aashisDevv/login-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserService interface {
	Login(context.Context, dto.LoginRequest) (service.LoginResult, *utils.AppError)
}
type UserController struct {
	userService UserService
	validator   *validator.Validate
}

func NewUserController(userService UserService, validator *validator.Validate) *UserController {
	return &UserController{
		userService: userService,
		validator: validator,
	}
}

func (h *UserController) Login(c *gin.Context) {
	r := c.Request
	w := c.Writer

	ctx := r.Context()

	var req dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if err := h.validator.Struct(req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	result, err := h.userService.Login(ctx, req)
	if err != nil {
		http.Error(w, err.Message, err.Code)
		return
	}
	
	resp := dto.LoginResponse{
		UserId: result.ID,
		Token: result.Token,
		Message: "User Logged in Successfully",
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("error in sending response: %v", err)
	}
}
