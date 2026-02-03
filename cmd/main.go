package main

import (
	"fmt"
	"log"

	"github.com/aashisDevv/login-api/config"
	"github.com/aashisDevv/login-api/controller"
	"github.com/aashisDevv/login-api/db"
	"github.com/aashisDevv/login-api/models"
	"github.com/aashisDevv/login-api/repository"
	"github.com/aashisDevv/login-api/router"
	"github.com/aashisDevv/login-api/security"
	"github.com/aashisDevv/login-api/service"
	"github.com/go-playground/validator/v10"
)

func main() {
	// Load Config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error in loading config: %v", err)
	}
	
	// Connect to DB
	db, err := db.ConnectDB(cfg.DBUrl)
	if err != nil {
		log.Fatalf("error in connecting to DB: %v", err)
	}
	
	// do migration in DB
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("error in migration in DB: %v", err)
	}
	
	// Inject dependencies
	userRepo := repository.NewUserRepository(db)
	tokenService := security.NewTokenService(&cfg.AuthConfig)
	
	userService := service.NewUserService(userRepo, tokenService)
	validator := validator.New()
	userHandler := controller.NewUserController(userService, validator)
	
	// Setup Routes
	r := router.SetupRouter(userHandler)
	
	// Run Server
	log.Printf("server listening on Port : %v", cfg.Port)
	if err := r.Run(fmt.Sprintf(":%v", cfg.Port)); err != nil {
		log.Fatalf("error in runnin server: %v", err)
	}
}