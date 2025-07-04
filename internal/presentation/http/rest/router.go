package rest

import (
	"biophilia/internal/application/services"
	"biophilia/internal/presentation/http/rest/handlers"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, biomoleculeService *services.BiomoleculeService) {
	handler := handlers.NewBiomoleculeHandler(biomoleculeService)
	api := e.Group("/api")
	v1 := api.Group("/v1")

	v1.POST("/biomolecules", handler.AddBiomolecule)
	v1.GET("/biomolecules", handler.GetBiomolecules)
	v1.GET("/biomolecules/:id", handler.GetBiomoleculeByID)
	v1.PUT("/biomolecules/:id", handler.UpdateBiomolecule)
	v1.DELETE("/biomolecules/:id", handler.DeleteBiomolecule)
}
