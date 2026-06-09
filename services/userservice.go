package services

import (
	"eattheitch/backend/models"
	"encoding/json"
	"errors"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetUserFormEmail(email string) (*models.User, error) {
	dat, err := os.ReadFile("models/mock/users.json")
	check(err)

	log.Printf("json dat: %s", dat)

	var users []models.User
	if err := json.Unmarshal(dat, &users); err != nil {
		panic(err)
	}

	for i := range users {
		if users[i].Email == email {
			log.Printf("found user: %s, %s, %s, %s", users[i].Email, users[i].Password, users[i].Username, users[i].Location)
			return &users[i], nil
		}
	}
	return nil, errors.New("user not found " + email)
}
