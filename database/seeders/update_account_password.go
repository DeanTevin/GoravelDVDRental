package seeders

import (
	"goravel/app/models"

	"github.com/goravel/framework/facades"
)

type UpdateAccountPassword struct {
}

// Signature The name and signature of the seeder.
func (s *UpdateAccountPassword) Signature() string {
	return "UpdateAccountPassword"
}

// Run executes the seeder logic.
func (s *UpdateAccountPassword) Run() error {
	hashed_password, _ := facades.Hash().Make("adminjuga")

	facades.Orm().Query().Model(&models.Staff{}).Where(models.Staff{ID: 1}).Update(models.Staff{Password: hashed_password})
	facades.Orm().Query().Model(&models.Staff{}).Where(models.Staff{ID: 2}).Update(models.Staff{Password: hashed_password})

	return nil
}
