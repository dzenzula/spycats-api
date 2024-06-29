package routes

import (
	"net/http"

	"cmd/catapi/internal/handlers"
	"cmd/catapi/internal/middleware"

	"cmd/catapi/docs"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() (*gin.Engine, error) {
	r := gin.New()
	docs.SwaggerInfo.Title = "Spy Cat API"
	docs.SwaggerInfo.Description = "This is a sample server for a spy cat management system."
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r.Use(gin.Recovery())
	r.Use(middleware.Logger())

	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{
		Path:     "/spycats-api/api",
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
	})
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	spycatsGroup := r.Group("/api/SpyCats")
	{
		spycatsGroup.POST("/CreateSpyCat", handlers.CreateSpyCat)
		spycatsGroup.DELETE("/RemoveSpyCat", handlers.RemoveSpyCat)
		spycatsGroup.PUT("/UpdateSpyCat", handlers.UpdateSpyCat)
		spycatsGroup.GET("/GetSpyCats", handlers.GetSpyCats)
		spycatsGroup.GET("/GetSpyCat", handlers.GetSpyCat)
	}

	missionsGroup := r.Group("/api/Missions")
	{
		missionsGroup.POST("/CreateMission", handlers.CreateMission)
		missionsGroup.DELETE("/RemoveMission", handlers.RemoveMission)
		missionsGroup.PUT("/UpdateMission", handlers.UpdateMission)
		missionsGroup.PUT("/UpdateMissionTargets", handlers.UpdateMissionTargets)
		missionsGroup.PUT("/AssignCat", handlers.AssignCat)
		missionsGroup.GET("/GetMissions", handlers.GetMissions)
		missionsGroup.GET("/GetMission", handlers.GetMission)
	}

	targetsGroup := r.Group("api/Targets")
	{
		targetsGroup.POST("/AddTarget", handlers.AddTarget)
		targetsGroup.DELETE("/RemoveTarget", handlers.RemoveTarget)
	}

	if err := r.Run(); err != nil {
		return nil, err
	}

	return r, nil
}
