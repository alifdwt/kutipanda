package app

import (
	"github.com/alifdwt/kutipanda-backend/internal/handler"
	"github.com/alifdwt/kutipanda-backend/internal/mapper"
	"github.com/alifdwt/kutipanda-backend/internal/repository"
	"github.com/alifdwt/kutipanda-backend/internal/service"
	"github.com/alifdwt/kutipanda-backend/pkg/auth"
	"github.com/alifdwt/kutipanda-backend/pkg/cloudinary"
	"github.com/alifdwt/kutipanda-backend/pkg/database/postgres"
	"github.com/alifdwt/kutipanda-backend/pkg/dotenv"
	"github.com/alifdwt/kutipanda-backend/pkg/hashing"
	"github.com/alifdwt/kutipanda-backend/pkg/logger"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// @title Kutipanda API
// @version 1.0
// @description REST API for Kutipanda App

// @contact.name Alif Dewantara
// @contact.url https://github.com/alifdwt
// @contact.email aputradewantara@gmail.com

// @host localhost:8080
// @BasePath /api

// @securityDefinitions.apikey BearerAuth
// @in Header
// @name Authorization
func Run() {
	log, err := logger.NewLogger()
	if err != nil {
		log.Error("Error while initializing logger", zap.Error(err))
	}

	err = dotenv.Viper()
	if err != nil {
		log.Error("Error while loading config", zap.Error(err))
	}

	db, err := postgres.NewClient()
	if err != nil {
		log.Error("Error while connecting to database", zap.Error(err))
	}

	myCloudinary, err := cloudinary.NewMyCloudinary()
	if err != nil {
		log.Error("Error while connecting to cloudinary", zap.Error(err))
	}

	hashing := hashing.NewHashingPassword()
	repository := repository.NewRepositories(db)

	token, err := auth.NewManager(viper.GetString("JWT_SECRET"))
	if err != nil {
		log.Error("Error while intializing token manager", zap.Error(err))
	}

	mapper := mapper.NewMapper()

	service := service.NewService(service.Deps{
		Repository: repository,
		Hashing:    *hashing,
		Token:      token,
		Logger:     *log,
		Mapper:     *mapper,
	})

	myHandler := handler.NewHandler(service, *myCloudinary, token)

	myHandler.InitHandler().Listen(":8080")
}
