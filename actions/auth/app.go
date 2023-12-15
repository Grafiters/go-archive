package auth

import (
	"archive/models"
	"fmt"

	"github.com/gobuffalo/buffalo"
)

func Configuration(app *buffalo.App) {
	public := app.Group("/auth")

	public.POST("/google", AuthLoginGoogle)
}

func validateAndCreate(user *models.User) (*models.User, error) {
	err := models.DB.Where("email = ?", user.Email).First(user)
	fmt.Println(err)
	if err != nil {
		check, err := createUser(user)
		if err != nil {
			return nil, err
		}

		return check, nil
	}

	return user, nil
}

func createUser(data *models.User) (*models.User, error) {
	tx, err := models.DB.NewTransaction()
	if err != nil {
		return nil, err
	}

	if err := tx.Create(data); err != nil {
		tx.TX.Rollback()
		return nil, err
	}

	tx.TX.Commit()

	return data, nil
}
