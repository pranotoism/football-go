package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pranotoism/football-go/handler"
	"github.com/pranotoism/football-go/middleware"
)

func Setup(
	authHandler *handler.AuthHandler,
	teamHandler *handler.TeamHandler,
	playerHandler *handler.PlayerHandler,
	matchHandler *handler.MatchHandler,
	reportHandler *handler.ReportHandler,
) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.ErrorHandler())

	v1 := r.Group("/api/v1")
	{
		// Public routes
		auth := v1.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}

		// Protected routes
		protected := v1.Group("")
		protected.Use(middleware.AuthMiddleware())
		{
			// Teams
			teams := protected.Group("/teams")
			{
				teams.POST("", teamHandler.Create)
				teams.GET("", teamHandler.FindAll)
				teams.GET("/:id", teamHandler.FindByID)
				teams.PUT("/:id", teamHandler.Update)
				teams.DELETE("/:id", teamHandler.Delete)
				teams.POST("/:id/players", playerHandler.Create)
				teams.GET("/:id/players", playerHandler.FindByTeam)
			}

			// Players (flat routes for individual operations)
			players := protected.Group("/players")
			{
				players.GET("/:id", playerHandler.FindByID)
				players.PUT("/:id", playerHandler.Update)
				players.DELETE("/:id", playerHandler.Delete)
			}

			// Matches
			matches := protected.Group("/matches")
			{
				matches.POST("", matchHandler.Create)
				matches.GET("", matchHandler.FindAll)
				matches.GET("/:id", matchHandler.FindByID)
				matches.PUT("/:id", matchHandler.Update)
				matches.DELETE("/:id", matchHandler.Delete)
				matches.POST("/:id/result", matchHandler.ReportResult)
				matches.GET("/:id/report", reportHandler.GetMatchReport)
			}

			// Reports
			reports := protected.Group("/reports")
			{
				reports.GET("/matches", reportHandler.GetAllMatchReports)
			}
		}
	}

	return r
}
