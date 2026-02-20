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

type TeamHandler struct {
	teamService *service.TeamService
}

func NewTeamHandler(teamService *service.TeamService) *TeamHandler {
	return &TeamHandler{teamService: teamService}
}

func (h *TeamHandler) Create(c *gin.Context) {
	var req dto.CreateTeamRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	team, err := h.teamService.Create(req)
	if err != nil {
		util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	util.SuccessResponse(c, http.StatusCreated, "Team created successfully", team)
}

func (h *TeamHandler) FindAll(c *gin.Context) {
	page, perPage := getPagination(c)

	teams, total, err := h.teamService.FindAll(page, perPage)
	if err != nil {
		util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	util.PaginatedSuccessResponse(c, http.StatusOK, "Teams retrieved successfully", teams, util.Meta{
		Page:       page,
		PerPage:    perPage,
		Total:      total,
		TotalPages: int(math.Ceil(float64(total) / float64(perPage))),
	})
}

func (h *TeamHandler) FindByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		util.ErrorResponse(c, http.StatusBadRequest, "invalid team ID")
		return
	}

	team, err := h.teamService.FindByID(uint(id))
	if err != nil {
		util.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	util.SuccessResponse(c, http.StatusOK, "Team retrieved successfully", team)
}

func (h *TeamHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		util.ErrorResponse(c, http.StatusBadRequest, "invalid team ID")
		return
	}

	var req dto.UpdateTeamRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	team, err := h.teamService.Update(uint(id), req)
	if err != nil {
		util.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	util.SuccessResponse(c, http.StatusOK, "Team updated successfully", team)
}

func (h *TeamHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		util.ErrorResponse(c, http.StatusBadRequest, "invalid team ID")
		return
	}

	if err := h.teamService.Delete(uint(id)); err != nil {
		util.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	util.SuccessResponse(c, http.StatusOK, "Team deleted successfully", nil)
}

func getPagination(c *gin.Context) (int, int) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))

	if page < 1 {
		page = 1
	}
	if perPage < 1 || perPage > 100 {
		perPage = 10
	}

	return page, perPage
}
