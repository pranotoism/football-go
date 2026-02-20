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

type MatchHandler struct {
	matchService *service.MatchService
}

func NewMatchHandler(matchService *service.MatchService) *MatchHandler {
	return &MatchHandler{matchService: matchService}
}

func (h *MatchHandler) Create(c *gin.Context) {
	var req dto.CreateMatchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	match, err := h.matchService.Create(req)
	if err != nil {
		util.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	util.SuccessResponse(c, http.StatusCreated, "Match created successfully", match)
}

func (h *MatchHandler) FindAll(c *gin.Context) {
	page, perPage := getPagination(c)

	matches, total, err := h.matchService.FindAll(page, perPage)
	if err != nil {
		util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	util.PaginatedSuccessResponse(c, http.StatusOK, "Matches retrieved successfully", matches, util.Meta{
		Page:       page,
		PerPage:    perPage,
		Total:      total,
		TotalPages: int(math.Ceil(float64(total) / float64(perPage))),
	})
}

func (h *MatchHandler) FindByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		util.ErrorResponse(c, http.StatusBadRequest, "invalid match ID")
		return
	}

	match, err := h.matchService.FindByID(uint(id))
	if err != nil {
		util.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	util.SuccessResponse(c, http.StatusOK, "Match retrieved successfully", match)
}

func (h *MatchHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		util.ErrorResponse(c, http.StatusBadRequest, "invalid match ID")
		return
	}

	var req dto.UpdateMatchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	match, err := h.matchService.Update(uint(id), req)
	if err != nil {
		util.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	util.SuccessResponse(c, http.StatusOK, "Match updated successfully", match)
}

func (h *MatchHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		util.ErrorResponse(c, http.StatusBadRequest, "invalid match ID")
		return
	}

	if err := h.matchService.Delete(uint(id)); err != nil {
		util.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	util.SuccessResponse(c, http.StatusOK, "Match deleted successfully", nil)
}

func (h *MatchHandler) ReportResult(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		util.ErrorResponse(c, http.StatusBadRequest, "invalid match ID")
		return
	}

	var req dto.ReportResultRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	match, err := h.matchService.ReportResult(uint(id), req)
	if err != nil {
		util.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	util.SuccessResponse(c, http.StatusOK, "Match result reported successfully", match)
}
