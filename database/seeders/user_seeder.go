package seeders

import (
	"fmt"
	"gons/app/models"

	"github.com/brianvoe/gofakeit/v6"

	"gorm.io/gorm"
)

func init() {
	RegisterSeeder(func(db *gorm.DB) error {
		dataPlan := 50
		var users []models.Users

		gofakeit.Seed(0)

		for i := 0; i < dataPlan; i++ {
			// create user data using gofakeit
			user := models.Users{
				Name:     gofakeit.Name(),
				Email:    gofakeit.Email(),
				Password: "P@ssw0rd",
			}

			users = append(users, user)
		}

		// execute batch insert to database
		err := db.Create(&users).Error
		if err != nil {
			return err
		}

		fmt.Printf("Successfully generate %d dummy users\n", dataPlan)
		return nil
	})
}
