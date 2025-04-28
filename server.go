package main

import (
	"edu-learn/config"
	"edu-learn/controller"
	"edu-learn/middleware"
	"edu-learn/repository"
	"edu-learn/usecase"
	"edu-learn/utils/service"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	// Instean usecase
	materialUC usecase.MaterialUseCase
	courseUC   usecase.CourseUseCase
	userUC     usecase.UserUseCase
	authUC     usecase.AuthenticationUseCase
	jwtService service.JwtService
	engine     *gin.Engine
	host       string
}

func NewServer() *Server {
	cfg, _ := config.NewConfig()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta", cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Database)

	// Setting menggunakan gorm
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	jwtService := service.NewJwtService(cfg.TokenConfig)

	// Instean repository
	userRepo := repository.NewUserRepository(db)
	courseRepo := repository.NewCourseRepository(db)
	materialRepo := repository.NewMaterialRepository(db)

	// Instean usecase
	userUseCase := usecase.NewUserUseCase(userRepo)
	courseUseCase := usecase.NewCourseUsecase(courseRepo, userRepo)
	materialUseCase := usecase.NewMaterialUseCase(materialRepo, courseRepo)

	// Auth usecase
	authUseCase := usecase.NewAuthenticationUsecase(userUseCase, jwtService)

	engine := gin.Default()

	host := fmt.Sprintf(":%s", cfg.ApiPort)

	return &Server{
		// Instean usecase
		materialUC: materialUseCase,
		courseUC:   courseUseCase,
		userUC:     userUseCase,
		authUC:     authUseCase,
		jwtService: jwtService,
		engine:     engine,
		host:       host,
	}
}

func (s *Server) initRoute() {
	rg := s.engine.Group("/api")

	authMiddleware := middleware.NewAuthMiddleware(s.jwtService)

	// Instean controller
	controller.NewAuthController(s.authUC, rg).Route()
	controller.NewUserController(s.userUC, rg, authMiddleware).Route()
	controller.NewCourseController(s.courseUC, rg, authMiddleware).Route()
	controller.NewMaterialController(s.materialUC, rg, authMiddleware).Route()
}

func (s *Server) Run() {
	s.initRoute()
	s.engine.SetTrustedProxies([]string{"127.0.0.1"})

	err := s.engine.Run(s.host)
	if err != nil {
		panic(fmt.Errorf("failed to start server on host %s: %v", s.host, err))
	}
}
