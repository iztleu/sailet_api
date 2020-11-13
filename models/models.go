package models

import (
	"strings"
)

type Account struct {
	Email    string `json:"email";bson:"email"`
	Password string `json:"password";bson:"password"`
	Info     string `json:"info";bson:"info"`
}

func (account *Account) Validate(base BaseAccount) (string, bool) {
	if !strings.Contains(account.Email, "@") {
		return "Email address is required", false
	}
	_, err := base.GetAccount(account.Email)
	if err != nil {
		return "Account already exist", false
	}

	return "Requirement passed", true
}

type BaseAccount interface {
	Login(string, string) (*Account, error)
	GetAccounts() ([]Account, error)
	GetAccount(string) (*Account, error)
	DeleteAccount(string) (bool, error)
	CreateOrUpdate(account *Account) (*Account, error)
}
