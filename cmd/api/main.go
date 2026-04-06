package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/project/wayt-page/config"
	"github.com/project/wayt-page/internal/handler"
	"github.com/project/wayt-page/internal/repository"
	"github.com/project/wayt-page/internal/service"
	"github.com/project/wayt-page/pkg/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := gorm.Open(postgres.Open(cfg.DB.DSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Repositories
	userRepo := repository.NewUserRepository(db)
	pricingRepo := repository.NewPricingRepository(db)
	testimonialRepo := repository.NewTestimonialRepository(db)
	settingRepo := repository.NewSettingRepository(db)

	// Services
	authSvc := service.NewAuthService(userRepo, cfg.Auth.JWTSecret)
	pricingSvc := service.NewPricingService(pricingRepo)
	testimonialSvc := service.NewTestimonialService(testimonialRepo)
	settingSvc := service.NewSettingService(settingRepo)

	// Seed admin
	if cfg.Auth.AdminPassword != "" {
		if err := authSvc.SeedAdmin(cfg.Auth.AdminUsername, cfg.Auth.AdminPassword); err != nil {
			log.Printf("seed admin skipped: %v", err)
		}
	}

	// Handlers
	authHandler := handler.NewAuthHandler(authSvc)
	pricingHandler := handler.NewPricingHandler(pricingSvc)
	testimonialHandler := handler.NewTestimonialHandler(testimonialSvc)
	settingHandler := handler.NewSettingHandler(settingSvc)

	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.LoadHTMLGlob("web/templates/*")

	// Pages
	r.GET("/", func(c *gin.Context) { c.HTML(200, "index.html", nil) })
	r.GET("/admin", func(c *gin.Context) { c.HTML(200, "admin.html", nil) })

	// Public API
	r.POST("/auth/login", authHandler.Login)
	r.GET("/api/pricing", pricingHandler.ListPublic)
	r.GET("/api/testimonials", testimonialHandler.ListPublic)
	r.GET("/api/settings", settingHandler.GetPublic)

	// Internal (authenticated)
	internal := r.Group("/internal", middleware.JWTAuth(cfg.Auth.JWTSecret))
	{
		internal.GET("/pricing", pricingHandler.List)
		internal.POST("/pricing", pricingHandler.Create)
		internal.PUT("/pricing/:id", pricingHandler.Update)
		internal.DELETE("/pricing/:id", pricingHandler.Delete)

		internal.GET("/testimonials", testimonialHandler.List)
		internal.POST("/testimonials", testimonialHandler.Create)
		internal.PUT("/testimonials/:id", testimonialHandler.Update)
		internal.DELETE("/testimonials/:id", testimonialHandler.Delete)

		internal.GET("/settings", settingHandler.List)
		internal.PUT("/settings", settingHandler.Update)

		internal.GET("/admins", authHandler.ListAdmins)
		internal.POST("/admins", authHandler.CreateAdmin)
		internal.PUT("/admins/:id/password", authHandler.UpdatePassword)
		internal.DELETE("/admins/:id", authHandler.DeleteAdmin)
	}

	addr := fmt.Sprintf(":%s", cfg.AppPort)
	log.Printf("wayt-page server running on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
