package main

import (
	"FinancialPlannerGo/app/middleware"
	"FinancialPlannerGo/app/routes"
	"FinancialPlannerGo/config"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	config.ConnectToDB()

	if config.GetEnv("ENV", "local") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery(), middleware.SecurityHeaders(), middleware.LimitRequestBody())

	routes.RegisterFinancialPlannerRoutes(router)

	port := config.GetEnv("APP_PORT", "8080")
	log.Println("Financial Planner service running on port:", port)

	server := &http.Server{
		Addr:              ":" + port,
		Handler:           router,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       20 * time.Second,
		WriteTimeout:      20 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("server start failed:", err)
	}
}
