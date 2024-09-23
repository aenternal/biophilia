package rest

import (
	"biophilia/internal/domain/services"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, proteinUsecase *services.BiomoleculeService) {
	handler := NewProteinHandler(proteinUsecase)

	// Инициализация маршрутов для работы с белками
	e.POST("/proteins", handler.AddProtein)
	e.GET("/proteins", handler.GetProteins)
	e.GET("/proteins/:id", handler.GetProteinByID)
	e.PUT("/proteins/:id", handler.UpdateProtein)
	e.DELETE("/proteins/:id", handler.DeleteProtein)
}
