package oauth

import (
	"github.com/artpar/go-oauth2-server/config"
	"github.com/jinzhu/gorm"
)

// Service struct keeps objects to avoid passing them around
type Service struct {
	cnf *config.Config
	db  *gorm.DB
}

// NewService starts a new Service instance
func NewService(cnf *config.Config, db *gorm.DB) *Service {
	return &Service{cnf: cnf, db: db}
}
