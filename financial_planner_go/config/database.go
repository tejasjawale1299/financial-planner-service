package config

import (
	dao "FinancialPlannerGo/app/domain/dao/financialPlanner"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	host := GetEnv("DB_HOST", "localhost")
	port := GetEnv("DB_PORT", "5433")
	user := GetEnv("DB_USER", "postgres")
	password := GetEnv("DB_PASSWORD", "postgres")
	dbname := GetEnv("DB_NAME", "financial_planner")
	sslmode := GetEnv("DB_SSLMODE", "disable")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("database connection failed: ", err)
	}

	DB = db

	if err := DB.AutoMigrate(&dao.FinancialPlannerReport{}); err != nil {
		log.Fatal("migration failed: ", err)
	}

	log.Println("database connected successfully")
}
