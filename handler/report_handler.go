package handler

import (
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pranotoism/football-go/service"
	"github.com/pranotoism/football-go/util"
)

type ReportHandler struct {
	reportService *service.ReportService
}

func NewReportHandler(reportService *service.ReportService) *ReportHandler {
	return &ReportHandler{reportService: reportService}
}

func (h *ReportHandler) GetMatchReport(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		util.ErrorResponse(c, http.StatusBadRequest, "invalid match ID")
		return
	}

	report, err := h.reportService.GetMatchReport(uint(id))
	if err != nil {
		util.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	util.SuccessResponse(c, http.StatusOK, "Match report retrieved successfully", report)
}

func (h *ReportHandler) GetAllMatchReports(c *gin.Context) {
	page, perPage := getPagination(c)

	reports, total, err := h.reportService.GetAllMatchReports(page, perPage)
	if err != nil {
		util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	util.PaginatedSuccessResponse(c, http.StatusOK, "Match reports retrieved successfully", reports, util.Meta{
		Page:       page,
		PerPage:    perPage,
		Total:      total,
		TotalPages: int(math.Ceil(float64(total) / float64(perPage))),
	})
}
