package q

import (
	"fmt"
	"jobsync-be/configs"
	"jobsync-be/models"
)

// Fungsi untuk membuat booking rental baru
func Create(user models.User) error {
	res := configs.DB.Create(&user)
	if res.Error != nil {
		return fmt.Errorf("failed to create on database: %v", res.Error)
	}
	return nil
}
