package sqlite

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Muchogoc/phone-numbers-exercise/service/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	RunningTestsEnvName = "IS_RUNNING_TEST"
)

type DBImpl struct {
	db *gorm.DB
}

func IsRunningTests() bool {
	envValue := os.Getenv(RunningTestsEnvName)

	value, err := strconv.ParseBool(envValue)
	if err != nil {
		return false
	}

	return value
}

func NewSqliteDBImpl() *DBImpl {
	var dsn string

	// For test use the virtual db
	if IsRunningTests() {
		dsn = "file::memory:?cache=shared"
	} else {
		dsn = "sample.db"
	}

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if IsRunningTests() {
		db.AutoMigrate(&Customer{})
	}

	return &DBImpl{
		db: db,
	}
}

func (d DBImpl) GetCustomers(ctx context.Context, filters *domain.FilterInput, pagination domain.PaginationInput) ([]*Customer, error) {
	var customers []*Customer

	query := d.db.Limit(pagination.Limit).Offset(pagination.Offset)

	// add country filter
	if filters.Country != nil {
		// Converts a country code to the form in the db i.e +254 to (254)
		code := fmt.Sprintf("(%s)", strings.Replace(filters.Country.Code(), "+", "", 1))
		query.Where("phone LIKE ?", "%"+code+"%")
	}

	if err := query.Find(&customers).Error; err != nil {
		return nil, fmt.Errorf("error fetching customers: %w", err)
	}

	return customers, nil
}
