package FinancialPlanner

import (
	service "FinancialPlannerGo/app/services/FinancialPlanner"
	"FinancialPlannerGo/app/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FinancialPlannerController struct {
	service service.FinancialPlannerService
}

func NewFinancialPlannerController(financialPlannerService service.FinancialPlannerService) *FinancialPlannerController {
	return &FinancialPlannerController{
		service: financialPlannerService,
	}
}

func (ctrl *FinancialPlannerController) GenerateReport(c *gin.Context) {

	var req utils.FinancialPlannerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("VAL_001", "invalid request payload"))
		return
	}
	utils.NormalizeFinancialPlannerRequest(&req)
	if err := utils.ValidateFinancialPlannerRequest(req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("VAL_001", err.Error()))
		return
	}
	userID := getUserID(c)
	_, apiResp, err := ctrl.service.GenerateFinancialReport(c.Request.Context(), userID, req)
	if err != nil {
		c.JSON(http.StatusBadGateway, utils.ErrorResponse("FP_001", err.Error()))
		return
	}
	// Return Policy Planner response directly
	c.JSON(http.StatusOK, apiResp.Raw)

}

func (ctrl *FinancialPlannerController) GetReport(c *gin.Context) {
	id, err := parseID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("VAL_001", "invalid report id"))
		return
	}

	report, err := ctrl.service.GetReportByID(id, getUserID(c))
	if err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse("FP_404", "report not found"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Report fetched successfully", utils.ToFinancialPlannerReportResponse(report)))
}

func (ctrl *FinancialPlannerController) DownloadReport(c *gin.Context) {
	id, err := parseID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("VAL_001", "invalid report id"))
		return
	}

	report, err := ctrl.service.GetReportByID(id, getUserID(c))
	if err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse("FP_404", "report not found"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("PDF URL fetched successfully", gin.H{
		"report_pdf_url":  report.ReportPDFURL,
		"report_pdf_path": report.ReportPDFPath,
	}))
}

func parseID(value string) (uint, error) {
	id, err := strconv.ParseUint(value, 10, 64)
	if err != nil || id == 0 {
		return 0, err
	}
	return uint(id), nil
}

func getUserID(c *gin.Context) uint {
	value := c.GetHeader("X-User-ID")
	if value == "" {
		return 0
	}

	id, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return 0
	}

	return uint(id)
}
