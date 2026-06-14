package services

import (
	"eattheitch/backend/models"
	"eattheitch/backend/utils"
	"encoding/json"
	"errors"
	"log"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

const usersFile = "models/mock/users.json"

func LoadUsers() ([]models.User, error) {
	data, err := utils.ReadJson(usersFile)
	if err != nil {
		log.Printf("could not read users from users.json - %s", err.Error())
		return []models.User{}, nil
	}

	var users []models.User
	if err := json.Unmarshal(data, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func SaveUsers(users []models.User) error {
	if err := utils.WriteJson(usersFile, users); err != nil {
		return err
	}
	return nil
}

func GetUserForEmail(email string) (*models.User, error) {
	users, err := LoadUsers()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if strings.EqualFold(user.Email, email) {
			return &user, nil
		}
	}
	return nil, errors.New("user not found " + email)
}

func GetUserForUsername(username string) (*models.User, error) {
	users, err := LoadUsers()
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		if strings.EqualFold(user.Username, username) {
			return &user, nil
		}
	}
	return nil, errors.New("user not found " + username)
}

func hashUserPassword(password string) ([]byte, error) {
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
