package main

import (
	"log"

	"github.com/pranotoism/football-go/config"
	"github.com/pranotoism/football-go/database"
	"github.com/pranotoism/football-go/handler"
	"github.com/pranotoism/football-go/model"
	"github.com/pranotoism/football-go/repository"
	"github.com/pranotoism/football-go/router"
	"github.com/pranotoism/football-go/service"
	"github.com/pranotoism/football-go/util"
)

func main() {
	// Load config
	cfg := config.Load()

	// Set JWT secret
	util.SetJWTSecret(cfg.JWTSecret)

	// Connect database
	db := database.Connect(cfg)

	// Auto migrate
	database.Migrate(db,
		&model.User{},
		&model.Team{},
		&model.Player{},
		&model.Match{},
		&model.Goal{},
	)

	// Repositories
	userRepo := repository.NewUserRepository(db)
	teamRepo := repository.NewTeamRepository(db)
	playerRepo := repository.NewPlayerRepository(db)
	matchRepo := repository.NewMatchRepository(db)
	goalRepo := repository.NewGoalRepository(db)

	// Services
	authService := service.NewAuthService(userRepo)
	teamService := service.NewTeamService(teamRepo, playerRepo)
	playerService := service.NewPlayerService(playerRepo, teamRepo)
	matchService := service.NewMatchService(matchRepo, teamRepo, goalRepo)
	reportService := service.NewReportService(matchRepo)

	// Handlers
	authHandler := handler.NewAuthHandler(authService)
	teamHandler := handler.NewTeamHandler(teamService)
	playerHandler := handler.NewPlayerHandler(playerService)
	matchHandler := handler.NewMatchHandler(matchService)
	reportHandler := handler.NewReportHandler(reportService)

	// Setup router
	r := router.Setup(authHandler, teamHandler, playerHandler, matchHandler, reportHandler)

	// Start server
	log.Printf("Server starting on port %s", cfg.AppPort)
	if err := r.Run(":" + cfg.AppPort); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
