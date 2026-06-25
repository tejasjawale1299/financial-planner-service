package main

import (
	dao "FinancialPlannerGo/app/domain/dao/financialPlanner"
	"FinancialPlannerGo/config"
	"log"
)

func main() {
	config.LoadEnv()
	config.ConnectToDB()

	if err := config.DB.AutoMigrate(&dao.FinancialPlannerReport{}); err != nil {
		log.Fatal("migration failed:", err)
	}

	log.Println("migration completed successfully")
}
