package usecases

import (
	"clean-arch/models"
	"golang.org/x/crypto/bcrypt"
)

type AddAccount struct{}

func (a AddAccount) Add(ac models.AddAccountModel) (*models.AccountModel, *error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(ac.Password), bcrypt.MinCost)
	if err != nil {
		return nil, &err
	}

	return &models.AccountModel{
		Id:       nil,
		Name:     ac.Name,
		Email:    ac.Email,
		Password: string(hashedPassword),
	}, nil
}
