package handlers

import (
	"biophilia/internal/domain/entities"
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

// AddBiomolecule godoc
//
//	@Summary Create a new biomolecule
//	@Description Create a new biomolecule in the database
//	@Tags biomolecules
//	@Accept  json
//	@Produce  json
//	@Param biomolecule body entities.AddBiomoleculeRequest true "Biomolecule to create"
//	@Success 201 {object} entities.Biomolecule
//	@Router /biomolecules [post]
func (h *BiomoleculeHandler) AddBiomolecule(c echo.Context) error {
	var biomoleculeRequest entities.AddBiomoleculeRequest
	if err := c.Bind(biomoleculeRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.String(http.StatusOK, "Biomolecule added")
}

// GetBiomolecules godoc
//
//	@Summary Get biomolecules
//	@Description Get biomolecules
//	@Tags biomolecules
//	@Accept  json
//	@Produce  json
//	@Success 200 {object} []entities.Biomolecule
//	@Router /biomolecules [get]
func (h *BiomoleculeHandler) GetBiomolecules(c echo.Context) error {
	return c.String(http.StatusOK, "List of biomolecules")
}

// GetBiomoleculeByID godoc
//
//	@Summary Get biomolecule by ID
//	@Description Get biomolecule by ID
//	@Tags biomolecules
//	@Accept  json
//	@Produce  json
//	@Param id path int true "Biomolecule ID"
//	@Success 200 {object} entities.Biomolecule
//	@Router /biomolecules/{id} [get]
func (h *BiomoleculeHandler) GetBiomoleculeByID(c echo.Context) error {
	return c.String(http.StatusOK, "Biomolecule by ID")
}

// UpdateBiomolecule godoc
//
//	@Summary Update biomolecule
//	@Description Update biomolecule
//	@Tags biomolecules
//	@Accept  json
//	@Produce  json
//	@Param id path int true "Biomolecule ID"
//	@Param biomolecule body entities.UpdateBiomoleculeRequest true "Biomolecule to create"
//	@Success 200 {object} entities.Biomolecule
//	@Router /biomolecules/{id} [put]
func (h *BiomoleculeHandler) UpdateBiomolecule(c echo.Context) error {
	return c.String(http.StatusOK, "Biomolecule updated")
}

// DeleteBiomolecule
//
//	@Summary Delete biomolecule
//	@Description Delete biomolecule
//	@Tags biomolecules
//	@Accept  json
//	@Produce  json
//	@Param id path int true "Biomolecule ID"
//	@Success 204
//	@Router /biomolecules/{id} [delete]
func (h *BiomoleculeHandler) DeleteBiomolecule(c echo.Context) error {
	return c.String(http.StatusOK, "Biomolecule deleted")
}
