package services

import (
	"eattheitch/backend/models"
	"eattheitch/backend/utils"
	"errors"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const usersFile = "models/mock/users.json"

func SaveUser(newUser *models.User) error {
	if err := utils.UpsertJSON(usersFile, newUser); err != nil {
		return err
	}
	return nil
}

func GetUserForId(userId uuid.UUID) (*models.User, error) {
	user, err := utils.GetJSON[*models.User](usersFile, userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(user models.User) error {
	if err := utils.UpsertJSON(usersFile, &user); err != nil {
		return err
	}
	return nil
}

func GetUserForEmail(email string) (*models.User, error) {
	users, err := utils.ReadJSON[*models.User](usersFile)
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if strings.EqualFold(user.Email, email) {
			return user, nil
		}
	}
	return nil, errors.New("user not found " + email)
}

func GetUserForUsername(username string) (*models.User, error) {
	users, err := utils.ReadJSON[*models.User](usersFile)
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		if strings.EqualFold(user.Username, username) {
			return user, nil
		}
	}
	return nil, errors.New("user not found " + username)
}

func HashUserPassword(password string) ([]byte, error) {
	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, err
	}
	return passwordHash, nil
}

func VerifyUserPassword(email string, password string) error {
	user, err := GetUserForEmail(email)
	if err != nil {
		return err
	}
	if compareResult := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); compareResult != nil {
		return compareResult
	}
	return nil
}
