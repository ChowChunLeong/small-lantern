package repository

import (
	"errors"
	"time"

	"github.com/ChowChunLeong/pineapple-language-api.git/form"
	"github.com/ChowChunLeong/pineapple-language-api.git/model"
	"gorm.io/gorm"
)

// FindOrCreateUser handles finding an existing user or creating a new one
func FindOrCreateUser(db *gorm.DB, request form.OAuthRequest) (*model.User, error) {
	var user model.User
	
	err := db.Where("email = ?", request.Email).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		// Create new user
		newUser := model.User{
			Email: request.Email,
			Name:  request.Name,
			Image: request.Image,
			Account: &model.Account{
				Provider:          request.Provider,
				ProviderAccountID: request.Email, // Use email as provider ID
			},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		if err := db.Create(&newUser).Error; err != nil {
			return nil, errors.New("failed to create user")
		}
		return &newUser, nil
	} else if err != nil {
		return nil, err
	}

	return &user, nil
}