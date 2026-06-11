package services

import (
	"eattheitch/backend/models"
	"encoding/json"
	"errors"
	"log"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

const usersFile = "models/mock/users.json"

func LoadUsers() ([]models.User, error) {
	data, err := readUserJson()
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
	data, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(usersFile, data, 0644)
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

func readUserJson() ([]byte, error) {
	if _, err := os.Stat(usersFile); errors.Is(err, os.ErrNotExist) {
		return nil, err
	}

	data, err := os.ReadFile(usersFile)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, errors.New("file empty")
	}
	return data, nil
}
