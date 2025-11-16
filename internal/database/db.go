package database

import (
	"fmt"
	"time"

	"exptracker/internal/config"
	"exptracker/internal/logger"

	models "exptracker/internal/domain/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect(cfg config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=Asia/Kolkata",
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.Port,
		cfg.Database.SSLMode,
	)

	gormLogger := glogger.Default.LogMode(glogger.Warn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})

	if err != nil {
		logger.L().Fatal().Err(err).Msg("Failed to connect to PostgreSQL")
	}

	sqlDB, err := db.DB()
	if err != nil {
		logger.L().Fatal().Err(err).Msg("Failed to get DB handle")
	}

	// Connection pool tuning
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(1 * time.Hour)

	logger.L().Info().Msg("Connected to PostgreSQL")

	err = db.AutoMigrate(
		&models.User{},
		&models.Account{},
		&models.Expense{},
		&models.RecurringExpense{},
		&models.AuditLog{},
	)
	if err != nil {
		logger.L().Fatal().Err(err).Msg("Failed DB migrations")
	}

	DB = db
	return db
}

func Close() {
	if DB == nil {
		return
	}

	sqlDB, err := DB.DB()
	if err == nil {
		sqlDB.Close()
		logger.L().Info().Msg("PostgreSQL connection closed")
	}
}
