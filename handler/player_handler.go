package handler

import (
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pranotoism/football-go/dto"
	"github.com/pranotoism/football-go/service"
	"github.com/pranotoism/football-go/util"
)

type PlayerHandler struct {
	playerService *service.PlayerService
}

func NewPlayerHandler(playerService *service.PlayerService) *PlayerHandler {
	return &PlayerHandler{playerService: playerService}
}

func (h *PlayerHandler) Create(c *gin.Context) {
	teamID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		util.ErrorResponse(c, http.StatusBadRequest, "invalid team ID")
		return
	}

	var req dto.CreatePlayerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	player, err := h.playerService.Create(uint(teamID), req)
	if err != nil {
		if err.Error() == "team not found" {
			util.ErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}
		util.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	util.SuccessResponse(c, http.StatusCreated, "Player created successfully", player)
}

func (h *PlayerHandler) FindByTeam(c *gin.Context) {
	teamID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		util.ErrorResponse(c, http.StatusBadRequest, "invalid team ID")
		return
	}

	page, perPage := getPagination(c)

	players, total, err := h.playerService.FindByTeam(uint(teamID), page, perPage)
	if err != nil {
		if err.Error() == "team not found" {
			util.ErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}
		util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	util.PaginatedSuccessResponse(c, http.StatusOK, "Players retrieved successfully", players, util.Meta{
		Page:       page,
		PerPage:    perPage,
		Total:      total,
		TotalPages: int(math.Ceil(float64(total) / float64(perPage))),
	})
}

func (h *PlayerHandler) FindByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		util.ErrorResponse(c, http.StatusBadRequest, "invalid player ID")
		return
	}

	player, err := h.playerService.FindByID(uint(id))
	if err != nil {
		util.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	util.SuccessResponse(c, http.StatusOK, "Player retrieved successfully", player)
}

func (h *PlayerHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		util.ErrorResponse(c, http.StatusBadRequest, "invalid player ID")
		return
	}

	var req dto.UpdatePlayerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	player, err := h.playerService.Update(uint(id), req)
	if err != nil {
		util.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	util.SuccessResponse(c, http.StatusOK, "Player updated successfully", player)
}

func (h *PlayerHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		util.ErrorResponse(c, http.StatusBadRequest, "invalid player ID")
		return
	}

	if err := h.playerService.Delete(uint(id)); err != nil {
		util.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	util.SuccessResponse(c, http.StatusOK, "Player deleted successfully", nil)
}
