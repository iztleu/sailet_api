package models

import (
	"errors"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MgModel struct {
	session *mgo.Session
}

func NewMgModel(s *mgo.Session) *MgModel {
	return &MgModel{
		session: s,
	}
}

func (m *MgModel) Login(email, password string) (*Account, error) {
	acc, err := m.GetAccount(email)
	if err != nil {
		return &Account{}, err
	}
	if acc.Password != password {
		return &Account{}, errors.New("Invalid login")
	}
	return acc, nil
}

func (m *MgModel) GetAccounts() ([]Account, error) {
	m.session.SetMode(mgo.Monotonic, true)

	c := m.session.DB("AccountService").C("Accounts")
	var Accounts []Account
	err := c.Find(bson.M{}).All(&Accounts)
	if err != nil {
		return Accounts, err
	}
	return Accounts, nil
}

func (m *MgModel) GetAccount(email string) (*Account, error) {

	c := m.session.DB("AccountService").C("Accounts")
	Account := &Account{}
	err := c.Find(bson.M{"email": email}).One(Account)

	if err != nil {
		return Account, err
	}

	return Account, nil
}

func (m *MgModel) DeleteAccount(id string) (bool, error) {

	c := m.session.DB("AccountService").C("Accounts")
	err := c.RemoveId(id)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (m *MgModel) CreateOrUpdateAccount(a *Account) (bool, error) {
	c := m.session.DB("AccountService").C("Accounts")
	_, err := c.UpsertId(a.Email, a)
	if err != nil {
		return false, err
	}
	return true, nil
}
