package handlers

import (
	"biophilia/internal/application/services"
	"biophilia/internal/domain/entities"
	"biophilia/internal/infrastructure/mappers"
	"biophilia/internal/presentation/http/rest/entities/requests"
	"biophilia/internal/presentation/http/rest/entities/responses"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
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
//	@Success 201 {object} responses.Biomolecule
//	@Router /biomolecules [post]
func (h *BiomoleculeHandler) AddBiomolecule(c echo.Context) error {
	var biomoleculeRequest requests.AddBiomoleculeRequest
	if err := c.Bind(&biomoleculeRequest); err != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse{Code: http.StatusBadRequest, Message: err.Error()})
	}

	biomolecule := mappers.MapCreateBiomoleculeRequestToDomain(biomoleculeRequest)
	created, err := h.biomoleculeService.AddBiomolecule(biomolecule)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, mappers.MapDomainBiomoleculeToResponse(*created))
}

// GetBiomolecules godoc
//
//	@Summary Get biomolecules
//	@Description Get biomolecules
//	@Tags biomolecules
//	@Accept  json
//	@Produce  json
//	@Success 200 {object} []responses.Biomolecule
//	@Router /biomolecules [get]
func (h *BiomoleculeHandler) GetBiomolecules(c echo.Context) error {
	biomolecules, err := h.biomoleculeService.GetBiomolecules()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	list := make([]entities.Biomolecule, len(biomolecules))
	for i, b := range biomolecules {
		list[i] = *b
	}
	return c.JSON(http.StatusOK, mappers.MapDomainBiomoleculesToResponse(list))
}

// GetBiomoleculeByID godoc
//
//	@Summary Get biomolecule by ID
//	@Description Get biomolecule by ID
//	@Tags biomolecules
//	@Accept  json
//	@Produce  json
//	@Param id path int true "Biomolecule ID"
//	@Success 200 {object} responses.Biomolecule
//	@Router /biomolecules/{id} [get]
func (h *BiomoleculeHandler) GetBiomoleculeByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse{Code: http.StatusBadRequest, Message: "invalid id"})
	}
	biomolecule, err := h.biomoleculeService.GetBiomoleculeByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, mappers.MapDomainBiomoleculeToResponse(*biomolecule))
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
//	@Success 200 {object} responses.Biomolecule
//	@Router /biomolecules/{id} [put]
func (h *BiomoleculeHandler) UpdateBiomolecule(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse{Code: http.StatusBadRequest, Message: "invalid id"})
	}
	var biomoleculeRequest requests.UpdateBiomoleculeRequest
	if err := c.Bind(&biomoleculeRequest); err != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse{Code: http.StatusBadRequest, Message: err.Error()})
	}
	biomolecule := mappers.MapUpdateBiomoleculeRequestToDomain(biomoleculeRequest)
	updated, err := h.biomoleculeService.UpdateBiomolecule(id, biomolecule)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, mappers.MapDomainBiomoleculeToResponse(*updated))
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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse{Code: http.StatusBadRequest, Message: "invalid id"})
	}
	if err := h.biomoleculeService.DeleteBiomolecule(id); err != nil {
		return c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}
