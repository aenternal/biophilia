package rest

import (
	"biophilia/internal/domain/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

type BiomoleculeHandler struct {
	biomoleculeService *services.BiomoleculeService
}

func NewBiomoleculeHandler(service *services.BiomoleculeService) *BiomoleculeHandler {
	return &BiomoleculeHandler{
		biomoleculeService: service,
	}
}

func (h *BiomoleculeHandler) AddProtein(c echo.Context) error {
	return c.String(http.StatusOK, "Protein added")
}

func (h *BiomoleculeHandler) GetProteins(c echo.Context) error {
	return c.String(http.StatusOK, "List of proteins")
}

func (h *BiomoleculeHandler) GetProteinByID(c echo.Context) error {
	return c.String(http.StatusOK, "Protein by ID")
}

func (h *BiomoleculeHandler) UpdateProtein(c echo.Context) error {
	return c.String(http.StatusOK, "Protein updated")
}

func (h *BiomoleculeHandler) DeleteProtein(c echo.Context) error {
	return c.String(http.StatusOK, "Protein deleted")
}
