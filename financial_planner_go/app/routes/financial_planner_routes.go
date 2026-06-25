package routes

import (
	controller "FinancialPlannerGo/app/controllers/FinancialPlanner"
	repo "FinancialPlannerGo/app/repositories/FinancialPlanner"
	service "FinancialPlannerGo/app/services/FinancialPlanner"
	"FinancialPlannerGo/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterFinancialPlannerRoutes(router *gin.Engine) {
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "Financial Planner service is running"})
	})

	financialPlannerRepo := repo.NewFinancialPlannerRepository(config.DB)
	financialPlannerService := service.NewFinancialPlannerService(financialPlannerRepo)
	financialPlannerController := controller.NewFinancialPlannerController(financialPlannerService)

	finance := router.Group("/finance")
	{
		finance.POST("/financial-planner/report", financialPlannerController.GenerateReport)
		finance.GET("/financial-planner/report/:id", financialPlannerController.GetReport)
		finance.GET("/financial-planner/report/download/:id", financialPlannerController.DownloadReport)
	}
}
