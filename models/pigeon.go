package models

import (
	"math/rand"
)

type Pigeon struct {
	Id          uint   `json:"id"`
	PigeonCode  int    `json:"pigeonCode"`
	YearOfBirth int    `json:"yearOfBirth"`
	Gender      string `json:"gender"`
	Colour      string `json:"colour"`
	Mother      int    `json:"mother"`
	Father      int    `json:"father"`
	OwnerEmail  string `json:"ownerEmail"`
}

func GeneratePigeonCode() int {
	code := rand.Intn(100000)
	return code
}
